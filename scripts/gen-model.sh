#!/usr/bin/env bash

#  根据ddl代码生成model文件
goctl model mysql ddl -src="./migrations/*.sql" -dir="./model" -c