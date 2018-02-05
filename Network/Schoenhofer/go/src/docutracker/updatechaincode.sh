CHANNELNAME=vertraulich
CHAINCODENAME=schoenhoferchaincode
VERSION=$1
peer chaincode upgrade -p . -C $CHANNELNAME -n $CHAINCODENAME -v $VERSION -c '{"Args":[]}'
