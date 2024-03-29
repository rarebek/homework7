CURRENT_DIR=$(shell pwd)

DB_URL := "postgres://postgres:nodirbek@localhost:5432/users?sslmode=disable"

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto:
	protoc --go_out=genproto/user_service --go-grpc_out=genproto/user_service protos/user_service/user.proto

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migrate-file:
	migrate create -ext sql -dir migrations/ -seq create_users_table

swag-gen:
  ~/go/bin/swag init -g ./api/router.go -o api/docs

