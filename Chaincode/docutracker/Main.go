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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// MAIN APPLICATION ENTRY POINT
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
package Chaincode

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"docutracker/document"
)
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// main function starts up the chaincode in the container during instantiation
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func main () {

	if err := shim.Start(new(document.Document)); err != nil {
		fmt.Printf("Error starting Document chaincode: %s", err)
	}
}