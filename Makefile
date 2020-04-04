generate:
	go generate ./...

build:
	go build -o ./output/postmanctl ./cmd/postmanctl

install:
	go install ./cmd/postmanctl

doc:
	go build -o ./output/genpostmanctldocs ./cmd/genpostmanctldocs
	./output/genpostmanctldocs

.PHONY: generate build install doc