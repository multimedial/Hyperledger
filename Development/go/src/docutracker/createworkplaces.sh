peer chaincode invoke -n mycc -c '{"Args":["createWorkplace", "workplace1", "Schoenhofer", "Siegburg"]}' -C myc &&
peer chaincode invoke -n mycc -c '{"Args":["createWorkplace", "workplace2", "LKA Nordrhein-Westfalen", "Duesseldorf"]}' -C myc &&
peer chaincode invoke -n mycc -c '{"Args":["createWorkplace", "workplace3", "Innenministerium", "Berlin"]}' -C myc