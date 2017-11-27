############################################################
# base image
############################################################
FROM hyperledger/fabric-ccenv:x86_64-1.0.4

############################################################
# set the workdir
############################################################
WORKDIR /opt/gopath/src/docutracker

############################################################
# copy stuff that is needed
############################################################
ADD ./devenv/msp /etc/hyperledger/msp
ADD ./go/src /opt/gopath/src
ADD ./go/src/docutracker /opt/gopath/src/docutracker

############################################################
# Environment variables
############################################################
ENV GOPATH /opt/gopath
ENV CORE_VM_ENDPOINT unix:///host/var/run/docker.sock
ENV CORE_LOGGING_LEVEL DEBUG
ENV CORE_PEER_ID example02
ENV CORE_PEER_ADDRESS peer:7051
ENV CORE_PEER_LOCALMSPID DEFAULT
ENV CORE_PEER_MSPCONFIGPATH /etc/hyperledger/msp

EXPOSE 7050