FROM golang:1.8-alpine
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
ENV repo  /go/src/github.com/mikeifomin/docker-sweetheart-data
RUN mkdir -p $repo
ADD ./ $repo

WORKDIR $repo
RUN go get
RUN go build -o main
CMD ["./main"]

