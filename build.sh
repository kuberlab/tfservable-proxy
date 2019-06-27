#!/usr/bin/env bash

go build -v -ldflags="-s -w" -o tfservable-proxy tfproxy.go

