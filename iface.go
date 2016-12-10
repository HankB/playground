package main

import (
	//"errors"
	"fmt"
	"net"
	"strings"
)

func main() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
	} else {
		// fmt.Println(addrs)
		for _, addr := range addrs {
			// fmt.Println("from:", addr.String(), addr.Network())
			if strings.Compare(addr.String(), "127.0.0.1/8") == 0 ||
				strings.Compare(addr.String(), "::1/128") == 0 {
				fmt.Println("skipping lo:", addr)
			} else {
				addrIP, ipNet, _ := net.ParseCIDR(addr.String())
				bcast := ipNet.IP
				for i := range ipNet.Mask {
					bcast[i] |= ^ipNet.Mask[i]
				}
				if addrIP.DefaultMask() != nil {
					fmt.Println("IPV4 ", addr, "ip is", addrIP, "net is", ipNet.IP, "mask is", ipNet.Mask, "bcast", bcast)
				} else {
					fmt.Println("IPV6 ", addr, "ip is", addrIP, "net is", ipNet.IP, "mask is", ipNet.Mask, "bcast", bcast)
				}
			}
		}
	}
}
