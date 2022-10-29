#!/usr/bin/env bash

# user="zero-ai-user"
# ip="106.52.162.117"
# port="10020"
# password="q)r88K!LdjM>P3]p@t"
# database="zero-ai"

user="paopao"
ip="localhost"
port="10020"
password="paopao"
database="paopao"
# 利用网络连接到db在
goctl model mysql datasource -url="$user:$password@tcp($ip:$port)/$database" -table="*" -dir="./model"
