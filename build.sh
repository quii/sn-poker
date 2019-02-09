#!/usr/bin/env bash

set -e

go test $(go list ./... | grep -v /vendor/)
go vet $(go list ./... | grep -v /vendor/)
go fmt $(go list ./... | grep -v /vendor/)
golint $(go list ./... | grep -v /vendor/)
