## Compile
FROM golang:latest AS builder

WORKDIR /go/src/app

COPY . .

RUN make

## Deploy
FROM debian:bookworm-slim

WORKDIR /

RUN apt-get update && apt-get install -y curl procps iputils-ping
RUN rm -rf /var/lib/apt/lists/*

COPY --from=builder /go/src/app/bin/eduid_ladok /eduid_ladok

HEALTHCHECK --interval=27s --timeout=30s CMD curl --connect-timeout 5 http://localhost:8080/health | grep -q STATUS_OK

CMD [ "./eduid_ladok" ]

