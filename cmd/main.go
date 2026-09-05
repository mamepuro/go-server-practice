package main

import (
	"fmt"
	"go-server-practice/cmd/internal/controller"
	"go-server-practice/cmd/internal/repository"
	"go-server-practice/cmd/internal/usecase"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := fmt.Sprintf(
		"host=localhost user=%s password=%s dbname=go_server_practice_dev port=5432 sslmode=disable TimeZone=Asia/Tokyo",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
	) // .env / dbconfig.yml と同じ接続情報
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userCtrl := controller.NewUserController(userUC)

	r := gin.Default()
	userCtrl.RegisterRoutes(r)
	r.Run(":8080")
}
