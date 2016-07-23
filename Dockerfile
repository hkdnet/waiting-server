FROM golang:1.6.3-alpine
ADD . $GOPATH/src/github.com/hkdnet/waiting-server

CMD ["go", "run", "/go/src/github.com/hkdnet/waiting-server/main.go"]

