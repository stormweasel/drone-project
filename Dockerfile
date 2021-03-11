FROM golang:latest
WORKDIR /go/src/project
COPY . .

RUN go get -d -v ./...

CMD ./cmd/server/go run     
