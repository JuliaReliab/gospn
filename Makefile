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

clean:
	rm -fR bin/*


