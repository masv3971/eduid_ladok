.PHONY: update clean build build-all run package deploy test authors dist

NAME 					:= eduid_ladok
VERSION                 := $(shell cat VERSION)
LDFLAGS                 := -ldflags "-w -s --extldflags '-static'"

default: linux

linux: build-eduid_ladok-linux build-fake_environment
mac_m1: build-eduid_ladok-arm build-fake_environment-arm

build-eduid_ladok-arm:
		@echo building eduid_ladok for darwin on arm
		CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -v -o ./bin/eduid_ladok ${LDFLAGS} ./cmd/eduid_ladok/main.go 
		@echo Done

build-fake_environment-arm:
		@echo building fake_environment for darwin on ARM
		CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -v -o ./bin/fake_environment ${LDFLAGS} ./cmd/fake_environment/main.go 
		@echo Done

build-fake_environment-linux:
		@echo building fake_environment for linux on amd64 	
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/fake_environment ${LDFLAGS} ./cmd/fake_environment/main.go

build-eduid_ladok-linux:
		@echo building-static
		CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/${NAME} ${LDFLAGS} ./cmd/eduid_ladok/main.go
		@echo Done
