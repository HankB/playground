package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

func usage() {
	fmt.Println("Usage:", os.Args[0], "nnnn")
	fmt.Println("'nnnn' => port number")
	os.Exit(1)
}

/* A Simple function to verify error */
/* copied from https://varshneyabhi.wordpress.com/2014/12/23/simple-udp-clientserver-in-golang/ */
func checkErr(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}
	port, err := strconv.Atoi(os.Args[1])
	if err != nil || port <= 0 || port > 0xFFFF {
		fmt.Println("Cannot interpret", os.Args[1], "as a port")
		usage()
	}
	fmt.Println("port:", port)
        server := fmt.Sprintf(":%d", port)
        //server := fmt.Sprintf("255.255.255.0:%d", port)
	addr, err := net.ResolveUDPAddr("udp", server)
	checkErr(err)
	serverConn, err := net.ListenUDP( "udp", addr)
	checkErr(err)
	defer serverConn.Close()
	fmt.Println(serverConn)

	buf := make([]byte, 1024)

	for {
		n, addr, err := serverConn.ReadFromUDP(buf)
		checkErr(err)
		fmt.Println("recv: \"%s\" from \"%s\" \n", string(buf[:n]), addr)
	}

}
//func ListenMulticastUDP(network string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)

