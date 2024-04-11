#!/bin/bash

if [ $# -lt 1 ]; then
    echo "Required image name."
    exit 1
fi

registry='localhost:5001'
name="$1"

if [ -z "$name" ]; then
    echo "Image name is empty."
    exit 1
fi

curl -v -sSL -X DELETE "http://${registry}/v2/${name}/manifests/$(
    curl -sSL -I \
        -H "Accept: application/vnd.docker.distribution.manifest.v2+json" \
        "http://${registry}/v2/${name}/manifests/$(
            curl -sSL "http://${registry}/v2/${name}/tags/list" | jq -r '.tags[0]'
        )" \
    | awk '$1 == "Docker-Content-Digest:" { print $2 }' \
    | tr -d $'\r' \
)"

docker exec -it kind-registry bin/registry garbage-collect /etc/docker/registry/config.yml