-include .env
export

gen-model:
	gentool \
	-db postgres \
	-dsn "host=localhost user=$$POSTGRES_USER password=$$POSTGRES_PASSWORD dbname=go_server_practice_dev port=5432 sslmode=disable" \
	-tables "users" \
	-outPath "./cmd/internal/infra/query" \
	-modelPkgName "model" \
	-fieldNullable \
	-fieldWithIndexTag \
	-fieldWithTypeTag