# Hyperledger

This is a POC for a hyperledger fabric based tracking system. Early work in progress.
Created by multimedial.de for Schoenhofer Sales and Engineering GmbH, Siegburg.
Get in touch at cleske_at_extern.schoenhofer.de

# How to set up the blockchain

The current example is taken from the simple blockchain sample in the fabric-samples directory. It creates 3 nodes with 3 users each and one orderer node.

To set up the custom blockchain:


Prerequisites:

- Download the required binaries for your platform (URL see tutorial)

- Download the fabric-samples


Copy the directory with the simple network example, and give it a custom name. 

In order to set up your custom blockchain, you need to define your own architecture (number of nodes + their names), as well as generate their digital certificates. 
These are self-signed X.509 certificates in this example, using the tool "cryptogen" to create them. 

Thus, after having defined your architecture in the "crypto-config.yaml" file, run the "generate.sh" script to generate the certificates.

You then need to modify the "docker-compose.yml" file to reflect your changes. 

Then, run the "start.sh" to start the docker containers and set up the blockchain.

