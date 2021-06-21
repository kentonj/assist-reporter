FROM golang:1.16

WORKDIR /go/src/app

COPY go.mod go.mod
COPY go.sum go.sum
COPY main.go main.go
COPY models models

RUN go get -d -v ./...

RUN go build
RUN chmod +x ./assist
CMD ["./assist"]
