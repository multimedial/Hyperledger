peer chaincode install -p chaincodedev/chaincode/document -n mycc -v 0
peer chaincode instantiate -n mycc -v 0 -c '{"Args":["title","abc"]}' -C myc
// populate ledger
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOCUMENT0", "VS1", "1", "ich"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOCUMENT1", "VS2", "1", "du"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOCUMENT2", "VS3", "1", "er"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOCUMENT3", "VS4", "1", "sie"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["createDocument", "DOCUMENT4", "VS5", "1", "es"]}' -C myc
peer chaincode invoke -n mycc -c '{"Args":["queryAllDocs"]}' -C myc