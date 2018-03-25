CHANNELNAME=vertraulich &&
CHAINCODENAME=schoenhoferchaincode &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createWorkplace", "workplace1", "Schoenhofer", "Siegburg"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createWorkplace", "workplace2", "LKA NRW", "Duesseldorf"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createWorkplace", "workplace3", "Innenministerium", "Berlin"]}' -C $CHANNELNAME