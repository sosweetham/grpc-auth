make-go-proto:
	rm -f pb/*.go && \
	protoc --proto_path=protos --go_out=. --go-grpc_out=. protos/*.proto \
	&& cd server && go mod tidy