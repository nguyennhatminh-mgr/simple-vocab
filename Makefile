include .env

createdb:
	docker exec -it simple-vocab-postgres-db createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} simple_vocabulary
dropdb:
	docker exec -it simple-vocab-postgres-db dropdb simple_vocabulary
migrateUp:
	migrate -path database/migration/ -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}/simple_vocabulary?sslmode=disable" -verbose up
migrateDown:
	migrate -path database/migration/ -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}/simple_vocabulary?sslmode=disable" -verbose down