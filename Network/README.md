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
- Aufziehen der Infrastruktur
- Vorbereitung und Ausführung des Chaincodes
- Visualisierung durch den Blockchain-Viewer (optional)
- Ausführung von Aufrufen*

**Schritt 1: Aufziehen der Infrastruktur**

Ins Verzeichnis "Hyperledger/Network/Schoenhofer" des Repository wechseln und die Docker-Container per Shell-Skript starten:

	cd Hyperledger/Network/Schoenhofer
	
	./start.sh
	
Dies startet die benötigten Docker Container und erstellt somit die benötigte Infrastruktur. Es werden erstellt: 

* drei Peer Nodes (jedes repräsentiert ein Amt bzw lokale Dienststelle)
* drei Certification Authorities (CA) (für jedes Amt eine),
* einen Orderer (verteilt und ordnet die Anfragen im Netzwerk), 
* einen CouchDB Container zur Verwaltung des WorldState für die Peer Nodes.
* ein MySQL Container für den optionalen Blockchain-Viewer (siehe "Einschub").

Zur Kontroller per docker ps feststellen, dass auch alle Container gestartet wurden.

	Einschub: MySQL-Container für den optionalen Blockchain-Viewer

	Der MySql-Container für den optionalen Blockchain-Viewer muss im Moment noch händisch gestartet und die benötigte Datenbank mit der [fabricexplorer.sql](https://github.com/multimedial/Hyperledger/blob/master/Network/db/fabricexplorer.sql) Datei konfiguriert werden. Dies erstellt die benötigten Tabellen. Starten des MySQL Server Docker Image mit Root-Passwort "123456" und Zulassen, dass der DB-Root-User sich auch von externen Hosts an der Datenbank anmelden darf:
		
			docker run -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_ROOT_HOST=% -p 3306:3306 --name mysql mysql/mysql-server

	In einem anderen Terminal dann in den mysql-client einloggen:
		
			docker exec -it mysql mysql -u root -p

	Man wird zur Eingabe des vorher definierten root-Passworts "123456" aufgefordert, dann Einfügen (per Copy-Paste) des Inhaltes von ["Hyperledger/Network/db/fabricexplorer.sql"](https://github.com/multimedial/Hyperledger/blob/master/Network/db/fabricexplorer.sql), damit die benötigten Datenbanktabellen erstellt werden.
	
	TODO: 
	Automatisieren bzw Initialisieren des Docker-Containers mit der SQL Datei.
	

**Schritt 2: Vorbereitung und Ausführung des Chaincodes**

In den CLI-Container der Blockchain-Infrastruktur wechseln:

	docker exec -it cli bash

Sicherstellen, im Verzeichnis "/opt/gopath/src/docutracker" zu sein:
	
	cd /opt/gopath/src/docutracker
	
Dann ausführen:

	./buildandinstall.sh
	
Der Chaincode wird dann gestartet, letzte Zeile im Terminal sollte sein 

	" [...] starting up ... "

Dieses Terminal zur Kontrolle offen belassen. In einem neuen Terminal wieder in den CLI-Container einloggen:

	docker exec -it cli bash
	
den Chaincode per Shell-Skript instanzieren und starten mit:

	cd /opt/gopath/src/docutracker
	
	./startcode.sh
	
Ergebniss sollte ohne Fehler sein, und im vorherigen, offengelassenen Terminal sollte nun stehen:

	...
	"#### Smartcontract struct initialized #####"



**Schritt 3: Visualisierung der Blockchain Operationen **

** ACHTUNG **: da dies ein separates Projekt ist, muss es vor dem ersten Aufruf gebaut werden mit:

	cd Hyperledger/Network
	npm install

Wenn der Blockchain-Viewer einmal erstellt wurde, kann der Blockchain-Viewer starten per Shell-Skript:

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