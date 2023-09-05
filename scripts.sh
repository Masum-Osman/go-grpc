protoc -Igreet/proto --go_opt=module=go-grpc --go_out=. --go-grpc_opt=module=go-grpc --go-grpc_out=. greet/proto/*.proto

go build -o bin/greet/client ./greet/client
go build -o bin/greet/server ./greet/server

protoc -Icalculator/proto --go_opt=module=go-grpc --go_out=. --go-grpc_opt=module=go-grpc --go-grpc_out=. calculator/proto/*.proto
go build -o bin/calculator/client ./calculator/client
go build -o bin/calculator/server ./calculator/server