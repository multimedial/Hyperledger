////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Document tracking system POC for Schoenhofer Sales and Engineering GmbH, Siegburg
//
// author: christophe leske, christophe.leske@extern.schoenhofer.de
// v: 1.0 (10th of nov 2017)
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

	////////////////////////////////////////////////////
	// standard packages
	////////////////////////////////////////////////////
	"fmt"
	"strconv"
	"encoding/json"
	"bytes"
	"time"
	////////////////////////////////////////////////////
	// external packages
	////////////////////////////////////////////////////
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"

	////////////////////////////////////////////////////
	// project-specific sub-packages
	////////////////////////////////////////////////////
	"docutracker/document"
	"docutracker/docuser"
	"docutracker/workplace"
	_"docutracker/datablob"
	_"strings"
)

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Document implements a document asset owned by a user
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type SmartContract struct {
	// this is the smart contract struct
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// init is called during chaincode instantiation to initialize any data.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (t *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {

	fmt.Println("######################## SmartContract struct initialized. ########################")
	return shim.Success(nil)

}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// invoke is called per transaction on the chaincode. Each transaction is either a 'get' or 'set' on the asset created
// by the Init function. The 'set' method may create a new asset by specifying a new key-value pair.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func (t *SmartContract) Invoke (stub shim.ChaincodeStubInterface) peer.Response {

	fn, args := stub.GetFunctionAndParameters()

	// some logging
	fmt.Println()
	fmt.Println("REQUESTED METHOD CALL: " + fn)
	fmt.Println()

	// some variables for later use
	var result string
	var err error

	//////////////////////////////////////////////////////////////////
	// main branching of functions
	//////////////////////////////////////////////////////////////////
	if fn == "createDocument" {
		return t.createDocument(stub, args)
	}

	if fn == "createUser" {
		return t.createUser(stub, args)
	}

	if fn == "createWorkplace" {
		return t.createWorkplace(stub, args)
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

	if fn == "queryAllWorkplaces" {
		return t.queryAllWorkplaces(stub)
	}

	if fn == "queryAll" {
		return t.queryAll(stub, args[0]) // only first argument passed
	}

	if fn == "getHistory" {
		return t.getHistory(stub, args[0]) // only first argument passed
	}

	if fn == "saveData" {
		return t.saveData(stub, args)
	}

	if fn == "getData" {
		return t.getData(stub, args)
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

	//////////////////////////////////////////////////////////////////
	// return the result as success payload
	//////////////////////////////////////////////////////////////////
	if result != "" {
		return shim.Success([]byte(result))
	} else {
		return shim.Error("Method call not recognized: '" + fn + "'")
	}

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
	// store the docid (first argument) as this represents the ID of the document
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	docid := args[0]

	fmt.Println("Creating new document with id " + args[0])

	title := args[1]
	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// set the version of the document
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	version,_ := strconv.Atoi(args[2])

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// check the owner of the document - it must be the username of an already registered one
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	owner := args[3]
	if (s.isUser(stub, owner)==false) {
		return shim.Error("The owner of the document is not a registered user. Please register the user first: " +owner)
	}

	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// get the security Level of the document and bail if there is an error
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	securityLevel,err := strconv.Atoi(args[4])
	if err != nil || securityLevel<0 || securityLevel>3 {
		// there was an error converting the securityLevel
		return shim.Error("Either could not convert security level, or level provided is out of range (must be 0 to 3).")
	}

	var doc = document.Document{Title: title, Version: version, Owner: owner, CurrentOwner: "", SecurityLevel: securityLevel}
	documentAsBytes, _ := json.Marshal(doc)
	stub.PutState(docid, documentAsBytes)
	return shim.Success(nil)

}

func (s *SmartContract) saveData(stub shim.ChaincodeStubInterface, args[] string) peer.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2: docid, data")
	}

	///////////////////////////////////////////////////
	// 1.st param: docid - document key for the search in the database
	///////////////////////////////////////////////////
	docid := args[0]
	data_base64encoded := args[1]

	////////////////////////////////////////////////////////////
	// check if documentid does exist and fetch object
	////////////////////////////////////////////////////////////
	docAsBytes,_ := stub.GetState(docid)
	if docAsBytes == nil {
		return shim.Error("Document does not exist.")
	}

	////////////////////////////////////////////////////////////
	// cast object from binary representation
	////////////////////////////////////////////////////////////
	var doc document.Document
	err := json.Unmarshal(docAsBytes, &doc)
	if err != nil {
		return shim.Error("Error while unmarshalling document from json represention.")
	}

	////////////////////////////////////////////////////////////
	// assign the data we received
	////////////////////////////////////////////////////////////
	blobid := "blob_"+docid
	err = stub.PutState(blobid, []byte(data_base64encoded))
	doc.DatablobID = blobid

	if err != nil {
		return shim.Error("Error while saving blob data.")
	}

	return shim.Success(nil)

}

func (s *SmartContract) getData(stub shim.ChaincodeStubInterface, args[] string) peer.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2: docid, userid")
	}

	docid := args[0]
	docAsBytes,err := stub.GetState(docid)
	if err != nil {
		return shim.Error("Error while getting the document from the ledger: " + docid)
	}

	userid := args[1]
	userAsBytes,_ := stub.GetState(userid)
	var user docuser.User
	if json.Unmarshal(userAsBytes, &user) != nil {
		return shim.Error("Error while unmarshalling user from json represention, the user probably doesn't exist: " + userid)
	}

	var doc document.Document
	if json.Unmarshal(docAsBytes, &doc) != nil {
		return shim.Error("Error while unmarshalling document from json represention.")
	}

	if user.SecurityLevel < doc.SecurityLevel {
		return shim.Error("The user is not allowed to get this document due to his or her security level: " + string(user.SecurityLevel) + " " + string(doc.SecurityLevel))
	}

	dataAsBytes,err := stub.GetState("blob_"+docid)

	if err != nil {
		return shim.Error("Error while getting the data for document: " + docid)
	}

	//
	return shim.Success( dataAsBytes )
}

func (s *SmartContract) createUser(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5: userid, firstname, lastname, workplace, securitylevel")
	}

	///////////////////////////////////////////////////
	// 1.st param: userid - key for the search in the database
	///////////////////////////////////////////////////
	userid := args[0]

	///////////////////////////////////////////////////
	// 2.nd param: firstname
	///////////////////////////////////////////////////
	firstname := args[1]

	///////////////////////////////////////////////////
	// 3.rd param: lastname
	///////////////////////////////////////////////////
	lastname := args[2]

	///////////////////////////////////////////////////
	// 4th param: workplace ID (!)
	///////////////////////////////////////////////////
	workplace := args[3]
	///////////////////////////////////////////////////
	// we need to check if the supplied workplaceID is
	// indeed a valid workplace! If not, webail out
	///////////////////////////////////////////////////
	if s.isWorkplace(stub, workplace)==false {
		return shim.Error("Incorrect workplace id. Must be a valid workplace id.")
	}

	///////////////////////////////////////////////////
	// 5th param: securitylevel, provided as a string
	// needs to be recasted to an integer
	///////////////////////////////////////////////////
	securityLevel, err := strconv.Atoi(args[4])
	if err != nil {
		return shim.Error("Incorrect security level passed. Must be an integer.")
	}

	//////////////////////////////////////////////////////////////////////////////////////////////////////
	// if we are here, everything seems to be fine, create new user object with infos supplied
	//////////////////////////////////////////////////////////////////////////////////////////////////////
	fmt.Println("Creating new user with id " + userid)
	var usr = docuser.User{FirstName: firstname, LastName: lastname, WorkplaceID: workplace, SecurityLevel: securityLevel}

	///////////////////////////////////////////////////
	// convert it to a JSON representation
	///////////////////////////////////////////////////
	var usrAsBytes,_ = json.Marshal(usr)

	///////////////////////////////////////////////////
	// put into the blockchain, use userid as key
	///////////////////////////////////////////////////
	stub.PutState(userid, usrAsBytes)

	///////////////////////////////////////////////////
	// return success
	///////////////////////////////////////////////////
	return shim.Success(usrAsBytes)
}

func (s *SmartContract) createWorkplace(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	fmt.Println("Creating new workplace with id " + args[0])
	/////////////////////////////////////////////////////////////
	// creates a new workplace object and stores it in the ledger
	/////////////////////////////////////////////////////////////
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 2: id, Name, Place")
	}

	id := args[0]
	name := args[1]
	place := args[2]

	var workplace = workplace.Workplace{Name: name, Place: place}

	/////////////////////////////////////////////////////////////
	// convert to JSON representation
	/////////////////////////////////////////////////////////////
	var workplaceAsBytes,err = json.Marshal(workplace)

	if err != nil {
		return shim.Error("Error during JSON conversion of workplace object.")
	}

	// store it
	stub.PutState(id, workplaceAsBytes)

	// return success!
	return shim.Success(workplaceAsBytes)

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
		if bArrayMemberAlreadyWritten {
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

	fmt.Printf("- queryAllDocs:\n%s\n", buffer.String()[0:255])

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

func (s *SmartContract) queryAllWorkplaces(stub shim.ChaincodeStubInterface) peer.Response {

	startKey := "workplace0"
	endKey := "workplace99999"

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

	fmt.Printf("- queryAllWorkplaces:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) queryAll(stub shim.ChaincodeStubInterface, key string) peer.Response {

	startKey := key+"0"
	endKey := key+"99999"

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

	fmt.Printf("- queryAll"+key+":\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) isUser(stub shim.ChaincodeStubInterface, userid string) bool {
	// quick test to see if a given user is a registered one
	_, err := stub.GetState(userid)
	return err==nil
}

func (s *SmartContract) isWorkplace(stub shim.ChaincodeStubInterface, workplace string) bool {
	// quick test to see if a given workplace is valid
	_,err := stub.GetState(workplace)
	return err==nil
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
	if json.Unmarshal(docAsBytes, &doc) != nil {
		return shim.Error("Error while unmarshalling document from json represention.")
	}

	var usr docuser.User
	if json.Unmarshal(usrAsBytes, &usr) != nil {
		return shim.Error("Error while unmarshalling user from json represention.")
	}

	// now check if security levels are same
	if usr.SecurityLevel >= doc.SecurityLevel {
		// we are ok, change it!
		fmt.Println("Security level match for user " + newOwnerid + " [level "+ strconv.Itoa(usr.SecurityLevel) +"] and document " + docid + " [level " + strconv.Itoa(doc.SecurityLevel) + "]: OK")
		// check if we can actually lend the document (CurrentOwner == Owner)
		if doc.CurrentOwner !="" {
			/////////////////////////////////////////////////
			// emit an event
			/////////////////////////////////////////////////
			fmt.Println("Document " + docid + "(level " + strconv.Itoa(doc.SecurityLevel) + ") is not available for lending out!")
			stub.SetEvent("Event_LendOut_Warning", []byte("Request for " + docid + " by " + newOwnerid + ", but " + doc.CurrentOwner + " has it already."))
			return shim.Success([]byte("ERROR: document " + docid + " is already checked out by " + doc.CurrentOwner +  " and not available at the moment."))
		} else {
			doc.CurrentOwner = newOwnerid
			docAsBytes, _ = json.Marshal(doc)
			stub.PutState(docid, docAsBytes)
			fmt.Println("Lending out successful!")
			return shim.Success([]byte("SUCCESS: Lending out successful."))
		}
	} else {
		fmt.Println()
		fmt.Println("####################### SECURITY WARNING #######################")
		fmt.Println(newOwnerid + " (security level:" + strconv.Itoa(usr.SecurityLevel) + ") tried to access document " + docid + " (security level:" + strconv.Itoa(doc.SecurityLevel)+")")
		fmt.Println("################################################################")
		fmt.Println()
		errorMsg := "User "+newOwnerid+" has not the required security level for " + docid + " which has security level " + strconv.Itoa(usr.SecurityLevel)
		stub.SetEvent("Event_Security_Error", []byte(errorMsg))
		return shim.Success([]byte("ERROR: Security Levels are not compatible."))
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
		fmt.Println()
		fmt.Println("####################### SECURITY ERRROR! #######################")
		fmt.Println(returningUser + " brought back document " + docid + " although it should have been " + doc.CurrentOwner)
		fmt.Println("################################################################")
		fmt.Println()
		stub.SetEvent("Event_Security_Error", []byte("Mitarbeiter " + returningUser + " versuchte Dokument " + docid + " zurückzubringen. Dokument war ausgeliehen von " + doc.CurrentOwner + "."))
		return shim.Success([]byte("ERROR: someone else brought the document back!"))
	}

	if doc.Owner == doc.CurrentOwner {
		// error - this document cannot be brought back,
		// as it is not lend out!
		return shim.Success([]byte("The document cannot be given back: current owner is the owner."))
	}

	// ok, we bring back the document
	doc.CurrentOwner = ""
	docAsBytes, err = json.Marshal(doc)
	if err != nil {
		return shim.Error("Something went wrong while marshalling the document.")
	}
	stub.PutState(docid,docAsBytes)
	return shim.Success([]byte("Das Dokument wurde erfolgreich zurückgegeben und kann wieder ausgeliehen werden."))

}

func (s *SmartContract) getHistory(stub shim.ChaincodeStubInterface, key string) peer.Response {

	historyIer, err := stub.GetHistoryForKey(key)

	if err != nil {
		fmt.Println("ERROR while reading the ledger!")
		return shim.Error("ERROR while reading the ledger!")
	}
	fmt.Println()
	fmt.Println("Returning ledger history for object '" + key + "':")

	var buffer bytes.Buffer

	for historyIer.HasNext() {
		modification, err := historyIer.Next()
		if err != nil {
			fmt.Println("ERROR while reading an entry in the ledger.")
			return shim.Error("ERROR while reading an entry in the ledger.")
		}

		buffer.WriteString("TX_ID: " + modification.TxId + "\n")
		buffer.WriteString(time.Unix(modification.Timestamp.Seconds, 0).String()+"\n")
		buffer.WriteString("VALUES: " + string(modification.Value))
		buffer.WriteString("\n\n")
	}
	fmt.Println(buffer.String())

	return shim.Success(buffer.Bytes())
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// main function starts up the chaincode in the container during instantiation
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func main () {

	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Error starting Document chaincode: %s", err)
	}

	shim.SetLoggingLevel(shim.LogWarning)

}