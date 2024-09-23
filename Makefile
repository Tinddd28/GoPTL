postgres:
	docker-compose up -d

outps:
	docker-compose down

gotodb:
	docker exec -it goptl-psql psql -U postgres crypto-platform

createdb:
	docker exec -it goptl-psql createdb -U postgres -O postgres crypto-platform

dropdb:
	docker exec -it goptl-psql dropdb -U postgres crypto-platform -f

migrateup:
	migrate -path internal/migrations -database "postgresql://postgres:root@localhost:5432/crypto-platform?sslmode=disable" -verbose up

migratedown:
	migrate -path internal/migrations -database "postgresql://postgres:root@localhost:5432/crypto-platform?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...


.PHONY: postgres outps gotodb createdb dropdb migrateup migratedown sqlc