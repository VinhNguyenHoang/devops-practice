FROM golang:1.19

# set working directory
WORKDIR /app

# copy go mod dependencies
COPY go.mod go.sum ./

# install dependencies
RUN go mod download

# copy source code
COPY *.go .
COPY cmd/ ./cmd/
COPY internal/ ./internal/


# compile go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /backend

# using port
EXPOSE 30000

# define the executable file
ENTRYPOINT [ "/backend" ]