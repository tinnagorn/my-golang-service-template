#!/bin/sh

set -e

PROJECT="my-project-gcp" \
    TAG="dev" \
    SVC_LIST="my-service" \
    VALUES="values" \
    NAMESPACE="default" \
    REGION="asia-southeast1" \
    ZONE="asia-southeast1-a" \
    KUBE="poc-cluster-cmmn" \
    .deploy-script/deploy.sh
