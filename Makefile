.PHONY: build docker run test generate travis dep tools

current_dir = $(shell pwd)
go_version = 1.10

tools:
	retool build

dep:
	retool do dep ensure -update
	retool do dep prune

docker:
	docker pull xena/go:$(go_version)
	docker build -t horseville/horsevilled .

run: docker
	docker-compose up -d

generate: tools
	retool do statik -src ./public -f
	sh -c 'cd ./internal/database/migrations && retool do go-bindata -pkg=dmigrations -o=../dmigrations/bindata.go .'

test: generate
	go test ./...

build: generate
	GOBIN=$(current_dir)/bin go install ./cmd/...

travis: build test
	true
