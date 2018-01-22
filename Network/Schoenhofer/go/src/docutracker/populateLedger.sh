peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC1", "Pananama Papers", "1", "user0", "0"]}' -C vertraulich &&
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC2", "WikiLeaks", "1", "user1","0"]}' -C vertraulich &&
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC3", "BKA", "1", "user2", "1"]}' -C vertraulich &&
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC4", "Deutsche Bundesbank", "1", "user3", "1"]}' -C vertraulich &&
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC5", "Schoenhofer", "1", "user4", "2"]}' -C vertraulich &&
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC6", "Bitcoin", "1", "user5", "2"]}' -C vertraulich &&
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC7", "UBS", "1", "user6", "2"]}' -C vertraulich &&
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC8", "Whitehouse", "1", "user7", "3"]}' -C vertraulich &&
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC9", "RSA", "1", "user8", "3"]}' -C vertraulich &&
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC10", "SPON", "1", "user8", "3"]}' -C vertraulich