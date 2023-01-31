refresh-db:
	migrate -source file://migrations -database "postgres://localhost:5432/teknologi62-test?sslmode=disable&user=postgres&password=kepler22b" down 1
	migrate -source file://migrations -database "postgres://localhost:5432/teknologi62-test?sslmode=disable&user=postgres&password=kepler22b" up 1