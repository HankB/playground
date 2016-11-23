package main

import (
	"fmt"
)

// analog of libc strlen on a []byte
func strlen(s []byte) int {
	i := 0
	for ; i < len(s); i++ {
		if s[i] == 0 {
			break
		}
	}
	return i
}

const fill = 16 // length to fill string 16 bytes/line

// endLine has two jobs. First pad remainder of line if <16
// bytes have been printed and then pad and add the ascii
// representation (and '\n') This is part of the code that will print
// binary strings using a format like 'hexdump -C'
/*
hbarta@swanky:~/Documents/go-wks/src/github.com/HankB/playground$ hexdump -C ~/go/bin/go|tail -3
0097e1f0  5f 73 74 61 72 74 00 63  72 6f 73 73 63 61 6c 6c  |_start.crosscall|
0097e200  5f 61 6d 64 36 34 00                              |_amd64.|
0097e207
hbarta@swanky:~/Documents/go-wks/src/github.com/HankB/playground$
*/
func endLine(str []byte) {
	fmt.Printf("len = %d, strlen = %d\n", len(str), strlen(str))
  /*
	add := fill - strlen(str)
	for i := 0; i < add; i++ {
		str = append(str, ' ')
	}
  */
	fmt.Printf("|%s|\n", str)
  fmt.Println("|0123456789ABCDEF|")
}

//const str = "abc"
var str []byte = make([]byte, 8) // = { '\x68', '\x64' }
func main() {
	str[0] = 'a'
	str[1] = 'b'
	str[2] = 'c'

	endLine(str)
}
