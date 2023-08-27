
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

