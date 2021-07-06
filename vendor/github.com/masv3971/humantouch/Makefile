.PHONY: update clean build build-all run package deploy test authors dist

NAME 					:= humantouch
VERSION                 := $(shell cat VERSION)

default: release-patch

release-patch: tidy test add commit-msg release-tag push-tag go-list
		@echo relese ${NAME}@${VERSION} 

tidy:
		@echo tidy up..
		go mod tidy

test:
		@echo test ${NAME}
		go test -v --cover .
add:
	git add .

commit-msg:
		git commit -m"Humantouch release ${VERSION}"

release-tag:
		git tag ${VERSION}

push-tag:
		git push origin ${VERSION}

go-list:	
		GOPROXY=proxy.golang.org go list -m github.com/masv3971/${NAME}@${VERSION}