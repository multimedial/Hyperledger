CHANNELNAME=vertraulich
CHAINCODENAME=schoenhoferchaincode
VERSION=$1
peer chaincode install -p . -n $CHAINCODENAME -v $VERSION && peer chaincode instantiate -n $CHAINCODENAME -v $VERSION -c '{"Args":[""]}' -C $CHANNELNAME
