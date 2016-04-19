VERSION ?= $(git rev-parse --short HEAD)
DOCKER_IMAGE_NAME := quay.io/arschles/goprox:${VERSION}

# dockerized development environment variables
REPO_PATH := github.com/arschles/goprox
DEV_ENV_IMAGE := quay.io/deis/go-dev:0.9.1
DEV_ENV_WORK_DIR := /go/src/${REPO_PATH}
DEV_ENV_PREFIX := docker run --rm -e GO15VENDOREXPERIMENT=1 -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR}
BUILD_ALPINE_CMD_PREFIX := ${DEV_ENV_PREFIX} -e CGO_ENABLED=0 -e GOOS=linux -e GOARCH=amd64 ${DEV_ENV_IMAGE}
TEST_CMD_PREFIX := ${DEV_ENV_PREFIX}

bootstrap:
	${DEV_ENV_CMD} glide install

build:
	go build -o goprox

build-alpine:
	${BUILD_ALPINE_CMD_PREFIX} go build -o rootfs/bin/goprox

docker-build:
	docker build --rm -t ${DOCKER_IMAGE_NAME} rootfs

docker-push:
	docker push ${DOCKER_IMAGE_NAME}

test:
	${TEST_CMD_PREFIX} go test $$(glide nv)
