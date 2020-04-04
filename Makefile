generate:
	go generate ./...

build: generate
	go build -o ./output/postmanctl ./cmd/postmanctl

install: generate
	go install ./cmd/postmanctl

doc:
	go build -o ./output/genpostmanctldocs ./cmd/genpostmanctldocs
	./output/genpostmanctldocs

.PHONY: generate build install doc