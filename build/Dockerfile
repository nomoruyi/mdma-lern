FROM golang:1.20 AS builder
LABEL stage=intermediate
ARG VERSION
COPY . /out
WORKDIR /out
ENV GO111MODULE=on
RUN CGO_ENABLED=0 go build -ldflags "-X main.commitHash=$(git rev-parse HEAD)" -o backend cmd/backend/*