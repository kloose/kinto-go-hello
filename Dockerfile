FROM golang:1.10.0

WORKDIR /go/src/app

COPY . .

RUN go get -d -v github.com/gorilla/mux

RUN go build -o kinto-go-hello

ENTRYPOINT ["/go/src/app/kinto-go-hello"]