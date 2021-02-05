#!/bin/bash

set -e

if test -z $TAG; then
    echo "empty tag"
    exit 1
fi

echo "Are you sure to deploy below services \n ${SVC_LIST// /\n } \nto k8s namespace: $NAMESPACE "
read -p "Please confirm [Y/N]? " -n 1 -r
echo # (optional) move to a new line

if [[ $REPLY =~ ^[Yy]$ ]]; then
    DEPLOYNAME="my-service"
    REPOSITORY="asia.gcr.io/$PROJECT/$DEPLOYNAME"
    IMAGE="$REPOSITORY:$TAG"
    HELM=".helm"

    read -p "Do you want to build code in this repo before deploy [Y/N]? " -n 1 -r
    echo # (optional) move to a new line
    if [[ $REPLY =~ ^[Yy]$ ]]; then
        docker build -t $IMAGE .
        docker push $IMAGE
    fi

    echo "Deploy helm script"

    echo ===========================================================
    helm upgrade --install --set name=$SVC_LIST $SVC_LIST -n $NAMESPACE $HELM \
        --set image.tag=$TAG \
        --set image.repository=$REPOSITORY \
        -f $HELM/$VALUES.yaml

fi
