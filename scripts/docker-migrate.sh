#!/usr/bin/env bash

# 当前路径
sh_dir=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)


# 获取docker
# 没有设置network的可用这么取,docker inspect --format '{{ .NetworkSettings.Networks.IPAddress }}' container名称或id
# 这里使用名称 mysql-master
# 这里的mysql使用了network,需要取NetworkSettings.Networks里面的
mysql_ip=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' mysql-master)
psw="q)r88K!LdjM>P3]p@t"

port=":3306"
ip=$(echo $mysql_ip)
# echo $port
# echo "($ip$port)"
ip_port=$(echo "($ip$port)")
echo 容器ip+端口: $ip_port

# 参考 https://github.com/golang-migrate/migrate
# mysql://user:password@tcp(host:port)/dbname?query
# run --rm  运行结束后删除container
docker run --rm -v $sh_dir/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database=mysql://zero-ai-user:$psw@tcp$ip_port/zero-ai up
