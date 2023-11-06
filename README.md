# cellphones

## build

docker build --tag backend:v1 -f ./deployments/Dockerfile .

## run docker

docker run --name backend -d -p 30000:8080 backend:v1 run -s 1
