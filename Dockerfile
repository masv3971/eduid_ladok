## Compile
FROM golang:1.17 AS builder

WORKDIR /go/src/app

COPY . .

RUN make

## Deploy
FROM alpine:3.14

WORKDIR /

RUN apk add curl

COPY --from=builder /go/src/app/bin/eduid_ladok /eduid_ladok

HEALTHCHECK --interval=27s CMD curl http://localhost:8080/health | grep -q STATUS_OK

CMD [ "./eduid_ladok" ]