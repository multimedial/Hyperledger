CHANNELNAME=vertraulich && 
CHAINCODENAME=schoenhoferchaincode &&
PDFDATA="$(cat Sample_PDFs/PDF_1.pdf|base64 -w 0)" &&
peer chaincode invoke -n $CHAINCODENAME -C $CHANNELNAME -c '{"Args":["saveData", "DOC1", "'$PDFDATA'"]}'