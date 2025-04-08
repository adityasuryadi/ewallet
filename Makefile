migrate-up:
	migrate -database "postgres://postgres:password@localhost:5432/ewallet_ums?sslmode=disable&search_path=public" -path ./db/migrations up