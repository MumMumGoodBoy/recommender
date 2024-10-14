generate:
	protoc --go_out=./internal/proto --go_opt=paths=source_relative \
	--go-grpc_out=./internal/proto --go-grpc_opt=paths=source_relative \
    recommender.proto

compose-up:
	docker-compose up -d

compose-down:
	docker-compose down

dev:
	make generate
	goreload main.go
