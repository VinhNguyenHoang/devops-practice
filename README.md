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

## delete image from local registry

### list all images in the repository

curl -X GET "http://localhost:5001/v2/<repository-name>/tags/list

### find image's digest

curl -sS -H "Accept: application/vnd.docker.distribution.manifest.v2+json" -o /dev/null -w "%header{Docker-Content-Digest}" http://localhost:5001/v2/<repository>/manifests/latest

### delete images's manifest

curl -X DELETE "http://localhost:5001/v2/<repository-name>/manifests/<tag>"

k label nodes kind-control-plane name=node1

MongoDB&reg; can be accessed on the following DNS name(s) and ports from within your cluster:

    mongodb-0.mongodb-headless.default.svc.cluster.local:27017
    mongodb-1.mongodb-headless.default.svc.cluster.local:27017

To get the root password run:

    export MONGODB_ROOT_PASSWORD=$(kubectl get secret --namespace default mongodb -o jsonpath="{.data.mongodb-root-password}" | base64 -d)

To get the password for "user1" run:

    export MONGODB_PASSWORD=$(kubectl get secret --namespace default mongodb -o jsonpath="{.data.mongodb-passwords}" | base64 -d | awk -F',' '{print $1}')

To get the password for "user2" run:

    export MONGODB_PASSWORD=$(kubectl get secret --namespace default mongodb -o jsonpath="{.data.mongodb-passwords}" | base64 -d | awk -F',' '{print $2}')

To connect to your database, create a MongoDB&reg; client container:

    kubectl run --namespace default mongodb-client --rm --tty -i --restart='Never' --env="MONGODB_ROOT_PASSWORD=$MONGODB_ROOT_PASSWORD" --image localhost:5001/bitnami/mongodb:7.0.3-debian-11-r6 --command -- bash

Then, run the following command:
mongosh admin --host "mongodb-0.mongodb-headless.default.svc.cluster.local:27017,mongodb-1.mongodb-headless.default.svc.cluster.local:27017" --authenticationDatabase admin -u $MONGODB_ROOT_USER -p $MONGODB_ROOT_PASSWORD
