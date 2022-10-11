#!/bin/bash
##

# cd /git/network/docker
# export PATH=$GOPATH/src/github.com/hyperledger/fabric/build/bin:${PWD}/../bin:${PWD}:$PATH
# export FABRIC_CFG_PATH=${PWD}

apt update
apt install -y figlet curl jq wget
#install yq
VERSION=v4.2.0
BINARY=yq_linux_amd64
wget https://github.com/mikefarah/yq/releases/download/${VERSION}/${BINARY}.tar.gz -O - |\
  tar xz && mv ${BINARY} /usr/bin/yq

wget https://github.com/moparisthebest/static-curl/releases/download/v7.85.0/curl-amd64
chmod +x curl-amd64
mv curl-amd64 /usr/bin/curl

alias figlet=echo

export CHANNEL_NAME=devchannel
export CHAINCODE=energy

export TOP=/opt/gopath/src/github.com/hyperledger/fabric/peer
cd $TOP

figlet "Start Chaincode"

echo ""
echo "#####  Wait till RAFT nodes are ready to accept transactions!  #######"
echo ""
#for kubernetes
#ORDERERS=$(kubectl get svc | awk '/orderer[0-9]*-org[0-9]*/ {print $3}')
ORDERERS=$(yq -j eval /git/network/docker/configtx.yaml | jq '.Orderer.EtcdRaft.Consenters[].Host' | tr -d '"')

while [ true ]; do
    for ORDERER in $ORDERERS; do
        IS_LEADER=$(curl -v --silent http://${ORDERER}:8443/metrics 2>&1 | grep sys-channel | awk '/^consensus_etcdraft_is_leader/ {print $2}')
        [ "${IS_LEADER}" == "1" ] && break
		echo ${ORDERER} : ${IS_LEADER}
    done
    [ "${IS_LEADER}" == "1" ] && break
    sleep 1
done
echo "ordering service is ready to accept transactions!"



echo ""
echo "##### Creating channel #########"
echo ""
echo "-Because we are using a Raft ordering service, you may get some status unavailable messages that you can safely ignore. The command will return the genesis block of the new channel"
peer channel create -o orderer.example.com:7050 -c $CHANNEL_NAME -f /etc/hyperledger/channel/channel.tx --tls --cafile /etc/hyperledger/crypto/orderer/msp/tlscacerts/tlsca.example.com-cert.pem

echo ""
echo "##### Join channel #########"
echo ""

peer channel fetch 0 bcchannel.block -c $CHANNEL_NAME  -o orderer.example.com:7050  --tls --cafile /etc/hyperledger/crypto/orderer/msp/tlscacerts/tlsca.example.com-cert.pem
peer channel join -b bcchannel.block  --tls --cafile /etc/hyperledger/crypto/orderer/msp/tlscacerts/tlsca.example.com-cert.pem


echo ""
echo "##### List channels #########"
echo ""

peer channel list

echo ""
echo "##### Update Anchor peer #########"
echo ""

peer channel update -o orderer.example.com:7050 -c $CHANNEL_NAME -f /etc/hyperledger/channel/Org1MSPanchors.tx --tls --cafile /etc/hyperledger/crypto/orderer/msp/tlscacerts/tlsca.example.com-cert.pem


echo ""
echo "##### Package chaincode #########"
echo ""
cd /chaincode/$CHAINCODE/go/
go mod init energy
go mod vendor
go get
go mod tidy
go build
cd $TOP

peer lifecycle chaincode package /chaincode/$CHAINCODE.tar.gz --path /chaincode/$CHAINCODE/go/bin --label $CHAINCODE --tls --cafile /etc/hyperledger/crypto/orderer/msp/tlscacerts/tlsca.example.com-cert.pem
# peer lifecycle chaincode package /chaincode/$CHAINCODE.tar.gz --path /chaincode/$CHAINCODE/go --label $CHAINCODE --tls --cafile /etc/hyperledger/crypto/orderer/msp/tlscacerts/tlsca.example.com-cert.pem


echo ""
echo "##### Install chaincode #########"
echo "(please wait! it can takes until 2 minutes)"
echo ""

peer lifecycle chaincode install /chaincode/$CHAINCODE.tar.gz --tls --cafile /etBlockchain4everc/hyperledger/crypto/orderer/msp/tlscacerts/tlsca.example.com-cert.pem

echo ""
echo "##### Query installed chaincode #########"
echo ""
peer lifecycle chaincode queryinstalled
export PACKAGE_ID=$(peer lifecycle chaincode queryinstalled -O json | jq '.installed_chaincodes[] | select(.label=="'$CHAINCODE'") | .package_id' | tr -d '"')
echo $PACKAGE_ID

echo ""
echo "##### Approve chaincode for org #########"
echo ""

peer lifecycle chaincode approveformyorg --channelID $CHANNEL_NAME --name $CHAINCODE --version 1.0 --init-required --package-id $PACKAGE_ID --sequence 1 -o orderer.example.com:7050 --tls --cafile /etc/hyperledger/crypto/orderer/msp/tlscacerts/tlsca.example.com-cert.pem

echo ""
echo "##### Check commit readiness #########"
echo ""

peer lifecycle chaincode checkcommitreadiness --channelID $CHANNEL_NAME --name $CHAINCODE --version 1.0 --init-required --sequence 1 -o orderer.example.com:7050 --tls --cafile /etc/hyperledger/crypto/orderer/msp/tlscacerts/tlsca.example.com-cert.pem

echo ""
echo "##### Commit chaincode #########"
echo ""

peer lifecycle chaincode commit \
-o orderer.example.com:7050 \
--channelID $CHANNEL_NAME --name $CHAINCODE --version 1.0 --sequence 1 --init-required \
--tls --cafile /etc/hyperledger/crypto/orderer/msp/tlscacerts/tlsca.example.com-cert.pem \
--tlsRootCertFiles /etc/hyperledger/crypto/peer/tls/ca.crt \
--peerAddresses peer0.org1.example.com:7051



peer lifecycle chaincode querycommitted  \
--channelID $CHANNEL_NAME --name $CHAINCODE \
--tls --cafile /etc/hyperledger/crypto/orderer/msp/tlscacerts/tlsca.example.com-cert.pem

echo ""
echo "##### Init chaincode #########"
echo ""


peer chaincode invoke \
-o orderer.example.com:7050 \
--tls --cafile /etc/hyperledger/crypto/orderer/msp/tlscacerts/tlsca.example.com-cert.pem \
-C $CHANNEL_NAME -n $CHAINCODE  \
--peerAddresses peer0.org1.example.com:7051 \
--tlsRootCertFiles /etc/hyperledger/crypto/peer/tls/ca.crt \
-c '{"Args":[]}' --isInit --waitForEvent



# echo ""
# echo "##### Test transactions - init $CHAINCODE #########"
# echo ""


# peer chaincode invoke \
# -o orderer.example.com:7050 \
# --tls --cafile /etc/hyperledger/crypto/orderer/msp/tlscacerts/tlsca.example.com-cert.pem \
# -C $CHANNEL_NAME -n $CHAINCODE \
# --peerAddresses peer0.org1.example.com:7051 \
# --tlsRootCertFiles /etc/hyperledger/crypto/peer/tls/ca.crt \
# -c '{"Args":["initMarble","marble1","blue","35","tom"]}' --waitForEvent



# echo ""
# echo "##### Query chaincode after first aid #########"
# echo ""


# peer chaincode query -C $CHANNEL_NAME -n $CHAINCODE -c '{"Args":["readMarble","marble1"]}'


echo "##########################################################"
echo "##### Network is finishing #########"
echo "##########################################################"

exit 0

