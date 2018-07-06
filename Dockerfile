FROM golang:1.10
ADD . /go/src/github.com/satuzcuoglu/authgo
WORKDIR /go/src/github.com/satuzcuoglu/authgo
RUN go get ./...
RUN go install
ENTRYPOINT ["/go/bin/authgo"]