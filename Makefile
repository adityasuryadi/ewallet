migrate-up:
	migrate -database "postgres://postgres:postgres@localhost:5432/ewallet_ums?sslmode=disable&search_path=public" -path ./db/migrations up