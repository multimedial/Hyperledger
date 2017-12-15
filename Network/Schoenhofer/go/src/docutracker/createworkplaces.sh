peer chaincode invoke -n mycc -c '{"Args":["createWorkplace", "workplace1", "Schoenhofer", "Siegburg"]}' -C vertraulich &&
peer chaincode invoke -n mycc -c '{"Args":["createWorkplace", "workplace2", "LKA Nordrhein-Westfalen", "Duesseldorf"]}' -C vertraulich &&
peer chaincode invoke -n mycc -c '{"Args":["createWorkplace", "workplace3", "Innenministerium", "Berlin"]}' -C vertraulich