Prerequisites:
**************
- für Hyperledger, siehe https://hyperledger-fabric.readthedocs.io/en/release/prereqs.html, insbesondere Docker und Go-Umgebung

- neueste Hyperledger Binaires und Docker images, zu beziehen per
	curl -sSL https://goo.gl/byy2Qj | bash -s 1.0.5
	(Quelle: https://hyperledger-fabric.readthedocs.io/en/release/samples.html#binaries)
	
- Dateien aus dem Repository (https://github.com/multimedial/Hyperledger)

- mysql Image, zu beziehen per "docker pull mysql/mysql-server"

- node.js sowie npm 

- für den Blockchain Explorer, siehe 
https://github.com/hyperledger/blockchain-explorer#requirements





Schritte zum Aufbau des Demo-Netzwerkes:
****************************************
Der Ablauf ist:
- Aufbau der Infrastruktur
- Vorbereitung und Ausführung des Chaincodes
- Visualisierung
- Ausführung von Aufrufen

Schritt 1: Aufbau der Infrastruktur

Ins Verzeichnis "Hyperledger/Network/blockchain-explorer/Schoenhofer" wechseln.

Dort:

	./start.sh
	
Dies startet die Container und erstellt somit die Infrastruktur der Demo:
drei Peers, drei CAs, einen Orderer, einen CouchDB Container sowie einen mysql-Container.

Zur Kontroller per docker ps feststellen, dass auch alle Container gestartet wurden.


[Einschub: 

hier muss der mysql-Container gestartet und mit der fabricexplorer.sql Datei beschickt werden, damit die Tabellen für den Blockchain-Viewer erstellt werden.

Ich mache dies noch händisch im Moment: 

	Starten des MySQL Servers:
	
		docker run -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_ROOT_HOST=% -p 3306:3306 mysql/mysql-server

	In einem anderen Terminalfenster per 
	
		docker ps 
	
	den Namen des SQL Server Containers ermitteln (oder beim Aufruf einen eigenen vergeben), dann...

		docker exec -it <NAME DES SQL CONTAINERS> mysql -u root -p

	Eingabe des Passworts ("123456"), dann pasten des Inhaltes von "Hyperledger/Network/blockchain-explorer/fabricexplorer.sql"

]


Schritt 2: Vorbereitung und Ausführung des Chaincodes

In den CLI-Container der Blockchain-Infrastruktur wechseln:

	docker exec -it cli bash

Ins Verzeichnis "/opt/gopath/src/docutracker" wechseln per
	
	cd /opt/gopath/src/docutracker
	
Dort dann ausführen:

	./buildandinstall.sh
	
Der Chaincode wird dann gestartet, letzte Zeile sollte sein " [...] starting up ... "

In einem neuen Terminal dann wieder in den CLI-Container wechseln und den Code instanziieren und starten mit:

	docker exec -it cli bash
	
	cd /opt/gopath/src/docutracker
	
	./startcode.sh
	
Ergebniss sollte ohne Fehler sein, und im vorherigen Terminalfenster sollte stehen "#### Smartcontract struct initialized #####"



Schritt 3: Visualisierung

Folgendes in einem nativen Terminal eingeben,also in keinem der Docker Container. 

Wechseln in 

	"Hyperledger/Network/blockchain-explorer/"

ACHTUNG: da dies ein separates Projekt ist, muss es vor dem ersten Aufruf gebaut werden mit 

	npm install
	
	dann

	./monitor.sh
	
dann im Browser http://localhost:8080 aufrufen.

Dort sollte dann der Blockchain-Viewer zu sehen sein mit den Peers und mindestens einem Block in der Chain.



Zur Demo: 

In den CLI Container wechseln, 

	docker exec -it cli bash

und dann 

	./demo.sh
	
Dies füllt die Blockchain mit Transaktionen und Objekten (Workplaces, User, Dokumente). Dies sollte im Browser zu sehen sein.



