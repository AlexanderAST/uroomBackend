FROM golang:1.19-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update && apt-get -y install postgresql-client

RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o uroomBackend ./cmd/uroom/main.go

LABEL authors="alexander"

CMD ["./uroomBackend"]