# syntax=docker/dockerfile:1
FROM golang:1.23 AS build-stage

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY src/go.mod src/go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY src /app/

# Build
RUN CGO_ENABLED=0  GOOS=linux GOARCH=arm go build -o /plcnGoApp


# Deploy the application binary into a lean image
FROM alpine AS build-release-stage

WORKDIR /
# GRPC Server Port 
EXPOSE 7121


# Run
ENTRYPOINT ["/main.go"]
