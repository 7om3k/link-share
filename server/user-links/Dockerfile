# syntax=docker/dockerfile:1

FROM golang:1.26.4
LABEL authors="github.com/7om3k"

WORKDIR /app

# Docker's COPY command is quite sensitive in its interpretation of the trailing slash
COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

# Binary is located in the root of the filesystem of the image
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-user-links

CMD ["/docker-user-links"]
