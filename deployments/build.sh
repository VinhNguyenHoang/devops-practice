#!/bin/bash

imageName="backend"

timestamp=$(date +%Y%m%d%H%M%S)

tag=$imageName:$timestamp
# latest=$imageName:latest

dir="$(dirname "$(realpath "$0")")"

# # build
# docker build -t "$tag" -f "$dir"/Dockerfile . --progress=plain --no-cache
docker build -t "$tag" -f "$dir"/Dockerfile . --no-cache

# # echo latest tag
# echo "$tag"