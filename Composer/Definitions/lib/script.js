/*
* Script file defining the callable methods to interact with the blockchain.
*
* version: v1.0 22nd of Sep 2017
* author: Christophe Leske, cleske@extern.schoenhofer.de
* 
* Schoenhofer Sales and Engineering GmbH, Siegburg, www.schoenhofer.de
*
*/ 

/**
 * Tracks the checkout of a document from one user to another
 * @param {org.proto.schoenhofer.CheckOut} operation - the checkout to process
 * @transaction
 */
function checkoutDocument(operation) {
  	
  	/////////////////////////////////////////////////////////////////
	// do not allow lending if not originalowner is current owner
  	// that means document is currently lent out
  	/////////////////////////////////////////////////////////////////
  	if (operation.document.owner != operation.document.originalowner)
      return False;
  
  	var checkoutNotification = 
        getFactory().newEvent('org.proto.schoenhofer', 'CheckOutNotification');
    	checkoutNotification.document = operation.document;
  		checkoutNotification.from_user = operation.document.owner;
  		checkoutNotification.to_user = operation.newowner;
        emit(checkoutNotification);
  
    operation.document.owner = operation.newOwner;
  
    return getAssetRegistry('org.proto.schoenhofer.Document')
        .then(function (assetRegistry) {
            return assetRegistry.update(operation.document);
        });
}

/**
 * Tracks the return of a document from one user to the originaluser
 * @param {org.proto.schoenhofer.ReturnDocument} operation - the return to process
 * @transaction
 */
function returnDocument(operation) {
  
  	operation.document.owner = operation.document.originalowner;
  
  	// returns the document to its original owner
  	return getAssetRegistry('org.proto.schoenhofer.Document')
        .then(function (assetRegistry) {
            return assetRegistry.update(operation.document);
        });
  }

/**
 * Returns the current owner of document
 * @param {org.proto.schoenhofer.WhoHas} operation - returns current owner
 * @transaction
 */
function whoHas(operation) {
  
    return getAssetRegistry('org.proto.schoenhofer.Document')
        .then(function (assetRegistry) {
      			return(operation.document.owner);
    });
  }