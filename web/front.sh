#!bin/env bash

docker stop paopaoweb
imgid=$(docker images | grep "paopao" | awk ' {print $3}')
docker rmi $imgid

localAddress=$(ifconfig -a | grep inet | grep -v 127.0.0.1 | grep -v inet6 | awk '{print $2}' | tr -d "addr:")
echo "local_address" $localAddress

sed -i 's/127.0.0.1/'$localAddress'/' ./.env.local

yarn build

docker build -t your/paopao-ce:web --build-arg USE_DIST=yes .

docker run --name paopaoweb --rm -d -p 8010:80 your/paopao-ce:web
