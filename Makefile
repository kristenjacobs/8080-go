BINARY=8080-space-invaders
TEST_BINARY=8080-test
CPUPROF=8080-go.cprof
MEMPROF=8080-go.mprof
TRACE=8080-go.trace

all: build fmt vet

.PHONY: build
build:
	go install -v ./cmd/...

.PHONY: clean
clean:
	rm -rf $(CPUPROF)
	rm -rf $(MEMPROF)
	rm -rf $(TRACE)

.PHONY: run
run:
	$(BINARY) -t 2>&1 | tee $(TRACE)

.PHONY: run-test
run-test:
	$(TEST_BINARY) -test -t -s 2>&1 | tee $(TRACE)

.PHONY: test
test:
	go test -v ./cmd/... ./pkg/...

.PHONY: fmt
fmt:
	go fmt ./cmd/... ./pkg/...

.PHONY: vet
vet:
	go vet ./cmd/... ./pkg/...

.PHONY: test-rom
test-rom:
	./test/bin2go.py > ./pkg/test/roms.go

.PHONY: cprof
cprof:
	$(BINARY) -cpuprofile $(CPUPROF)

.PHONY: mprof
mprof:
	$(BINARY) -memprofile $(MEMPROF)

.PHONY: ctop
ctop:
	go tool pprof -top -cum $(BINARY) $(CPUPROF)

.PHONY: mtop
mtop:
	go tool pprof -top -cum $(BINARY) $(MEMPROF)
