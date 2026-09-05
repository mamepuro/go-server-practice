package repository

import (
	"context"
	"go-server-practice/cmd/internal/domain"
	"go-server-practice/cmd/internal/infra/model"
	"go-server-practice/cmd/internal/infra/query"
	"go-server-practice/cmd/internal/usecase"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

type userRepository struct {
	q *query.Query
}

func NewUserRepository(db *gorm.DB) usecase.UserRepository {
	return &userRepository{q: query.Use(db)}
}

func (r *userRepository) FindAll(ctx context.Context) ([]*domain.User, error) {
	rows, err := r.q.User.WithContext(ctx).Find()
	if err != nil {
		return nil, err
	}
	users := make([]*domain.User, 0, len(rows))
	for _, row := range rows {
		u, err := toDomainUser(row)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *userRepository) FindByID(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	row, err := r.q.User.WithContext(ctx).Where(r.q.User.ID.Eq(id.String())).First()
	if err != nil {
		return nil, err
	}
	return toDomainUser(row)
}

func (r *userRepository) CreateUser(ctx context.Context, u *domain.User) error {
	row := &model.User{
		Name:  u.Name,
		Email: u.Email,
		// ID はDB側の DEFAULT gen_random_uuid() に任せるので指定しない
	}
	if err := r.q.User.WithContext(ctx).Create(row); err != nil {
		return err
	}

	created, err := toDomainUser(row)
	if err != nil {
		return err
	}
	*u = *created
	return nil
}

func (r *userRepository) UpdateUser(ctx context.Context, u *domain.User) error {
	_, err := r.q.User.WithContext(ctx).
		Where(r.q.User.ID.Eq(u.ID.String())).
		Updates(&model.User{Name: u.Name, Email: u.Email})
	return err
}

func (r *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	_, err := r.q.User.WithContext(ctx).Where(r.q.User.ID.Eq(id.String())).Delete()
	return err
}

func toDomainUser(row *model.User) (*domain.User, error) {
	id, err := uuid.Parse(row.ID)
	if err != nil {
		return nil, err
	}
	u := &domain.User{
		ID:    id,
		Name:  row.Name,
		Email: row.Email,
	}
	if row.CreatedAt != nil {
		u.CreatedAt = *row.CreatedAt
	}
	if row.UpdatedAt != nil {
		u.UpdatedAt = *row.UpdatedAt
	}
	return u, nil
}
