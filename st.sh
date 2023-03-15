#!/bin/bash
clc -s -e eval
go mod tidy
go fmt .
staticcheck .
go vet .
golangci-lint run
git st
