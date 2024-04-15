#!/bin/bash

if [ $# -lt 1 ]; then
    echo "Required image name."
    exit 1
fi

acceptM="application/vnd.docker.distribution.manifest.v2+json"
acceptML="application/vnd.docker.distribution.manifest.list.v2+json"
registry='localhost:5001'
name="$1"

if [ -z "$name" ]; then
    echo "Image name is empty."
    exit 1
fi

curl -H "Accept: ${acceptM}" \
    -H "Accept: ${acceptML}" \
    -v -sSL -X DELETE "http://${registry}/v2/${name}/manifests/$(
    curl -sSL -I \
        -H "Accept: ${acceptM}" \
        -H "Accept: ${acceptML}" \
        "http://${registry}/v2/${name}/manifests/$(
            curl -sSL "http://${registry}/v2/${name}/tags/list" | jq -r '.tags[0]'
        )" \
    | awk '$1 == "Docker-Content-Digest:" { print $2 }' \
    | tr -d $'\r' \
)"

docker exec kind-registry /bin/registry garbage-collect   /etc/docker/registry/config.yml --delete-untagged