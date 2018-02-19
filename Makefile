BINARY=8080-go
CPUPROF=8080-go.cprof
MEMPROF=8080-go.mprof
TRACE=8080-go.trace

all: build fmt vet

.PHONY: build
build:
	go install -v ./cmd/...

.PHONY: clean
clean:
	rm -rf $(BINARY)
	rm -rf $(CPUPROF)
	rm -rf $(MEMPROF)
	rm -rf $(TRACE)

.PHONY: run
run:
	./$(BINARY) -t 2>&1 | tee $(TRACE)

.PHONY: run-test
run-test:
	./$(BINARY) -test -t -s 2>&1 | tee $(TRACE)

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

.PHONY: cprof
cprof:
	./$(BINARY) -cpuprofile $(CPUPROF)

.PHONY: mprof
mprof:
	./$(BINARY) -memprofile $(MEMPROF)

.PHONY: ctop
ctop:
	go tool pprof -top -cum $(BINARY) $(CPUPROF)

.PHONY: mtop
mtop:
	go tool pprof -top -cum $(BINARY) $(MEMPROF)
