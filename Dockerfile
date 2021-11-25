## Compile
FROM golang:1.17 AS builder

WORKDIR /go/src/app

COPY . .

RUN ls -l
#RUN go mod download
RUN make

## Deploy
FROM alpine:3.14

WORKDIR /

COPY --from=builder /go/src/app/bin/eduid_ladok /eduid_ladok

ENV LOGXI=* \
    DEBUG=true \
    HOST=":8080" \
    SCHOOL_NAMES="kf,lnu" \
    KF_SAML_NAME="student.konstfack.se" \
    LNU_SAML_NAME="lnu.se" \
    LADOK_URL="https://api.integrationstest.ladok.se" \
    LADOK_ATOM_PERIODICITY="60" \
    EDUID_IAM_URL="https://api.dev.eduid.se/scim/test" \
    JWT_URL="https://auth-test.sunet.se" \
    LADOK_CERTIFICATE_FOLDER="cert" \
    REDIS_ADDR="redis:6379" \
    REDIS_DB="3"

COPY cert ./cert

CMD [ "./eduid_ladok" ]