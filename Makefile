all: build fmt vet

.PHONY: build
build:
	go build

.PHONY: clean
clean:
	rm -rf ./8080-go
	rm -rf trace.txt

.PHONY: run
run:
	./8080-go -t 2>&1 | tee trace.txt

.PHONY: run-test
run-test:
	./8080-go -test -t 2>&1 | tee trace.txt

.PHONY: test
test:
	go test -v

.PHONY: fmt
fmt:
	go fmt

.PHONY: vet
vet:
	go vet

.PHONY: test-rom
test-rom:
	./test/bin2go.py > test_rom.go
