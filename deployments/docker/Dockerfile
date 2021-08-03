FROM golang:1.16-alpine

RUN apk --no-cache update \
    && apk add --no-cache ca-certificates tzdata

WORKDIR /workspace

COPY . .

RUN GOOS=linux go build -o quotations services/quotations/cmd/main.go && \
    go build -o btc services/btc/cmd/main.go

USER nobody