# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

ADD . /go/src/go-server-boilerplate
WORKDIR /go/src/go-server-boilerplate

RUN go-wrapper download   # "go get -d -v ./..."

RUN go-wrapper install    # "go install -v ./..."

RUN go build

# Document that the service listens on port 8080.
EXPOSE 80
EXPOSE 3000

ENTRYPOINT
