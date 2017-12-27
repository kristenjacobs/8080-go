all: build fmt vet

build:
	go build

clean:
	rm -rf ./8080-go
	rm -rf trace.txt

run:
	./8080-go 2>&1 | tee trace.txt

test:
	go test -v

fmt:
	go fmt

vet:
	go vet
