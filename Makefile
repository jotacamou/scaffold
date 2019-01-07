# Makefile for sre/scaffold
# Build:      RPM
# Arch:       x86_64
# Maintainer: jcamou
# Type:       CLI

OWNER?=      sre
PROGRAM?=    scaffold
VERSION?=    v0.0.2
GOARCH?=     amd64
GOOS?=       linux
MAINTAINER?= jcamou@leaf.io

ifeq ($(GOARCH),amd64)
	ARCH := x86_64
endif

.PHONY: dep test build clean

all: dep test build

dep:
	dep ensure -v

test:
	go test ./...

build: dep
	CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -x -v

rpm: build
	rpmdev-setuptree
	cp ${PROGRAM} ${HOME}/rpmbuild/SOURCES
	cp etc/config.yaml ${HOME}/rpmbuild/SOURCES
	rpmbuild -bb package.spec --define "_version ${VERSION}"

clean:
	go clean -x -v
	rm -rf vendor
	rm -rf ${HOME}/rpmbuild/RPMS/${ARCH}/${PROGRAM}-${VERSION}-*.rpm
