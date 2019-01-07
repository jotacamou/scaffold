# Package:    {{.Name}}
# Build:      docker
# Arch:       x86_64
# Maintainer: {{.Maintainer}}
# Type:       Container

OWNER?=      {{.Owner}}
PROGRAM?=    {{.Program}}
VERSION?=    $(shell git rev-list HEAD --max-count=1 --abbrev-commit)
MAINTAINER?= {{.Maintainer}}
REGISTRY?=   docker-registry.dmdmedia.net:443

.PHONY: dep test build clean

all: dep test build

dep:
	dep ensure

test:
	go test ./...

build: dep
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -x -v

docker: build
	docker build --no-cache -t ${PROGRAM}:${VERSION} .
	docker tag ${PROGRAM}:${VERSION} ${REGISTRY}/${PROGRAM}:${VERSION}
	docker push ${REGISTRY}/${PROGRAM}:${VERSION}

clean:
	go clean -x -v
	rm -rf vendor
