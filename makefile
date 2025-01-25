DB_DSN := "postgres://postgres:7819900@localhost:5432/postgres?sslmode=disable"
MIGRATE := migrate -path ./migrations -database $(DB_DSN)

# Создание новой миграции
migrate-new:
	migrate create -ext sql -dir ./migrations $(NAME)

# Применение миграций
migrate:
	$(MIGRATE) up

# Откат миграций
migrate-down:
	$(MIGRATE) down

# Установка зависимостей
deps:
	go mod tidy

# Запуск приложения
run:
	go run cmd/app/main.go

# Генерация API
gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go

# Линтер - инструмент, который анализирует нашу кодовую базу и указывает на все места в которых есть ошибки.
lint:
	golangci-lint run --out-format=colored-line-number

gen:
	oapi-codegen -config openapi/.openapi -include-tags tasks -package tasks openapi/openapi.yaml > ./internal/web/tasks/api.gen.go
	oapi-codegen -config openapi/.openapi -include-tags users -package users openapi/openapi.yaml > ./internal/web/users/api.gen.go