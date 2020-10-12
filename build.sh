#!/bin/bash
#
# go version go1.14.9 linux/amd64

#
echo 'go version '
go version

export CGO_ENABLED=0
export GO111MODULE=on
export GOPROXY="https://mirrors.aliyun.com/goproxy/"

#
echo 'go build ...'
go build

## 查看依赖
## ldd xxx

#
echo 'docker tag ...'
docker build -t harbor.hd.com/gin-gorm-demo .
#docker tag gin-gorm-demo harbor.hd.com/gin-gorm-demo
