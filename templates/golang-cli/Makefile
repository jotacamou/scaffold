# Package:    {{.Name}}
# Build:      RPM
# Arch:       x86_64
# Maintainer: jcamou@leaf.io
# Type:       CLI

OWNER?=   {{.Owner}}
PROGRAM?= {{.Program}}
VERSION?= $(shell git rev-list HEAD --max-count=1 --abbrev-commit)

.PHONY: dep test build clean

all: dep test build

dep:
	dep ensure

test:
	go test ./...

build: dep
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -x -v

rpm: build
	rpmdev-setuptree
	cp ${PROGRAM} ${HOME}/rpmbuild/SOURCES
	rpmbuild -bb package.spec --define "_version ${VERSION}"

clean:
	go clean -x -v
	rm -rf vendor
