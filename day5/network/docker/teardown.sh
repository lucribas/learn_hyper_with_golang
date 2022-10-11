#!/bin/bash
#set -e

echo "************************************"
echo "TEARDOWN"
echo "************************************"
docker-compose -f docker-compose.yaml down

# Shut down the Docker containers for the system tests.
docker-compose -f docker-compose.yaml kill && docker-compose -f docker-compose.yaml down

# remove the local state
rm -f ~/.hfc-key-store/*

#TODO: ADD FILTER LABEL to remove only for this project
## remover apenas as imagens do ORDERER e do PEER
## manter os do java para ficar mais rapido a subida

# remove chaincode docker images
docker rm $(docker ps -aq)
# remove volumes
docker volume rm $(docker volume ls -q)

###docker-compose -f docker-compose.yaml up --build --force-recreate

# remove previous crypto material and config transactions
sudo chmod -R 0755 ./crypto-config
sudo rm -fr config/*
sudo rm -fr crypto-config/ordererOrganizations/*
sudo rm -fr crypto-config/peerOrganizations/*
sudo rm -f *.block

# Your system is now clean
