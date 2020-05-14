
build:
	GITBRANCH=$(shell git rev-parse --abbrev-ref HEAD); \
	GITCOMMIT=$(shell git rev-parse --short HEAD); \
	DATE=$(shell date -u +%Y-%m-%dT%H:%M:%SZ); \
	GOVERSION=$(shell go version | awk '{print $$3}'); \
	GOPLATFORM=$(shell go version | awk '{print $$4}'); \
	go build -ldflags "\
		-X github.com/kevinswiber/postmanctl/internal/runtime/cmd.version=v0.1.0-dev.$$GITBRANCH+$$GITCOMMIT \
		-X github.com/kevinswiber/postmanctl/internal/runtime/cmd.commit=$$GITCOMMIT \
		-X github.com/kevinswiber/postmanctl/internal/runtime/cmd.date=$$DATE \
		-X github.com/kevinswiber/postmanctl/internal/runtime/cmd.goVersion=$$GOVERSION \
		-X github.com/kevinswiber/postmanctl/internal/runtime/cmd.platform=$$GOPLATFORM \
		" -o ./output/postmanctl ./cmd/postmanctl

test: vet
	$(info ******************** running tests ********************)
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

vet:
	$(info ******************** vetting ********************)
	go vet ./...

generate:
	go generate ./...

release:
	GOVERSION=$(shell go version | awk '{print $$3}'); \
	GOPLATFORM=$(shell go version | awk '{print $$4}'); \
	goreleaser --rm-dist

release-snapshot:
	GOVERSION=$(shell go version | awk '{print $$3}'); \
	GOPLATFORM=$(shell go version | awk '{print $$4}'); \
	goreleaser --snapshot --skip-publish --rm-dist

install:
	go install ./cmd/postmanctl

doc:
	go build -o ./output/genpostmanctldocs ./cmd/genpostmanctldocs
	./output/genpostmanctldocs

.PHONY: generate build install doc release release-snapshot vet