#!/bin/bash
##
# Exit on first error, print all commands.
# set -ev

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

cd /git/network/docker
export PATH=$GOPATH/src/github.com/hyperledger/fabric/build/bin:${PWD}/../bin:${PWD}:$PATH
export FABRIC_CFG_PATH=${PWD}
CHANNEL_NAME=devchannel

# remove previous crypto material and config transactions
apk add figlet

figlet "Generate"

chmod -R 0755 ./crypto-config
rm -fr config/*
rm -fr crypto-config/ordererOrganizations/*
rm -fr crypto-config/peerOrganizations/*
rm -f *.block

# generate crypto material
cryptogen generate --config=./crypto-config.yaml
if [ "$?" -ne 0 ]; then
  echo "Failed to generate crypto material..."
  exit 1
fi

# generate genesis block for orderer
configtxgen -profile OrdererGenesis -channelID $CHANNEL_NAME -outputBlock ./config/genesis.block
configtxgen -profile OrdererGenesis -channelID sys-channel -outputBlock ./config/genesis.block
if [ "$?" -ne 0 ]; then
  echo "Failed to generate orderer genesis block..."
  exit 1
fi

# generate channel configuration transaction
configtxgen -profile BasicChannel -outputCreateChannelTx ./config/channel.tx -channelID $CHANNEL_NAME
if [ "$?" -ne 0 ]; then
  echo "Failed to generate channel configuration transaction..."
  exit 1
fi

# generate anchor peer transaction
configtxgen -profile BasicChannel -outputAnchorPeersUpdate ./config/Org1MSPanchors.tx -channelID $CHANNEL_NAME -asOrg Org1MSP
if [ "$?" -ne 0 ]; then
  echo "Failed to generate anchor peer update for Org1MSP..."
  exit 1
fi
