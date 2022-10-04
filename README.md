
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
sudo apt install docker-compose
sudo usermod -a -G docker $USER
sudo apt install gconf2 gconf-service libappindicator1
```
- Dockstation (Optional)
```bash
wget -c https://github.com/DockStation/dockstation/releases/download/v1.5.1/dockstation_1.5.1_amd64.deb
sudo dpkg -i dockstation_1.5.1_amd64.deb
```
- [Vscode](https://code.visualstudio.com/)
  + Dev Containers extension: https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers


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
- https://www.digitalocean.com/community/tutorials/defining-structs-in-go