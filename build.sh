#!/bin/sh
cd `dirname $0`
GOOS=js GOARCH=wasm go build -o webroot/add_cookie/js/main.wasm src/main.go
