go test -bench GoSPN -benchmem -o pprof/test.bin  -cpuprofile pprof/cpu.out .
go tool pprof --svg pprof/test.bin pprof/cpu.out > pprof/test.svg
go tool pprof gospn_test.go pprof/cpu.out
