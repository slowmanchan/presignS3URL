FROM golang:1.9-alpine

ENV GOBIN /go/bin

COPY Gopkg.toml Gopkg.lock ./

COPY . /go/src/github.com/slowmanchan/presignS3URL
WORKDIR /go/src/github.com/slowmanchan/presignS3URL

RUN dep ensure -vendor-only
RUN go build 

CMD ["go", "run", "main.go"]
