FROM golang:1.16-alpine

ARG service

RUN apk --no-cache update \
    && apk add --no-cache ca-certificates tzdata

WORKDIR /workspace

COPY . .

RUN GOOS=linux go build -o ${service} services/${service}/cmd/main.go

USER nobody