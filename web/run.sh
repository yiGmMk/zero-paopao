#!bin/env bash

docker stop paopaoweb
imgid=$(docker images | grep "paopao" | awk ' {print $3}')
docker rmi $imgid

yarn build

docker build -t your/paopao-ce:web --build-arg USE_DIST=yes .

docker run --name paopaoweb --rm -d -p 8010:80 your/paopao-ce:web
