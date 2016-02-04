# Makefile
#
# Targets:
#   build build-local

build build-local:
	godep go build -race .
.PHONY: build build-local
