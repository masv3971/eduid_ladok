NAME 					:= eduid_ladok
VERSION                 := $(shell cat VERSION)
LDFLAGS                 := -ldflags "-w -s --extldflags '-static'"

default: build

build:
		$(info building static binary)
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/${NAME} ${LDFLAGS} ./cmd/main.go
		$(info Done)

test:
		$(info running tests)
		go test -v -cover ./...

update_packages:
		$(info updating packages)
		go get -u ./...

container-build:
		$(info building container)
		docker-compose build

container-start:
		$(info running container)
		docker-compose up -d --remove-orphans

container-stop:
		$(info stopping container)
		docker-compose rm -s -f

container-logs-eduid:
		$(info showing logs for eduid)
		docker logs -f eduid_ladok_service