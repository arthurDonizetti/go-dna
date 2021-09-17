FROM golang:1.15

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...

CMD ["go", "run", "cmd/main.go"]