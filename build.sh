#!/usr/bin/env bash
set -euo pipefail

go version | grep -q 'go1.11' || (
    go version
    echo error: go1.11 required
    exit 1
)

go_opts=""
if [ "${1-}" = "ci" ]; then
    go_opts="-mod=readonly"
fi

go fmt ./...
go test $go_opts -cover ./...
golint ./... || echo "golint not installed"
go build -o cmd/poker-app/sn-poker cmd/poker-app/main.go

CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o cmd/poker-app/sn-poker-linux cmd/poker-app/main.go
