CHAINCODEPATH=.
CHAINCODENAME=schoenhoferchaincode
CHAINCODEVERSION=1
CHANNELNAME=vertraulich
peer chaincode install -p $CHAINCODEPATH -n $CHAINCODENAME -v $CHAINCODEVERSION && peer chaincode instantiate -n $CHAINCODENAME -v $CHAINCODEVERSION -c '{"Args":[]}' -C $CHANNELNAME