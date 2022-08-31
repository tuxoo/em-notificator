.DEFAULT_GOLA := proto-gen

proto-gen:
	protoc  --go_out=internal/transport/grpc --go_opt=paths=import \
	 --go-grpc_out=internal/transport/grpc --go-grpc_opt=paths=import \
	  internal/transport/grpc/proto/em-notificator.proto

build:
	go build -o ./.bin/app ./cmd/em-notificator/main.go