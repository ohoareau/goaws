ARCH ?= linux-amd64
platform_temp = $(subst -, ,$(ARCH))
GOOS = $(word 1, $(platform_temp))
GOARCH = $(word 2, $(platform_temp))
GOPROXY = https://proxy.golang.org

export CI

arch:
	@echo $(shell go env GOOS)-$(shell go env GOARCH)

format:
	@gofmt -w *.go

install:
	@go mod vendor

test:
	@go test *.go

.DEFAULT_GOAL := install
.PHONY: arch \
		format \
		install \
		test
