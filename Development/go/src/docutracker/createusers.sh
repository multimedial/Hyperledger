peer chaincode invoke -n mycc -c '{"Args":["createUser", "user0", "Chad", "Palaniuk", "workplace1", "0"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user1", "Wolfgang", "Petry", "workplace1", "0"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user2", "Dude", "Lebowski", "workplace1", "0"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user3", "Samuel", "Jackson", "workplace2", "1"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user4", "Wolf", "Gang", "workplace2", "1"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user5", "Jim", "Pansen", "workplace2", "1"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user6", "Bratislav", "Methulski", "workplace3", "2"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user7", "Arvo", "Paert", "workplace3", "2"]}' -C myc && \
sleep 1 && \
peer chaincode invoke -n mycc -c '{"Args":["createUser", "user8", "Towa", "Tei", "workplace3", "2"]}' -C myc