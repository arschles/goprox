VERSION ?= $(shell git rev-parse --short HEAD)
DOCKER_IMAGE_NAME := quay.io/arschles/goprox:${VERSION}

# dockerized development environment variables
REPO_PATH := github.com/arschles/goprox
DEV_ENV_IMAGE := quay.io/deis/go-dev:v0.22.0
DEV_ENV_WORK_DIR := /go/src/${REPO_PATH}
DEV_ENV_PREFIX := docker run --rm -v ${CURDIR}:${DEV_ENV_WORK_DIR} -w ${DEV_ENV_WORK_DIR}
DEV_ENV_CMD := ${DEV_ENV_PREFIX} ${DEV_ENV_IMAGE}

SERVER_BUILD_CMD := go build -o cmd/server/goproxd ./cmd/server
CLI_BUILD_CMD := go build -o cmd/cli/goprox ./cmd/cli

build:
ifdef DOCKER
	${DEV_ENV_CMD} ${CLI_BUILD_CMD}
	${DEV_ENV_CMD} ${SERVER_BUILD_CMD}
else
	${CLI_BUILD_CMD}
	${SERVER_BUILD_CMD}
endif

docker-build:
	@echo "nothing to do yet"
	# docker build --rm -t ${DOCKER_IMAGE_NAME} rootfs

docker-push:
	@echo "nothing to do yet"
	# docker push ${DOCKER_IMAGE_NAME}

test:
ifdef DOCKER
	${DEV_ENV_CMD} go test $$(glide nv)
else
	go test $$(glide nv)
endif

codegen:
	make -C ./_proto codegen
