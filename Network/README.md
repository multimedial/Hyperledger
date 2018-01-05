# Schönhofer Document Tracking POC

## Technische Voraussetzungen
**************
- für Hyperledger:
	- siehe [Hyperledger Dokumentation](https://hyperledger-fabric.readthedocs.io/en/release/prereqs.html), insbesondere Docker und die Go-Umgebung
	
- neueste Hyperledger Binaires und Docker images, laut [Quelle](https://hyperledger-fabric.readthedocs.io/en/release/samples.html#binaries) zu beziehen per 
	> curl -sSL https://goo.gl/byy2Qj | bash -s 1.0.5
	
- die Dateien aus dem [Repository](https://github.com/multimedial/Hyperledger)

- mysql-Docker-Image, zu beziehen per 
	> docker pull mysql/mysql-server
	
- node.js sowie npm für den [Blockchain Explorer](https://github.com/hyperledger/blockchain-explorer#requirements) (ist bereits Teil des lokalen Repository, muss nicht bezogen werden)


****************************************
## Schritte zum Aufbau des Demo-Netzwerkes: 
Der Ablauf ist:
- Aufbau der Infrastruktur
- Vorbereitung und Ausführung des Chaincodes
- Visualisierung
- Ausführung von Aufrufen*

**Schritt 1: Aufbau der Infrastruktur**

Ins Verzeichnis "Hyperledger/Network/Schoenhofer" des Repository wechseln:

	cd Hyperledger/Network/Schoenhofer
	./start.sh
	
Dies startet die Docker Container und erstellt die benötigte Infrastruktur:

* drei Peer Nodes (jedes repräsentiert ein Amt bzw lokale Dienststelle)
* drei Certification Authorities (CA) (für jedes Amt eine),
* einen Orderer (verteilt und ordnet die Anfragen im Netzwerk), 
* einen CouchDB Container zur Verwaltung des WorldState für die Peer Nodes.

Der zusätzliche MySQL-Container für den Blockchain-Viewer muss händisch erstellt werden, siehe unten.

Zur Kontroller per docker ps feststellen, dass auch alle Container gestartet wurden.


	Einschub: 

	hier muss der mysql-Container gestartet und mit der [fabricexplorer.sql](https://github.com/multimedial/Hyperledger/blob/master/Network/db/fabricexplorer.sql) SQL-Datei gefüttert werden, damit die Tabellen für den Blockchain-Viewer definiert sind.

	Im Moment geschieht dies noch etwas umständlich manuell. Starten des MySQL Server Docker Image mit Root-Passwort "123456" und Akzeptanz jeglichen Hosts von aussen für root-Operationen an der Datenbank:
		
			docker run -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_ROOT_HOST=% -p 3306:3306 --name mysql mysql/mysql-server

	In einem anderen Terminal dann in den mysql-client einloggen:
		
			docker exec -it mysql mysql -u root -p

	Dann Eingabe des vorher definierten root-Passworts "123456", dann Einfügen (per Copy-Paste) des Inhaltes von ["Hyperledger/Network/db/fabricexplorer.sql"](https://github.com/multimedial/Hyperledger/blob/master/Network/db/fabricexplorer.sql), damit die Datenbanktabellen erstellt werden.

	

**Schritt 2: Vorbereitung und Ausführung des Chaincodes**

In den CLI-Container der Blockchain-Infrastruktur wechseln:

	docker exec -it cli bash

Sicherstellen, dass man im Verzeichnis "/opt/gopath/src/docutracker" ist. Ansonsten
	
	cd /opt/gopath/src/docutracker
	
Dort dann ausführen:

	./buildandinstall.sh
	
Der Chaincode wird dann gestartet, letzte Zeile sollte sein " [...] starting up ... "

In einem neuen Terminal dann wieder in den CLI-Container einloggen:

	docker exec -it cli bash
	
den Chaincode instanziieren und starten mit:

	cd /opt/gopath/src/docutracker

	./startcode.sh
	
Ergebniss sollte ohne Fehler sein, und im vorherigen Terminalfenster sollte stehen 

	...
	"#### Smartcontract struct initialized #####"



**Schritt 3: Visualisierung der Blockchain Operationen **

** ACHTUNG **: da dies ein separates Projekt ist, muss es vor dem ersten Aufruf gebaut werden mit:

	cd Hyperledger/Network
	npm install

Wenn das Projekt einmal gebaut wurde, den Blockchain-Viewer starten:

	./monitor.sh
	
Nun sollte es möglich sein im Browser http://localhost:8080 aufzurufen. Dort sollte dann der Blockchain-Viewer zu sehen sein mit den Peers und mindestens einem Block in der Chain.


**Demo des Chaincodes**

In den CLI Container wechseln:

	docker exec -it cli bash

und dann ausführen:

	./demo.sh
	
Dies füllt die Blockchain mit Transaktionen und Objekten (Workplaces, User, Dokumente). Es werden in folgender Reihenfolge erstellt: 

- drei Arbeitsorte ("workplace1", "workplace2", "workplace3")
- diesen werden 9 Benutzer zugeordnet, drei je workplace1
- es werden ausserdem 9 Dokumente erstellt mit unterschiedlichen Sicherheitsstufen

Dies sollte im Browser als Transaktionen zu sehen sein:

![Blockchain-Viewer](https://raw.githubusercontent.com/multimedial/Hyperledger/master/Network/BlockchainViewer.jpg "Blockchain-Viewer")