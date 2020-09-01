FROM golang:latest

# copy over server files
COPY . /go/src/app/

# set working directory
WORKDIR /go/src/app

# build go binary
RUN go build server.go

# run go binary
CMD ["./server"]