
# 1. Hyperledger blockchain

## 1.1 Install docker and docker-compose
```bash
sudo apt install docker
sudo apt install docker-compose
```

## 1.2 For development use start a container with Portainer.io (Optional)

We are using the Portainer CE (Comunity Edition) Portainer container service tools to help container management for Kubernetes, Docker, Swarm and ACI.

```bash
docker volume create portainer_data
docker run -d -p 8000:8000 -p 9000:9000 --name=portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer-ce
docker run -d -p 9001:9001 --name portainer_agent --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v /var/lib/docker/volumes:/var/lib/docker/volumes portainer/agent
```
Open: http://172.17.0.2:9000/

## 1.3 Start containers in Docker-compose


```bash
# $GIT_TOP_PRJ is where your repo is
cd $GIT_TOP_PRJ/blockchain-backend/network/docker

./start_all.sh
```

Open the block explorer: <http://localhost:8080/#/login>
- Network="Test Network"
- User="exploreradmin"
- Password="exploreradminpw"

as specified in ./explorer/connection-profile/test-network.json



# Throubleshooting

Open the Dockstation as GUEST and check if all containers are running.

Check the logs:
```bash
docker logs peer0.org1.example.com -f
docker logs example.com -f
docker logs couchdb -f
```


## 1.2 Visual Studio Code

### 1.2.1 Prerequisites:
Install the following extensions:

### 1.2.1 Open workspace

Open the workpace file:
`$GIT_TOP_PRJ\blockchain-backend\network\docker\workspace.code-workspace`

Or open workspace in folder `$GIT_TOP_PRJ\blockchain-backend\network\docker\` and start devcontainer when asked.

### 1.2.2 Start devcontainer

The Visual Studio Code will ask for open devcontainer. Accept it and wait for workspace start.


### 1.2.3 Open block explorer

Open the block explorer: <http://localhost:8080/#/login>
- Network="Test Network"
- User="exploreradmin"
- Password="exploreradminpw"

as specified in .\explorer\connection-profile\test-network.json

### Execute a contract test


#### Hello

Usin VSC open a ash in hello folder:

```bash
cd /git/network/docker/chaincode/hello
mkdir vendor

#go mod init hello/m
go get
go mod tidy
go build

# add gcc:
apk add build-base
go test -v
```


# References

CouchDB:
http://localhost:5984/_utils/#login

BlockExplorer:
http://localhost:8080/#/

## VSC

https://code.visualstudio.com/docs/editor/multi-root-workspaces


## Hyperledger

https://hyperledger-fabric-ca.readthedocs.io/en/latest/operations_guide.html
https://hyperledger-fabric.readthedocs.io/en/release-2.2/architecture.html

## Blockchain explorer

https://github.com/hyperledger/blockchain-explorer


## Go

docker exec -it go-container bash

https://medium.com/@adiach3nko/package-management-with-go-modules-the-pragmatic-guide-c831b4eaaf31



## TODO
- HOW debug:
  - https://github.com/golang/vscode-go/blob/master/docs/debugging.md
  - https://golangforall.com/en/post/go-docker-delve-remote-debug.html
  - https://davidkel.github.io/docs/vscode/godebug.html
  - https://medium.com/@kaperys/delve-into-docker-d6c92be2f823






# JAVA
apk add maven