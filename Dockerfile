ARG GO_VERSION=1
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -o /run-app .

FROM debian:bookworm

RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates
RUN apt-get install curl -y

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz -C /usr/local/bin

COPY --from=builder /run-app /usr/local/bin/

COPY templates /templates
COPY static /static
COPY migrations /migrations

CMD ["run-app"]
