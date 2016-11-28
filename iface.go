package main

import (
	//"errors"
	"fmt"
	"net"
)

func main() {

	addr, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(addr)
	}
}
