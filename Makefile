GOSPN_VERSION=0.9.0

deps:
	go get github.com/antlr/antlr4/runtime/Go/antlr

build: deps
	mkdir -p bin
	go build -o bin/gospn cmd/main.go

test: test_matout test_petrinet test_parser

test_matout:
	cd pkg/matout/; go test -cover; cd ../../

test_petrinet:
	cd pkg/petrinet/; go test -cover; cd ../../

test_parser:
	cd pkg/parser/; go test -cover; cd ../../

test_benchmark:
	cd test/; go test -bench GoSPN -benchmem -o pprof/test.bin  -cpuprofile pprof/cpu.out .; go tool pprof --svg pprof/test.bin pprof/cpu.out > pprof/test.svg; cd ../

build_all: build build_linux build_darwin build_windows

build_linux: deps
	mkdir -p bin/linux
	GOOS=linux GOARCH=amd64 go build -o bin/linux/gospn cmd/main.go
	cd bin/linux; tar -czvf gospn-$(GOSPN_VERSION)-linux-amd64.tar.gz gospn; cd ../../

build_darwin: deps
	mkdir -p bin/darwin
	GOOS=darwin GOARCH=amd64 go build -o bin/darwin/gospn cmd/main.go
	cd bin/darwin; tar -czvf gospn-$(GOSPN_VERSION)-darwin-amd64.tar.gz gospn; cd ../../

build_windows: deps
	mkdir -p bin/windows
	GOOS=windows GOARCH=amd64 go build -o bin/windows/gospn.exe cmd/main.go
	cd bin/windows; zip gospn-$(GOSPN_VERSION)-windows-amd64.zip gospn.exe; cd ../../

clean:
	rm -fR bin/*


