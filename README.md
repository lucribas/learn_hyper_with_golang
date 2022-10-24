
![Gopher](assets/images/gopher.png)

# Description

Learning Golang by examples with TDD
- Using vscode with devcontainer (+docker-compose)
- Integrated Github Actions CI/CD (build and test) (with docker-compose)


# Requirements

## Windows

Install WSL with Ubuntu 20.04:
https://ubuntu.com/wsl

## Windows (WSL + Ubuntu 20.04) / Linux (Ubuntu 20.04+)

Initial setup
```bash
sudo apt update
```


- GIT
```bash
sudo apt install git
```
- Docker
```bash
sudo apt install docker docker-compose
sudo usermod -a -G docker $USER
sudo apt install gconf2 gconf-service libappindicator1
```

No windows seguir o [tutorial](https://medium.com/codigorefinado/docker-no-linux-dentro-do-windows-10-com-wsl-2-f52b91931267)


- Vscode
  Install vscode from link: [vscode]](https://code.visualstudio.com/)
  Install extension:
    + Dev Containers extension: https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers

- (on Linux only) Dockstation (Optional)
```bash
wget -c https://github.com/DockStation/dockstation/releases/download/v1.5.1/dockstation_1.5.1_amd64.deb
sudo dpkg -i dockstation_1.5.1_amd64.deb
```


# Initial Setup

```bash
git clone https://github.com/lucribas/learn_hyper_with_golang.git
cd learn_hyper_with_golang
```

## day 1

```bash
cd day1
code .
# Open in with Dev Containers extension
```


References:
- Go Lang do Zero : https://www.youtube.com/playlist?list=PL5aY_NrL1rjucQqO21QH8KclsLDYu1BIg
- https://www.digitalocean.com/community/tutorials/defining-structs-in-go

https://github.com/0xAX/go-algorithms

https://www.youtube.com/watch?v=NF1pwjL9-DE&ab_channel=Computerphile
https://www.youtube.com/watch?v=NmM9HA2MQGI&ab_channel=Computerphile

https://github.com/avelino/awesome-go




-----
new directory:

mkdir exX
cd exX
add your file.go
go mod init
go mod tidy
go work use .

go run file.go