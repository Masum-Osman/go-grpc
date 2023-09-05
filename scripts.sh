protoc -Igreet/proto --go_opt=module=go-grpc --go_out=. --go-grpc_opt=module=go-grpc --go-grpc_out=. greet/proto/*.proto

go build -o bin/greet/client ./greet/client
go build -o bin/greet/server ./greet/server