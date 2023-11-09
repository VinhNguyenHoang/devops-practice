# cellphones

## build

docker build --tag backend:v1 -f ./deployments/Dockerfile .

## run docker

docker run --name backend -d -p 30000:8080 backend:v1 run -s 1

## get list of images in local registry

curl -X GET "http://localhost:5001/v2/\_catalog"

## get tags of an image

curl -X GET "http://localhost:5001/v2/<image>/tags/list"

sha256:6bde8e70a73153b22430ef2f547f3eb651491dcde52ba7d10fc9258ee5c00aed
