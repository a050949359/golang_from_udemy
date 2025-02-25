# golang_from_udemy
Backend Master Class [Golang + Postgres + Kubernetes + gRPC]

viper gin migration sqlc

# at host

postgres:
    docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=123456 -d postgres:17.3-alpine

createdb:
    docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
    docker exec -it postgres dropdb simple_bank


# in container

migrate_create:
    migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
    migrate -path db/migration -database "postgresql://root:123456@postgres:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
    migrate -path db/migration -database "postgresql://root:123456@postgres:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
    sqlc init
    sqlc generate

test: 
    go test -v -cover ./...

lesson 9 descibe mysql postgresql isolation level & diff

鬼打牆的 = =
mockgen -destination db/mock.store.go github.com/golang_from_udemy/db/sqlc Store