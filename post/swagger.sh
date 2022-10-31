#!bin/env bash

goctl api go -api post.api -dir .

goctl api plugin -plugin goctl-swagger="swagger -filename post.json -host 172.28.192.100:8008 -basepath /" -api post.api -dir .

docker run --rm -p 8083:8080 -e SWAGGER_JSON=/foo/post.json -v $PWD:/foo swaggerapi/swagger-ui
