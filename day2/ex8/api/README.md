

Protobuf compiler installing:

go get -u github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go
go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

sudo apt install protobuf-compiler 

protoc --go_out=. --go-grpc_out=. api/user.proto  --plugin=$(go env GOPATH)/bin/protoc-gen-go-grpc