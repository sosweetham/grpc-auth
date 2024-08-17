make-go-proto:
	rm -f server/pkg/pb/*.go && \
	protoc --proto_path=protos --go_out=. --go-grpc_out=. protos/*.proto \
	&& cd server && go mod tidy