FROM golang:alpine as builder

WORKDIR /ecommerce_app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ." -command="./ecommerce_app"
