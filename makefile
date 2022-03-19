generate:
	@protoc --go_out=plugins=grpc:. ./chat.proto

run:
	@echo "---- Running Server ----"
	@go run main.go

run_client:
	@echo "---- Running Client ----"
	@go run ./client/client.go