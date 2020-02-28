#!/usr/bin/env bash
set -e

docker run --rm -e http_proxy=$http_proxy -e https_proxy=$https_proxy -e no_proxy=$no_proxy -v "$(dirname "$PWD")":/go/src -w /go/src/api golang:1.13-alpine go build -o api
docker build -t ekaraplatform/api .
