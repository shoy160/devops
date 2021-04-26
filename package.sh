#!/bin/bash
cd fed 
yarn build

cd .. 
EXPORT CGO_ENABLED=0
EXPORT GOOS=linux
EXPORT GOARCH=amd64
go build -o devops main.go
version=latest
docker build -t hub.app-chengdu1.yunzhicloud.com/yzcloud/devops:$version .
docker push hub.app-chengdu1.yunzhicloud.com/yzcloud/devops:$version
docker rmi hub.app-chengdu1.yunzhicloud.com/yzcloud/devops:$version
