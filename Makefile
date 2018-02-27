BINARY=8080-space-invaders
TEST_BINARY=8080-test
CPUPROF=8080-go.cprof
MEMPROF=8080-go.mprof

all: build fmt vet

.PHONY: build
build:
	go install -v ./cmd/...

.PHONY: clean
clean:
	rm -rf $(CPUPROF)
	rm -rf $(MEMPROF)

.PHONY: run
run:
	$(BINARY)

.PHONY: run-test
run-test:
	$(TEST_BINARY) -t -s

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
