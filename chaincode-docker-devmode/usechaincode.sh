peer chaincode install -p . -n mycc -v 0 && peer chaincode instantiate -n mycc -v 0 -c '{"Args":["title","abc"]}' -C myc
