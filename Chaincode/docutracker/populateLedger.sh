// populate ledger
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC1", "Pananama Papers", "1", "ich"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC2", "WikiLeaks", "1", "du"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC3", "BKA", "1", "er"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC4", "Deutsche Bundesbank", "1", "sie"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC5", "Schoenhofer", "1", "es"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC6", "Bitcoin", "1", "es"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC7", "UBS", "1", "es"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC8", "Whitehouse", "1", "es"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC9", "RSA", "1", "es"]}' -C mycc
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOC10", "SPON", "1", "es"]}' -C mycc