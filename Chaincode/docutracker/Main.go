////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Document tracking system POC for
// Schoenhofer Sales and Engineering GmbH, Siegburg
//
// 	author: christophe leske, christophe.leske@extern.schoenhofer.de
//	v: 1.0 (10th of nov 2017)
//
// Requirements:
// *************
// requires the Hyperledger fabric binaries, tools (docker images,cryptogen, and others),
// to be downloaded by using cURL: curl -sSL https://goo.gl/5ftp2f | bash
// (see https://hyperledger-fabric.readthedocs.io/en/release/samples.html#binaries)
//
// as well as the fabric samples (https://hyperledger-fabric.readthedocs.io/en/release/samples.html)
// git clone https://github.com/hyperledger/fabric-samples.git
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
package main
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// models a trackable document for the document tracking system
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
import (
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"encoding/json"
	"bytes"
	"docutracker/document"
	"docutracker/docuser"
)
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Document implements a document asset owned by a user
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type SmartContract struct {

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// init is called during chaincode instantiation to initialize any data.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("######################## Document struct initialized. ########################")

	/*
	args := stub.GetStringArgs()
	if len(args) != 3 {
		return shim.Error("Incorrect arguments. Expecting a key and a value")
	}

	// Set up any variables or assets here by calling stub.PutState()
	// We store the key and the value on the ledger

	err := stub.PutState("title", []byte(args[0]))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}

	err = stub.PutState("version", []byte(args[1]))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[1]))
	}

	err = stub.PutState("owner", []byte(args[2]))
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[2]))
	}
*/

	return shim.Success(nil)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// invoke is called per transaction on the chaincode. Each transaction is either a 'get' or 'set' on the asset created
// by the Init function. The 'set' method may create a new asset by specifying a new key-value pair.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (t *SmartContract) Invoke (stub shim.ChaincodeStubInterface) peer.Response {

	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error

	if fn == "createDocument" {
		return t.createDocument(stub, args)
	}

	if fn == "createUser" {
		return t.createUser(stub, args)
	}

	if fn == "lendDocument" {
		return t.lendDocument(stub, args)
	}

	if fn == "returnDocument" {
		return t.returnDocument(stub, args)
	}

	if fn == "queryAllDocs" {
		return t.queryAllDocs(stub)
	}

	if fn == "queryAllUser" {
		return t.queryAllUser(stub)
	}

	if fn == "set" {
		result, err = set(stub, args)
	}

	if fn == "get" {
		result, err = get(stub, args)
	}

	if err != nil {
		return shim.Error(err.Error())
	}

	// return the result as success payload
	return shim.Success([]byte(result))

}

func get(stub shim.ChaincodeStubInterface, args []string) (string, error){

	if len(args) != 1 {
		return "", fmt.Errorf("Inccorrect arguments. Expecting a key.")
	}

	// ok, we got ONE args, fetch it
	value, err := stub.GetState(args[0])

	if err != nil {
		return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err)
	}

	if value == nil {
		return "", fmt.Errorf("Asset not found: %s", args[0])
	}

	return string(value), nil
}

func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {

	if len(args) != 2 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value.")
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// still here? so we got 2 args (key, value pair) -> put them onto the ledger!
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	err := stub.PutState(args[0], []byte(args[1]))

	if err != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}

	// everything went fine
	return args[1], nil

}

func (s *SmartContract) createDocument(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	 // check arguments provided - we need 5
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5.")
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// store the identityKey (first argument) as this represents the ID of the document
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	identityKey := args[0]

	title := args[1]
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// set the version of the document
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	version,_ := strconv.Atoi(args[2])

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// get the security Level of the document and bail if there is an error
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	securityLevel,err := strconv.Atoi(args[4])
	if err != nil || securityLevel<0 || securityLevel>3 {
		// there was an error converting the securityLevel
		return shim.Error("Either could not convert security level, or level provided is out of range (must be 0 to 3).")
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// check the owner of the document - it must be the username of an already registered one
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	owner := args[3]
	if !s.isUser(stub, owner) {
		return shim.Error("The owner of the document is not a registered user. Please register the user first!")
	}

	var doc = document.Document{Title: title, Version: version, Owner: owner, CurrentOwner: owner, SecurityLevel: securityLevel}
	documentAsBytes, _ := json.Marshal(doc)
	stub.PutState(identityKey, documentAsBytes)
	return shim.Success(nil)
}

func (s *SmartContract) createUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 5.")
	}

	// args[0] => KEY for multirange retrieval
	userid := args[0]

	securityLevel, err := strconv.Atoi(args[3])

	if err != nil {
		return shim.Error("Incorrect security level passed. Must be an integer.")
	}

	var usr = docuser.User{FirstName: args[1], LastName: args[2], SecurityLevel: securityLevel}
	var usrAsBytes,_ = json.Marshal(usr)
	stub.PutState(userid, usrAsBytes)
	return shim.Success(nil)
}

func (s *SmartContract) queryAllDocs(stub shim.ChaincodeStubInterface) peer.Response {

	startKey := "DOC0"
	endKey := "DOC99999"

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllDocs:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryAllUser(stub shim.ChaincodeStubInterface) peer.Response {

	startKey := "user0"
	endKey := "user99999"

	resultsIterator, err := stub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllUser:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) isUser(stub shim.ChaincodeStubInterface, userid string) bool {
	// quick test to see if a given user is a registered one
	owner,_ := stub.GetState(userid)
	return owner != nil
}

func (s *SmartContract) lendDocument(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	// we only give a document out IF...
	// ... the current owner is the original owner
	// ... and the security level of the new owner is same or higher than security level of doc

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2: id of document, id of new user")
	}

	docid := args[0]
	newOwnerid := args[1]

	////////////////////////////////////////////////////////////
	// check if documentid does exist
	////////////////////////////////////////////////////////////
	docAsBytes,_ := stub.GetState(docid)
	if docAsBytes == nil {
		return shim.Error("Document does not exist.")
	}

	////////////////////////////////////////////////////////////
	// check if user does exist
	////////////////////////////////////////////////////////////
	usrAsBytes,_ := stub.GetState(newOwnerid)
	if usrAsBytes == nil {
		return shim.Error("User does not exist.")
	}

	///////////////////////////////////////////////////////////////////
	// check if security level of new user is same or higher as of doc
	///////////////////////////////////////////////////////////////////
	var doc document.Document
	err := json.Unmarshal(docAsBytes, &doc)
	if err != nil {
		return shim.Error("Error while unmarshalling document from json represention.")
	}

	var usr docuser.User
	err = json.Unmarshal(usrAsBytes, &usr)
	if err != nil {
		return shim.Error("Error while unmarshalling user from json represention.")
	}

	// check if we can actually lend the document (CurrentOwner == Owner)
	if doc.CurrentOwner != doc.Owner {
		return shim.Error("The document is checked out and cannot be lent at the moment!")
	}

	// now check if security levels are same
	if usr.SecurityLevel >= doc.SecurityLevel {
		// we are ok, change it!
		doc.CurrentOwner = newOwnerid
		docAsBytes, _ = json.Marshal(doc)
		stub.PutState(docid, docAsBytes)
	} else {
		return shim.Error("Security Levels are not compatible.")
	}

	return shim.Success(nil)

}

func (s *SmartContract) returnDocument(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2: id of document to return, and userid of returner.")
	}

	docid := args[0]
	returningUser := args[1]

	docAsBytes,err := stub.GetState(docid)
	var doc document.Document
	err = json.Unmarshal(docAsBytes, &doc)
	if err!=nil {
		return shim.Error("There was an error unmarshalling the document.")
	}

	// sanity check: who brings back the book?
	if doc.CurrentOwner != returningUser {
		// this is odd, someone brings the book back that didn't lend it out!
		// TBA: Event
		return shim.Error("WARNING: someone else brought the document back!")
	}

	if doc.Owner == doc.CurrentOwner {
		// error - this document cannot be brought back,
		// as it is not lend out!
		return shim.Error("The document cannot be given back: current owner is the owner.")
	}

	// ok, we bring back the document
	doc.CurrentOwner = doc.Owner
	docAsBytes, err = json.Marshal(doc)
	if err != nil {
		return shim.Error("Something went wrong while marshalling the document.")
	}
	stub.PutState(docid,docAsBytes)
	return shim.Success(nil)

}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// main function starts up the chaincode in the container during instantiation
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func main () {

	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Error starting Document chaincode: %s", err)
	}

}