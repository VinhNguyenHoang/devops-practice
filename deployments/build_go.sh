#!/bin/bash

action="$1"

DOCKERFILE_PATH="multistage.Dockerfile"
IMAGE_NAME="backend"
REGISTRY_HOST="localhost:5001"
REGISTRY_REPO="personal"

dir="$(dirname "$(realpath "$0")")"

# TODO: use timestamp as tag
IMAGE_FULL_PATH=$REGISTRY_HOST/$REGISTRY_REPO/$IMAGE_NAME":latest"

buildandpush () {
    buildonly
    echo "Pushing to local registry..."
    docker push $IMAGE_FULL_PATH
}

buildonly () {
    if docker images --format "{{.Repository}}:{{.Tag}}" | grep -q $IMAGE_FULL_PATH; then
        echo "Deleting $IMAGE_FULL_PATH"
        docker rmi $IMAGE_FULL_PATH
    fi

    echo "Building image..."
    docker build -t $IMAGE_FULL_PATH -f "$dir/"$DOCKERFILE_PATH . --no-cache
}

case $action in
    # build
    p)
        buildandpush
        ;;
    # push
    b)
        buildonly
        ;;
    *)
        echo "action not supported"
        ;;
esac
