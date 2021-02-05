#!/bin/sh

set -e

if test -z $TAG; then
    echo "empty tag"
    exit 1
fi

ZONE="asia-southeast1-a"
DEPLOYNAME="my-service"
REPOSITORY="asia.gcr.io/$PROJECT/$DEPLOYNAME"
IMAGE="$REPOSITORY:$TAG"

docker build -t $IMAGE .
docker push $IMAGE
