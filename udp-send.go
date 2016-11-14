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
        server := fmt.Sprintf("127.0.0.1:%d", port)

	// get server handle
	ServerAddr, err := net.ResolveUDPAddr("udp", server)
	checkErr(err)

	// get our handle
	LocalAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:0")
	checkErr(err)

        // establish connection
	conn, err := net.DialUDP("udp", LocalAddr, ServerAddr)
	checkErr(err)

        defer conn.Close()

        msg := "Hello UDP World!\n"
        buf := []byte(msg)
        _,err = conn.Write(buf)
	checkErr(err)

        fmt.Println("msg sent, Sir!")

}
