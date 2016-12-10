package main

import (
	//"errors"
	"fmt"
	"net"
	"strings"
)

func GetBroadcastIPs(includeIPV6 bool) ([]net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	var bcast_addrs []net.IP
	if err != nil {
		return nil, err
	} else {
		for _, addr := range addrs {
			if strings.Compare(addr.String(), "127.0.0.1/8") == 0 ||
				strings.Compare(addr.String(), "::1/128") == 0 {
			} else {
				addrIP, ipNet, _ := net.ParseCIDR(addr.String())
				bcast := ipNet.IP
				for i := range ipNet.Mask {
					bcast[i] |= ^ipNet.Mask[i]
				}
				if addrIP.DefaultMask() != nil {
					bcast_addrs = append(bcast_addrs, bcast)
					// fmt.Println("IPV4 ", addr, "ip is", addrIP, "net is", ipNet.IP, "mask is", ipNet.Mask, "bcast", bcast)
				} else if includeIPV6 {
					bcast_addrs = append(bcast_addrs, bcast)
					// fmt.Println("IPV6 ", addr, "ip is", addrIP, "net is", ipNet.IP, "mask is", ipNet.Mask, "bcast", bcast)
				}
			}
		}
	}
return bcast_addrs, nil
}

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
	ba, _ := GetBroadcastIPs(true)
	fmt.Println("From get_broadcast_addrs", ba)
	ba, _ = GetBroadcastIPs(false)
	fmt.Println("From get_broadcast_addrs w/out IPV6", ba)
}
