#!/bin/bash
# Copyright IBM Corp All Rights Reserved
# SPDX-License-Identifier: Apache-2.0
#####################################################
# modified by Christophe Leske
# v1.0: 16th of feb 2018 cleske@extern.schoenhofer.de
#####################################################
# Exit on first error, print all commands.
set -ev
# don't rewrite paths for Windows Git Bash users
export MSYS_NO_PATHCONV=1

#########################################
# stop any eventually running instances
#########################################
docker-compose -f docker-compose.yml down

#########################################
# start up the network for the blockchain
#########################################
docker-compose -f docker-compose.yml up -d 

#########################################
# due to docker being unreliable, we wait
# three seconds, then start up again to 
# make sure that all containers are running
# NOTE: this cludge is suggested by IBM
#########################################
export FABRIC_START_TIMEOUT=3
sleep ${FABRIC_START_TIMEOUT}

#########################################
# start up again
#########################################
docker-compose -f docker-compose.yml up -d 

#########################################################################################
# Create the channel vertraulich by executing the channel creation transaction channel.tx
#########################################################################################
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@amt1.aufsicht.de/msp" peer1.amt1.aufsicht.de peer channel create -o orderer.aufsicht.de:7050 -c vertraulich -f /etc/hyperledger/configtx/channel.tx

#########################################################################################
# then joins the channel using the information provided in the genesis.block
#########################################################################################
docker exec -e "CORE_PEER_LOCALMSPID=Org1MSP" -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@amt1.aufsicht.de/msp" peer1.amt1.aufsicht.de peer channel join -b vertraulich.block -o orderer.aufsicht.de:7050 --logging-level=notice

#########################################################################################
# peer 2 fetches the genesis.block from the orderer
#########################################################################################
docker exec peer1.amt2.aufsicht.de peer channel fetch config genesis.block -o orderer.aufsicht.de:7050 -c vertraulich
#########################################################################################
# then joins the channel using the information provided in the genesis.block
#########################################################################################
docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@amt2.aufsicht.de/msp" peer1.amt2.aufsicht.de peer channel join -b genesis.block

#########################################################################################
# peer 3 fetches the genesis.block from the orderer
#########################################################################################
docker exec peer1.amt3.aufsicht.de peer channel fetch config genesis.block -o orderer.aufsicht.de:7050 -c vertraulich
#########################################################################################
# then joins the channel using the information provided in the genesis.block
#########################################################################################
docker exec -e "CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/msp/users/Admin@amt3.aufsicht.de/msp" peer1.amt3.aufsicht.de peer channel join -b genesis.block

#########################################################################################
# check to see if all containers are running properly
#########################################################################################
docker ps

docker exec cli buildandinstall.sh
docker exec cli usechaincode.sh
docker exec cli demo.sh