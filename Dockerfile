## Compile
FROM golang:1.17 AS builder

WORKDIR /go/src/app

COPY . .

RUN make

## Deploy
FROM alpine:3.14

WORKDIR /

COPY --from=builder /go/src/app/bin/eduid_ladok /eduid_ladok

CMD [ "./eduid_ladok" ]