## Schönhofer Network Config

Note that this basic configuration uses pre-generated certificates and
key material, and has  predefined transactions to initialize a 
channel named "vertraulich".

To regenerate this material, simply run ``generate.sh``.

To start the network, run ``start.sh``.
To stop it, run ``stop.sh``
To completely remove all incriminating evidence of the network
on your system, run ``teardown.sh``.


------------------------------------------------------------------------------------------------------------------------------------------------------------------
------------------------------------------------------------------------------------------------------------------------------------------------------------------

README

Dieses ReadMe beschreibt den Aufbau und Start des Schönhofer-Netzwerkes.
Die aktuelle Version ist abgeleitet vom first-network Beispiel aus den fabric-samples. 


Um ein eigenes Netzwerk zu definieren, müssen mehrere Skripte aufgerufen werden: 

- generate.sh (um die notwendigen Docker Images, Zertifikate, logischen Kanäle und Konfiguration der Blockchain einzurichten)

- start.sh (um die Docker Images zu starten)

- stop.sh (stoppt die Blockchain)

- teardown.sh (löscht die Docker Images)


generate.sh
Generate.sh definiert vor allen Dingen die digitalen Zertifikate und logischen Kanäle (=Ledgers) der Blockchain. Jeder Kanal ist eine Blockchain.

start.sh
Startet die Docker Container mit der eigentlichen Blockchain.

stop.sh
Stoppt die laufenden Docker Container 

teardown.sh 
Beendet und löscht die Docker Images.
