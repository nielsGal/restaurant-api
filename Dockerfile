FROM golang:1.14

WORKDIR /go/src/github.com/nielsGal/restaurant-api/
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build server.go

CMD ["./server"]