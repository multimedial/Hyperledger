CHANNELNAME=vertraulich
CHAINCODENAME=schoenhoferchaincode
VERSION=$1
go build && CORE_CHAINCODE_ID_NAME=$CHAINCODENAME:$VERSION ./docutracker