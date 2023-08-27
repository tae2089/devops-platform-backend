
run:
	go run run/main.go

build:
	go build -o server run/main.go

generate-ent:
	go generate ./ent

install:
	go mod tidy

new-schema:
	go run -mod=mod entgo.io/ent/cmd/ent new ${schema}

docker-compose-run:
	docker compose -f build/cd/docker-compose.yaml --env-file .env up -d

docker-compose-only-db:
	docker compose -f build/cd/docker-compose.postgres.yaml --env-file .env up -d