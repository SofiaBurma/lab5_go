migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/api_db?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/api_db?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/api_db?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/api_db?sslmode=disable" -verbose down 1

run:
	go run cmd/api/main.go

.PHONY: migrateup migrateup1 migratedown migratedown1 run