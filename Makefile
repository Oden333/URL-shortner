postgres:
	docker run --name URL-shortener -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=qwerty -d postgres:12-alpine

createdb:
	docker exec -it URL-shortener createdb --username=root --owner=root urls

dropdb:
	docker exec -it URL-shortener dropdb urls

migrateup:
	migrate -path schema -database "postgresql://root:qwerty@localhost:5432/urls?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:qwerty@localhost:5432/urls?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown