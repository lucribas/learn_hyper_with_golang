#!/bin/bash

docker exec -it cli-network figlet "Start Explorer"

echo "##########################################################"
echo "##### Dev network is starting #########"
echo "##########################################################"

# Starting hyperledger fabricchannel
docker-compose -f docker-compose.yaml up -d
