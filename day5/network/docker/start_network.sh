#!/bin/bash


echo "************************************"
echo "START NETWORK"
echo "************************************"

echo "##########################################################"
echo "##### Dev network is starting #########"
echo "##########################################################"

# Shutting down exisiting network
docker-compose -f docker-compose.yaml down

# Starting hyperledger fabricchannel
docker-compose -f docker-compose.yaml up -d orderer.example.com peer0.org1.example.com couchdb cli-network
# docker-compose -f docker-compose.yaml up -d cli-network
