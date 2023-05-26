package parser

import (
	"fmt"
	"testing"
)

func TestPNreader1(t *testing.T) {
	net, imark, err := PNreadFromFile("../../example/spnp_example1.spn")
	if err != nil {
		t.Error("error")
	}
	fmt.Println(net)
	fmt.Println(imark)
}

func TestPNreader2(t *testing.T) {
	net, imark, err := PNreadFromFile("../../example/spnp_example2.spn")
	if err != nil {
		t.Error("error")
	}
	fmt.Println(net)
	fmt.Println(imark)
}

func TestPNreader3(t *testing.T) {
	net, imark, err := PNreadFromFile("../../example/spnp_example3.spn")
	if err != nil {
		t.Error("error")
	}
	fmt.Println(net)
	fmt.Println(imark)
}

func TestPNreader4(t *testing.T) {
	net, imark, err := PNreadFromFile("../../exmple/spnp_example4.spn")
	if err != nil {
		t.Error("error")
	}
	fmt.Println(net)
	fmt.Println(imark)
}

func TestPNreader5(t *testing.T) {
	net, imark, err := PNreadFromFile("../../example/spnp_example5.spn")
	if err != nil {
		t.Error("error")
	}
	fmt.Println(net)
	fmt.Println(imark)
}

func TestPNreader6(t *testing.T) {
	net, imark, err := PNreadFromFile("../../example/spnp_example6.spn")
	if err != nil {
		t.Error("error")
	}
	fmt.Println(net)
	fmt.Println(imark)
}

func TestPNreader7(t *testing.T) {
	net, imark, err := PNreadFromFile("../../example/iaas_cloud.spn")
	if err != nil {
		t.Error("error")
	}
	fmt.Println(net)
	fmt.Println(imark)
}

func TestPNreader8(t *testing.T) {
	net, imark, err := PNreadFromFile("../../example/raid6.spn")
	if err != nil {
		t.Error("error")
	}
	fmt.Println(net)
	fmt.Println(imark)
}

func TestPNreader9(t *testing.T) {
	net, imark, err := PNreadFromFile("../../example/raid10.spn")
	if err != nil {
		t.Error("error")
	}
	fmt.Println(net)
	fmt.Println(imark)
}
