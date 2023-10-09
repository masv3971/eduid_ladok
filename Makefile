NAME 					:= eduid_ladok
VERSION                 := $(shell cat VERSION)
LDFLAGS                 := -ldflags "-w -s --extldflags '-static' -X main.version=$(git describe --always --long --dirty)"

default: build

build:
		$(info building static binary)
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/${NAME} ${LDFLAGS} ./cmd/main.go
		$(info Done)

test:
		$(info running tests)
		go test -v -cover ./...

get_release-tag:
	@date +'%Y%m%d%H%M%S%9N'

ifndef VERSION
VERSION := latest
endif

RELEASE_TAG := docker.sunet.se/eduid/eduid_ladok:$(VERSION)

hard_restart: docker-stop docker-start

docker-push:
	$(info Pushing docker images)
	docker push $(RELEASE_TAG)

update_packages:
		$(info updating packages)
		go get -u ./...

docker-build:
		$(info building in docker)
		docker build --tag $(RELEASE_TAG) --file Dockerfile .

docker-start:
		$(info running docker)
		docker-compose up -d --remove-orphans

docker-stop:
		$(info stopping docker)
		docker-compose rm -s -f

docker-logs-eduid:
		$(info showing logs for eduid)
		docker logs -f eduid_ladok_service