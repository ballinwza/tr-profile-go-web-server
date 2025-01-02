#!/usr/bin/env bash
export GO_ENV=prod
go build -o ./app main.go rooms.go routes.go stats.go
