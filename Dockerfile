FROM golang:1.7

WORKDIR /go/src/github.com/rdsubhas/go-elastictable
COPY Godeps /go/src/github.com/rdsubhas/go-elastictable/Godeps
RUN go get -u -v github.com/tools/godep && godep restore -v

COPY . /go/src/github.com/rdsubhas/go-elastictable
CMD ["go", "test", "./..."]
