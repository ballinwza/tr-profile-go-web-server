#!/usr/bin/env bash
GO_ENV=production go build -o ./app main.go rooms.go routes.go stats.go
