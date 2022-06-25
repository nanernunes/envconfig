GOCMD:=$(shell which go)

test:
	@$(GOCMD) test -v ./...

cover:
	@$(GOCMD) test -v ./... -coverprofile=coverage.txt -covermode=atomic

deps:
	@$(GOCMD) mod download
