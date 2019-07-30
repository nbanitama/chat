FROM golang:1.8

WORKDIR /go/src/nbanitama/chat
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build

CMD ["./chat"]