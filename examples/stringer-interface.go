package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var res []string
	for _, b := range ip {
		res = append(res, strconv.Itoa(int(b)))
	}
	return strings.Join(res, ",")
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
