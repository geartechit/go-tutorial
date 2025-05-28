FROM golang:1.24-alpine3.21 AS builder
WORKDIR /src
COPY . ./
RUN go mod download
ENV GOOS=linux
RUN go build

FROM alpine:3.17.9 AS runner