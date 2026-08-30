GOSPN_VERSION = 0.18.0
GOSPN_LDFLAGS = -X main.version=$(GOSPN_VERSION)

deps:
	go get github.com/antlr/antlr4/runtime/Go/antlr@4.7.2
	cd pkg/parser/ && antlr4 -Dlanguage=Go -listener JSPNL.g4

build: deps
	mkdir -p bin
	go build -ldflags "$(GOSPN_LDFLAGS)" -o bin/gospn cmd/gospn.go

# The test/ directory exists, so without this `make test` says "up to date".
.PHONY: test test_e2e test_result test_jsonout test_npzout test_analysis test_mt \
	test_matout test_petrinet test_parser test_mxgraph test_benchmark

# Everything, which is what CI runs. `make test` used to name four pkg/ directories by
# hand, so the packages added in 0.18.0 and the golden tests in test/ were not in it at
# all: a passing `make test` meant less than it looked. The per-package targets below
# are still here for running one suite verbosely.
test:
	go test ./...

test_e2e:
	go test -v ./test/

test_result:
	cd pkg/result/ && go test -v -cover

test_jsonout:
	cd pkg/jsonout/ && go test -v -cover

test_npzout:
	cd pkg/npzout/ && go test -v -cover

test_analysis:
	cd pkg/analysis/ && go test -v -cover

test_mt:
	cd pkg/mt/ && go test -v -cover

test_matout:
	cd pkg/matout/ && go test -v -cover

test_petrinet:
	cd pkg/petrinet/ && go test -v -cover

test_parser:
	cd pkg/parser/ && go test -v -cover

test_mxgraph:
	cd pkg/mxgraph/ && go test -v -cover

test_benchmark:
	cd test/ &&\
	go test -v -bench GoSPN -benchmem -o pprof/test.bin  -cpuprofile pprof/cpu.out . &&\
	go tool pprof --svg pprof/test.bin pprof/cpu.out > pprof/test.svg

build_all: build build_linux build_darwin build_windows

build_linux: deps
	mkdir -p bin/linux
	GOOS=linux GOARCH=amd64 go build -ldflags "$(GOSPN_LDFLAGS)" -o bin/linux/gospn cmd/gospn.go
	cd bin/linux && tar -czvf gospn-$(GOSPN_VERSION)-linux-amd64.tar.gz gospn

build_darwin: deps
	mkdir -p bin/darwin
	GOOS=darwin GOARCH=amd64 go build -ldflags "$(GOSPN_LDFLAGS)" -o bin/darwin/gospn cmd/gospn.go
	cd bin/darwin && tar -czvf gospn-$(GOSPN_VERSION)-darwin-amd64.tar.gz gospn
	GOOS=darwin GOARCH=arm64 go build -ldflags "$(GOSPN_LDFLAGS)" -o bin/darwin/gospn cmd/gospn.go
	cd bin/darwin && tar -czvf gospn-$(GOSPN_VERSION)-darwin-arm64.tar.gz gospn

build_windows: deps
	mkdir -p bin/windows
	GOOS=windows GOARCH=amd64 go build -ldflags "$(GOSPN_LDFLAGS)" -o bin/windows/gospn.exe cmd/gospn.go
	cd bin/windows && zip gospn-$(GOSPN_VERSION)-windows-amd64.zip gospn.exe

clean:
	rm -fR bin/*


