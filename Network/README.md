# Schönhofer Document Tracking POC

## Technische Voraussetzungen
**************
- für Hyperledger:
	- siehe https://hyperledger-fabric.readthedocs.io/en/release/prereqs.html, insbesondere Docker und die Go-Umgebung
- neueste Hyperledger Binaires und Docker images, zu beziehen per 
	> curl -sSL https://goo.gl/byy2Qj | bash -s 1.0.5
[Quelle](https://hyperledger-fabric.readthedocs.io/en/release/samples.html#binaries)
- die Dateien aus dem [Repository](https://github.com/multimedial/Hyperledger)
- mysql-Docker-Image, zu beziehen per 
	docker pull mysql/mysql-server
- node.js sowie npm 
- für den [Blockchain Explorer](https://github.com/hyperledger/blockchain-explorer#requirements)


## Schritte zum Aufbau des Demo-Netzwerkes: 
****************************************
*Der Ablauf ist:
- Aufbau der Infrastruktur
- Vorbereitung und Ausführung des Chaincodes
- Visualisierung
- Ausführung von Aufrufen*

**Schritt 1: Aufbau der Infrastruktur**

Ins Verzeichnis "Hyperledger/Network/Schoenhofer" wechseln.

Dort:

	./start.sh
	
Dies startet die Container und erstellt somit die Infrastruktur der Demo:
drei Peers, drei CAs, einen Orderer, einen CouchDB Container. Der MySQL Container für den Blockchain Viewer muss händisch erstellt werden, siehe unten.

Zur Kontroller per docker ps feststellen, dass auch alle Container gestartet wurden.


	Einschub: 

	hier muss der mysql-Container gestartet und mit der fabricexplorer.sql Datei beschickt werden, damit die Tabellen für den Blockchain-Viewer erstellt werden.

	Ich mache dies im Moment manuell:

		Starten des MySQL Servers:
		
			docker run -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_ROOT_HOST=% -p 3306:3306 --name mysql mysql/mysql-server

		In einem anderen Terminal:
		
			docker exec -it mysql mysql -u root -p

		Eingabe des Passworts ("123456"), dann pasten des Inhaltes von "Hyperledger/Network/db/fabricexplorer.sql", damit die Datenbanktabellen erstellt werden können.

	

**Schritt 2: Vorbereitung und Ausführung des Chaincodes**

In den CLI-Container der Blockchain-Infrastruktur wechseln:

	docker exec -it cli bash

Sicherstellen, dass man im Verzeichnis "/opt/gopath/src/docutracker" ist. Ansonsten
	
	cd /opt/gopath/src/docutracker
	
Dort dann ausführen:

	./buildandinstall.sh
	
Der Chaincode wird dann gestartet, letzte Zeile sollte sein " [...] starting up ... "

In einem neuen Terminal dann wieder in den CLI-Container, den Code instanziieren und starten mit:

	docker exec -it cli bash
	
	cd /opt/gopath/src/docutracker
	
	./startcode.sh
	
Ergebniss sollte ohne Fehler sein, und im vorherigen Terminalfenster sollte stehen 

	"#### Smartcontract struct initialized #####"



**Schritt 3: Visualisierung**

In einem neuen unabhängigen Terminal in 

	"Hyperledger/Network/"

wechseln und den Blockchain-Viewer starten. 

** ACHTUNG **: da dies ein separates Projekt ist, muss es vor dem ersten Aufruf gebaut werden mit 

	npm install
	
dann

	./monitor.sh
	
Nun sollte es möglich sein im Browser http://localhost:8080 aufzurufen. Dort sollte dann der Blockchain-Viewer zu sehen sein mit den Peers und mindestens einem Block in der Chain.



**Demo des Chaincodes**

In den CLI Container wechseln, 

	docker exec -it cli bash

und dann 

	./demo.sh
	
Dies füllt die Blockchain mit Transaktionen und Objekten (Workplaces, User, Dokumente). Dies sollte im Browser zu sehen sein.

![Blockchain-Viewer](https://raw.githubusercontent.com/multimedial/Hyperledger/master/Network/BlockchainViewer.jpg "Blockchain-Viewer")