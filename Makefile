DOCKER_IMAGE_NAME := quay.io/arschles/goprox:devel

build:
	go build -o goprox

docker-build:
	docker build --rm -t ${DOCKER_IMAGE_NAME} .

test:
	go test $$(glide nv)
