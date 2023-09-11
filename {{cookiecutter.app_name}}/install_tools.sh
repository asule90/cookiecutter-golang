#!/bin/sh
# set -o errexit -eo pipefail

# Database Migrations
go install \
  -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1 

# Linter
go install \
  github.com/golangci/golangci-lint/cmd/golangci-lint@v1.43.0 

# Mockgen
go install github.com/golang/mock/mockgen@v1.6.0

# Cobra
go get -d github.com/spf13/cobra/cobra

# Openapi Code Generator
go get -d github.com/deepmap/oapi-codegen/cmd/oapi-codegen