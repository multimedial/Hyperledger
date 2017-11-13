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
)
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Document implements a document asset owned by a user
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
type SmartContract struct {
}

type Document struct {
	Title string
	Version int
	Owner string
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

	if fn == "queryAllDocs" {
		return t.queryAllDocs(stub)
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

	// args[0] => KEY for multirange retrieval

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	version,_ := strconv.Atoi(args[2])
	var doc = Document{Title: args[1], Version: version, Owner: args[3]}

	documentAsBytes, _ := json.Marshal(doc)
	stub.PutState(args[0], documentAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllDocs(stub shim.ChaincodeStubInterface) peer.Response {

	startKey := "DOCUMENT0"
	endKey := "DOCUMENT999"

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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// main function starts up the chaincode in the container during instantiation
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func main () {

	if err := shim.Start(new(SmartContract)); err != nil {
		fmt.Printf("Error starting Document chaincode: %s", err)
	}

}