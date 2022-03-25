package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var s []string
	for _, _v := range ip {
		v := _v
		s = append(s, strconv.Itoa(int(v)))
	}

	return fmt.Sprintf(strings.Join(s, "."))
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
