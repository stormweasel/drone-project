FROM golang:1.16-alpine3.13
WORKDIR /go/src/project
RUN apk update && apk add bash
SHELL ["/bin/bash", "-o", "pipefail", "-c"]
EXPOSE 8080:8080
COPY ./ ./

RUN go get -d ./...
# run go build ./cmd/server/main.go

ENV dbDriver
ENV dbUser
ENV dbPass
ENV dbHost
ENV dbName

CMD [ "go","run","./cmd/server/main.go" ]