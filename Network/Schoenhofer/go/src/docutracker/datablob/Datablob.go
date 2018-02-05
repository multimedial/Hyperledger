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
package datablob
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// models a trackable document for the document tracking system
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
import (
"fmt"
_ "strconv"
"github.com/hyperledger/fabric/core/chaincode/shim"
_ "github.com/hyperledger/fabric/protos/peer"
_ "encoding/json"
_ "bytes"
)
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Document implements a document asset owned by a user
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type Datablob struct {
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// to be JSONED
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	DocID string
	Data string
}

func Get(stub shim.ChaincodeStubInterface, args []string) (string, error){

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

func Set(stub shim.ChaincodeStubInterface, args []string) (string, error) {

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
