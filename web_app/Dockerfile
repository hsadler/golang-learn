FROM golang:latest

# copy over server files
COPY . /go/src/

# set working directory
WORKDIR /go/src/

# use CompileDaemon to hot-reload building and spinning-up the server
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon \
    --build="go build server.go" \
    --command=./server


