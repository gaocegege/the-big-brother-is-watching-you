# Makefile
#
# Targets:
#   build build-local

build build-local:
	godep go build -race .
.PHONY: build build-local

get-deps:
	go get ./...
.PHONY: get-deps

build-travis:
	go build -race .
.PHONY: build-travis