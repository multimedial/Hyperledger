CHANNELNAME=vertraulich
CHAINCODENAME=schoenhoferchaincode
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createWorkplace", "workplace1", "Schoenhofer", "Siegburg"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createWorkplace", "workplace2", "LKA Nordrhein-Westfalen", "Duesseldorf"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createWorkplace", "workplace3", "Innenministerium", "Berlin"]}' -C $CHANNELNAME &&
sleep 2 &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user0", "Chad", "Palaniuk", "workplace1", "0"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user1", "Wolfgang", "Petry", "workplace1", "0"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user2", "Dude", "Lebowski", "workplace1", "0"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user3", "Samuel", "Jackson", "workplace2", "1"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user4", "Wolf", "Gang", "workplace2", "1"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user5", "Jim", "Pansen", "workplace2", "1"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user6", "Bratislav", "Methulski", "workplace3", "2"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user7", "Arvo", "Paert", "workplace3", "2"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createUser", "user8", "Towa", "Tei", "workplace3", "2"]}' -C $CHANNELNAME &&
sleep 2 &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createDocument", "DOC1", "Pananama Papers", "1", "user0", "0"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createDocument", "DOC2", "WikiLeaks", "1", "user1","0"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createDocument", "DOC3", "BKA", "1", "user2", "1"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createDocument", "DOC4", "Deutsche Bundesbank", "1", "user3", "1"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createDocument", "DOC5", "Schoenhofer", "1", "user4", "2"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createDocument", "DOC6", "Bitcoin", "1", "user5", "2"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createDocument", "DOC7", "UBS", "1", "user6", "2"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createDocument", "DOC8", "Whitehouse", "1", "user7", "3"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createDocument", "DOC9", "RSA", "1", "user8", "3"]}' -C $CHANNELNAME &&
peer chaincode invoke -n $CHAINCODENAME -c '{"Args":["createDocument", "DOC10", "SPON", "1", "user8", "3"]}' -C $CHANNELNAME