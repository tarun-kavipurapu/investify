DB_URL=postgresql://root:secret@localhost:5432/lendify?sslmode=disable
#migrate -path PATH_TO_YOUR_MIGRATIONS -database YOUR_DATABASE_URL force VERSION


migrateup:
	go run cmd/migrations/init/init.sql.go up




migratedown:
	go run cmd/migrations/init/init.sql.go down


sqlc:
	sqlc generate


server:
	go run cmd/app/main.go
