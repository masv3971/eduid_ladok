## Compile
FROM golang:1.20.1 AS builder

WORKDIR /go/src/app

COPY . .

RUN make

## Deploy
FROM debian:bullseye

WORKDIR /

RUN apt-get update && apt-get install -y curl procps
RUN rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/src/app/bin/eduid_ladok /eduid_ladok

HEALTHCHECK --interval=27s CMD curl --connect-timeout 5 http://localhost:8080/health | grep -q STATUS_OK

CMD [ "./eduid_ladok" ]