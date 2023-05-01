.SILENT:

run-dev:
	go run ./cmd/main.go
	docker-compose up -d

migrate:
	echo "Database migrate up..."
	migrate -path ./database/migration -database 'postgres://web:web@localhost:5454/web?sslmode=disable' up

migrate-down:
	echo "Database migration down..."
	migrate -path ./database/migration -database 'postgres://web:web@localhost:5454/web?sslmode=disable' down
