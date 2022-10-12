

# first time
cd /git/network/docker/network/docker/application-go
rm -rf go.mod go.sum
go mod init go_app
go mod vendor

# Compilação Estática Com Golang
cd /git/network/docker/network/docker/application-go
CGO_ENABLED=0 go build
./go_app
