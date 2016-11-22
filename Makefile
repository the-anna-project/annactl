.PHONY: all annactl clean



export SHELL := /bin/bash
export PATH := ${PATH}:${GOPATH}/bin



GIT_COMMIT := $(shell git rev-parse --short HEAD)
ifndef GOARCH
	GOARCH := $(shell go env GOARCH)
endif
ifndef GOOS
	GOOS := $(shell go env GOOS)
endif
GO_VERSION=$(shell go version | cut -d ' ' -f 3)
PROJECT_VERSION=$(shell cat VERSION)



all: annactl

annactl:
	@go build \
		-o ${GOPATH}/bin/annactl \
		-ldflags " \
			-X main.gitCommit=${GIT_COMMIT} \
			-X main.goArch=${GOARCH} \
			-X main.goOS=${GOOS} \
			-X main.goVersion=${GO_VERSION} \
			-X main.projectVersion=${PROJECT_VERSION} \
		" \
		.

clean:
	@rm -rf ${GOPATH}/bin/annactl
