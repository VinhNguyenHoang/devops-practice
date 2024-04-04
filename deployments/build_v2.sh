#!/bin/bash

DOCKERFILE_PATH="Dockerfile"
IMAGE_NAME="backend"
REGISTRY_HOST="localhost:5001"
REGISTRY_REPO="personal"

dir="$(dirname "$(realpath "$0")")"

IMAGE_FULL_PATH=$REGISTRY_HOST/$REGISTRY_REPO/$IMAGE_NAME

docker build -t $IMAGE_FULL_PATH -f "$dir/"$DOCKERFILE_PATH .

docker push $IMAGE_FULL_PATH
