VERSION ?= $(shell git rev-parse --short HEAD)
DOCKER_IMAGE_NAME := quay.io/arschles/goprox:${VERSION}

# dockerized development environment variables
REPO_PATH := github.com/arschles/goprox
DEV_ENV_IMAGE := quay.io/deis/go-dev:0.9.1
DEV_ENV_WORK_DIR := /go/src/${REPO_PATH}
DEV_ENV_PREFIX := docker run --rm -e GO15VENDOREXPERIMENT=1 -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR}
DEV_ENV_CMD := ${DEV_ENV_PREFIX} ${DEV_ENV_IMAGE}
BUILD_ALPINE_CMD_PREFIX := ${DEV_ENV_PREFIX} -e CGO_ENABLED=0 -e GOOS=linux -e GOARCH=amd64 ${DEV_ENV_IMAGE}

bootstrap:
	glide install

build: 
	make -C cmd/server build
	make -C cmd/cli build

docker-build:
	docker build --rm -t ${DOCKER_IMAGE_NAME} rootfs

docker-push:
	docker push ${DOCKER_IMAGE_NAME}

test:
	go test $$(glide nv)

codegen:
	make -C ./_proto codegen
