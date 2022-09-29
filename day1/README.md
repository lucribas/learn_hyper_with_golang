


$ docker build -t my-golang-app .
$ docker run -it --rm --name my-running-app my-golang-app

cd ex1
go mod init example.com/hello
cd ..

cd ex1_solved
go mod init example.com/hello_solved
cd ..

go work init ./ex1
go work use ./ex1_solved

