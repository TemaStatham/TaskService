docker:
	docker compose up -d

createdb:
	docker exec -it dapomogu_task_postgres createdb --username=postgres --owner=postgres dapomogu_task_db

dropdb:
	docker exec -it dapomogu_task_postgres dropdb --username=postgres dapomogu_task_db

migrateup:
	migrate -path migrations -database "postgresql://postgres:dapomogu_password@localhost:5436/dapomogu_task_db?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations -database "postgresql://postgres:dapomogu_password@localhost:5436/dapomogu_task_db?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown