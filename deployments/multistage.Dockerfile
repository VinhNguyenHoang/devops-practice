# Use a GoLang base image for building the application
FROM golang:1.19 AS builder

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

# Use a minimal base image for the final image
FROM alpine:3.17

# Set the working directory inside the container
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /backend .

# using port
EXPOSE 30000

# define the executable file
ENTRYPOINT [ "/app/backend" ]