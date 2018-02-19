for i in {1..10}
do
	echo
	echo '{ "proposal": { "chaincodeId": "schoenhoferchaincode", "fcn": "saveData", "args": [ "DOC'$i'", "' >> _trunk.txt &&
	base64 -w0 Sample_PDFs/PDF_$i.pdf >> _trunk.txt &&
	echo "\"] }}" >> _trunk.txt &&
	echo "Saving binary data for doc"$i""
	curl -s -X POST -H "Content-Type: application/json" -d @_trunk.txt -u 'chris:secret' 'http://fabric-rest:3000/api/fabric/1_0/channels/vertraulich/transactions' &&
	rm _trunk.txt
done
echo
