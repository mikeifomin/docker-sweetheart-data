FROM golang:1.8-alpine

ADD . /go/app
WORKDIR /go/app
RUN go get
RUN go build -o main
CMD ["./main"]
