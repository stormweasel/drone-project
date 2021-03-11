FROM golang:1.16-alpine3.13
WORKDIR /go/src/project
COPY . .

RUN go get -d ./...
# run go build ./cmd/server/main.go

WORKDIR ./cmd/server
CMD [ "go", "run", "main.go"]