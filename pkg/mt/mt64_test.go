package mt

import (
	"fmt"
	"testing"
)

func TestMT64(t *testing.T) {
	mt := NewMT64()
	mt.InitByArray([]uint64{0x12345, 0x23456, 0x34567, 0x45678})
	for i := 0; i < 1000; i++ {
		fmt.Println(mt.Float64())
	}
}
