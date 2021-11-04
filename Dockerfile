## Compile
FROM golang:1.17-stretch AS build

WORKDIR /

COPY internal ./
COPY pkg ./ 
COPY cmd ./
COPY Makefile ./
COPY go.mod ./
COPY go.sum ./
COPY VERSION ./

RUN go mod download
RUN make

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /bin/eduid_ladok-linux /eduid_ladok

USER nonroot:nonroot

ENTRYPOINT ["eduid_ladok"]