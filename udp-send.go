package main

import (
    "os"
    "fmt"
    "strconv"
    )

func usage() {
	fmt.Println("Usage:", os.Args[0], "nnnn")
	fmt.Println("'nnnn' => port number")
	os.Exit(1)
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
}
