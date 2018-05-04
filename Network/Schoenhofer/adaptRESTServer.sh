docker cp ../../../fabric-sdk-rest/packages/fabric-rest/server/config.json $1:/node_modules/fabric-rest/server/config.json
docker cp ../../../fabric-sdk-rest/packages/fabric-rest/server/datasources.json $1:/node_modules/fabric-rest/server/datasources.json
docker commit $1 fabric-rest:latest