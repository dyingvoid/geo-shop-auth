.PHONY: gen-grpc migrate test

ENV ?= local
CONFIG_FILE := config.$(ENV).yaml

gen_grpc:
	protoc --go_out=. --go-grpc_out=. proto/*.proto

migrate:
	go run scripts/get_db_conn.go --config $(CONFIG_FILE) > .connstring.tmp
	migrate -path db/migrations postgres -database "$$(cat .connstring.tmp)" up

test:
	go test ./...