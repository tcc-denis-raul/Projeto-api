#!/bin/bash

#Um simples script para executar o servidor web

mkdir -p ${GOPATH}/src/paloma_api
cp -r . ${GOPATH}/src/paloma_api
cd ${GOPATH}/src/paloma_api/api

go get .
go run *.go

sudo rm -r ${GOPATH}/src/paloma_api
