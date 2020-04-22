#!/usr/bin/env bash

go build -o compiler cmd/compiler/main.go
./compiler
go build -o out output.go
./out