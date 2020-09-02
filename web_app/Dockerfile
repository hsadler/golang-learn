FROM golang:latest

# copy over server files
COPY . /go/src/

# set working directory
WORKDIR /go/src/

# build go binary
RUN go build server.go

# run go binary
CMD ["./server"]

