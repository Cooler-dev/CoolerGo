PACKAGE_NAME := github.com/Cooler-dev/CoolerGo
VERSION      ?= $(shell git describe --tags --abbrev=0)
COMMIT_HASH  := $(shell git rev-parse --short HEAD)
DEV_VERSION  := dev-${COMMIT_HASH}
GO_VERSION   ?= 1.18

install:
	go mod tiny

build:
	./scripts/build-backend.sh
