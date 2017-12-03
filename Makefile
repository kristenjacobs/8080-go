all: build

build:
	go build

run:
	./8080-go

test:
	go test -v

fmt:
	go fmt

vet:
	go vet
