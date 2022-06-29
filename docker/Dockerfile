FROM golang:1.18.3-alpine3.16

RUN go install golang.org/x/tools/cmd/godoc@latest

COPY . /go/src/github.com/deblasis/saga-alien-invasion/

WORKDIR /go/src/github.com/deblasis/saga-alien-invasion/
RUN go install github.com/deblasis/saga-alien-invasion

CMD "bash"