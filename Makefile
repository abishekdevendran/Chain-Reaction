# Makefile

.PHONY: proto db-gen server migrate-up

# Command to regenerate all protobuf/gRPC code
proto:
	@protoc --proto_path=proto \
		--go_out=./backend \
		--go-grpc_out=./backend \
		proto/*.proto

# Command to regenerate all sqlc database code
db-gen:
	@cd backend && sqlc generate

# Command to run the backend server
server:
	@echo "Starting backend server..."
	@cd backend && go run ./server/main.go

# Command to run all "up" database migrations
migrate-up:
	@migrate -database "$$DATABASE_URL" -path backend/db/migration up