#!/bin/sh -ex

export GOPROXY="https://goproxy.cn,direct"

go mod tidy

go build -mod=mod -o bin/state-machine cmd/main.go 