peer chaincode invoke -n mycc -c '{"Args":["createUser", "user0", "Chad", "Palaniuk", "0"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user1", "Wolfgang", "Petry", "0"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user2", "Dude", "Lebowski", "0"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user3", "Samuel", "Jackson", "1"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user4", "Wolf", "Gang", "1"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user5", "Jim", "Pansen", "1"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user6", "Bratislav", "Methulski", "2"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user7", "Arvo", "Paert", "2"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user8", "Towa", "Tei", "2"]}' -C myc