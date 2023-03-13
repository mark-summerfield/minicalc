#!/bin/bash
clc -sS
go mod tidy
go fmt .
staticcheck .
go vet .
golangci-lint run
git st
