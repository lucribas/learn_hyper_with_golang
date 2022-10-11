#!/bin/bash
##
# Exit on first error, print all commands.
# set -ev

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

echo
echo " ____    _____      _      ____    _____ "
echo "/ ___|  |_   _|    / \    |  _ \  |_   _|"
echo "\___ \    | |     / _ \   | |_) |   | |  "
echo " ___) |   | |    / ___ \  |  _ <    | |  "
echo "|____/    |_|   /_/   \_\ |_| \_\   |_|  "
echo


echo "##########################################################"
echo "##### Dev network is starting #########"
echo "##########################################################"

# Shutting down exisiting network
docker-compose -f docker-compose.yaml down

# Starting hyperledger fabricchannel
docker-compose -f docker-compose.yaml up -d cli-network

# wait for Hyperledger Fabric to start
# incase of errors when running later commands, issue export FABRIC_START_TIMEOUT=<larger number>
export FABRIC_START_TIMEOUT=2
echo "waiting "${FABRIC_START_TIMEOUT}"s.."
sleep ${FABRIC_START_TIMEOUT}

docker exec -it cli-network /git/network/docker/cli_generate.sh

