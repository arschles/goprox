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
	${DEV_ENV_CMD} glide install

build:
	go build -o goprox

build-alpine:
	${BUILD_ALPINE_CMD_PREFIX} go build -o rootfs/bin/goprox

build-cli:
	make -C cli build

docker-build:
	docker build --rm -t ${DOCKER_IMAGE_NAME} rootfs

docker-push:
	docker push ${DOCKER_IMAGE_NAME}

docker-deploy: build-alpine docker-build docker-push

test:
	${DEV_ENV_CMD} go test $$(glide nv)

run-test:
	docker run --rm -e AWS_KEY=${GOPROX_AWS_KEY} -e AWS_SECRET=${GOPROX_AWS_SECRET} ${DOCKER_IMAGE_NAME}

codegen-admin:
	make -C ./_proto codegen-admin

deploy-to-deis:
	${DEIS_BINARY_NAME} pull ${DOCKER_IMAGE_NAME} -a ${DEIS_APP_NAME}
