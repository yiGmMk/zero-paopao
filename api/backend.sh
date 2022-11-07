#!bin/env bash

docker stop paopaoswagger

goctl api go -api api.api -dir .

localAddress=$(ifconfig -a | grep inet | grep -v 127.0.0.1 | grep -v inet6 | awk '{print $2}' | tr -d "addr:")
echo "local_address" $localAddress

goctl api plugin -plugin goctl-swagger="swagger -filename api.json -host ${localAddress}:8008 -basepath /" -api api.api -dir .

docker run --rm --name paopaoswagger -d -p 8083:8080 -e SWAGGER_JSON=/foo/api.json -v $PWD:/foo swaggerapi/swagger-ui

go run api.go
