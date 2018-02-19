.PHONY: build docker run test generate travis dep tools

current_dir = $(shell pwd)
go_version = 1.10

# To be run outside containers
docker:
	docker pull xena/go:$(go_version)
	docker build -t horseville/core -f Dockerfile.core .
	docker build -t horseville/horsevilled .

run: docker
	docker-compose up -d

# To be run inside containers
tools:
	retool build

dep:
	retool do dep ensure -update
	retool do dep prune

generate: tools
	retool do statik -src ./public -f
	sh -c 'cd ./internal/database/migrations && retool do go-bindata -pkg=dmigrations -o=../dmigrations/bindata.go .'
	sh -c 'cd ./rpc/horseville && retool do sh ./regen.sh'

test: generate
	go test ./...

build: generate
	GOBIN=$(current_dir)/bin go install ./cmd/...

travis: build test
	true
