CHANNELNAME=vertraulich &&
CHAINCODENAME=schoenhoferchaincode &&
sleep 2 &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user1", "Chad", "Palaniuk", "workplace1", "0"]}' -C $CHANNELNAME && \
sleep 1 &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user2", "Wolfgang", "Petry", "workplace1", "0"]}' -C $CHANNELNAME && \
sleep 1 && \
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user3", "Dude", "Lebowski", "workplace1", "0"]}' -C $CHANNELNAME && \
sleep 1 && \
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user4", "Samuel", "Jackson", "workplace2", "1"]}' -C $CHANNELNAME && \
sleep 1 && \
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user5", "Wolf", "Gang", "workplace2", "1"]}' -C $CHANNELNAME && \
sleep 1 && \
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user6", "Jim", "Pansen", "workplace2", "1"]}' -C $CHANNELNAME && \
sleep 1 && \
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user7", "Bratislav", "Methulski", "workplace3", "2"]}' -C $CHANNELNAME && \
sleep 1 && \
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user8", "Arvo", "Paert", "workplace3", "2"]}' -C $CHANNELNAME && \
sleep 1 && \
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user9", "Towa", "Tei", "workplace3", "3"]}' -C $CHANNELNAME