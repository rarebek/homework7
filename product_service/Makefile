CURRENT_DIR=$(shell pwd)

DB_URL := "postgres://postgres:nodirbek@localhost:5432/products?sslmode=disable"

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

migrate-up:
	migrate -path migrations -database "$(DB_URL)" -verbose up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" -verbose down

migrate-file:
	migrate create -ext sql -dir migrations/ -seq create_users_products_table

proto-p:
	protoc --go_out=genproto/product_service --go-grpc_out=genproto/product_service protos/product_service/product.proto

proto-u:
	protoc --go_out=genproto/user_service --go-grpc_out=genproto/user_service protos/user_service/user.proto