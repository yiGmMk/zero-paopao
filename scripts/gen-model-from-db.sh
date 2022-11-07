#!/usr/bin/env bash

user="paopao"
ip="localhost"
port="10020"
password="paopao"
database="paopao"
# 利用网络连接到db在
goctl model mysql datasource -url="$user:$password@tcp($ip:$port)/$database" -table="*" -dir="./model"
