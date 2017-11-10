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

package amt

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// Document implements a document asset owned by a user
type Document struct {

	owner docutrackeruser.User
	title string = "No Name"
	version int = 0

}

// init is called during chaincode instantiation to initialize any data.
func (t *Document) Init(stub shim.ChaincodeStubInterface) peer.Response {

	// get Args from transaction protocol
	args := stub.GetStringArgs()
	if len(args) != 2 {
		return shim.Error("Incorrect arguments. Expecting a key and a value!")
	}

	// set up any vars or assets here by calling stub.PutState()
	// store the key and value on the ledger
	err := stub.PutState(args[0], []byte(args[1]))

	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
	}

	return shim.Success(nil)

}


// invoke is called per transaction on the chaincode. Each transaction is either a 'get' or 'set' on the asset created
// by the Init function. The 'set' method may create a new asset by specifying a new key-value pair.
func (t *Document) Invoke (stub shim.ChaincodeStubInterface) peer.Response {

	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error

	if fn == "getData" {
		result, err = getData(stub, args)
	}

	if fn == "set" {
		result, err = set(stub, args)
	} else {
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

func getData(stub shim.ChaincodeStubInterface, args []string) (string, error) {

	key := args[0]
	result, err := stub.GetQueryResult(key)
	return string(result), err

}

// main function starts up the chaincode in the container during instantiation
func main () {

	if err := shim.Start(new(Document)); err != nil {
		fmt.Printf("Error starting Document chaincode: %s", err)
	}

}