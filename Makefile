SHELL := /bin/bash

VERSION := $(shell git rev-parse --short HEAD 2> /dev/null || echo 'unknown')

BUILD_DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)

.PYTHON: build-mipsle
build-mipsle:
	GOOS=linux GOARCH=mipsle go build -ldflags "-s -w -X main.Version=$(VERSION) -X main.BuildDate=$(BUILD_DATE)" -o bin/lnutnc_mipsle
	upx bin/lnutnc # upx 3.95 must be used.

.PYTHON: build
build:
	GOOS=linux GOARCH=mipsle go build -ldflags "-s -w -X main.Version=$(VERSION) -X main.BuildDate=$(BUILD_DATE)" -o bin/lnutnc

