#!/bin/bash

./teardown.sh

./generate.sh
./start_network.sh
./start_chaincode.sh
./start_explorer.sh

# Because containers runs as root we need fix the file permissions:
sudo chown -R $USER.$USER chaincode/
sudo chown -R $USER.$USER crypto-config/
sudo chown -R $USER.$USER config/
sudo find . -type f -exec chmod 655 -- {} +
sudo find . -type f -iname "*.sh" -exec chmod 755 -- {} +