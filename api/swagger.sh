#!bin/env bash

docker stop paopaoswagger

goctl api go -api api.api -dir .

goctl api plugin -plugin goctl-swagger="swagger -filename api.json -host 172.28.192.100:8008 -basepath /" -api api.api -dir .

docker run --rm --name paopaoswagger -d -p 8083:8080 -e SWAGGER_JSON=/foo/api.json -v $PWD:/foo swaggerapi/swagger-ui
