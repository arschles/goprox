#!/usr/bin/env bash
#
# Build and push Docker images to Docker Hub and quay.io.
#

cd "$(dirname "$0")" || exit 1

docker login -e="$QUAY_EMAIL" -u="$QUAY_USERNAME" -p="$QUAY_PASSWORD" quay.io
make -C .. docker-push
curl -sSL http://deis.io/deis-cli/install-v2.sh | bash
./deis login --username=$DEIS_USERNAME --password=$DEIS_PASSWORD deis.arschles.net
DEIS_BINARY_NAME=./_scripts/deis DEIS_APP_NAME=goprox make -C .. deploy-to-deis
