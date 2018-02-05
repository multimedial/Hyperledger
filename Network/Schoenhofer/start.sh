#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#
# Exit on first error, print all commands.
set -ev

# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1


docker-compose -f docker-compose.yml down

docker-compose -f docker-compose.yml up -d 


# wait for Hyperledger Fabric to start
# incase of errors when running later commands, issue export FABRIC_START_TIMEOUT=<larger number>
export FABRIC_START_TIMEOUT=3
sleep ${FABRIC_START_TIMEOUT}

docker-compose -f docker-compose.yml up -d 


# Create the channel
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@amt1.aufsicht.de/msp" peer1.amt1.aufsicht.de peer channel create -o orderer.aufsicht.de:7050 -c vertraulich -f /etc/hyperledger/configtx/channel.tx


# peer 1 joins
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@amt1.aufsicht.de/msp" peer1.amt1.aufsicht.de peer channel join -b vertraulich.block --logging-level=notice

# let peer 2 join
docker exec peer1.amt2.aufsicht.de peer channel fetch config vertraulich.block -o orderer.aufsicht.de:7050 -c vertraulich
docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@amt2.aufsicht.de/msp" peer1.amt2.aufsicht.de peer channel join -b vertraulich.block

# let peer 3 join
docker exec peer1.amt3.aufsicht.de peer channel fetch config vertraulich.block -o orderer.aufsicht.de:7050 -c vertraulich
docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@amt3.aufsicht.de/msp" peer1.amt3.aufsicht.de peer channel join -b vertraulich.block

# populate the msql container with the SQL definitition for the table
# TBA!