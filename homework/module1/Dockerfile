FROM golang:latest
WORKDIR $GOPATH/src/httpserver
COPY httpserver.go $GOPATH/src/httpserver/httpserver.go
RUN go build httpserver.go
EXPOSE 8080
ENTRYPOINT  ["./httpserver"]