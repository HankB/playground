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
// representation (and '\n')
func endLine(str []byte) {
	fmt.Printf("len = %d, strlen = %d\n", len(str), strlen(str))
	add := fill - strlen(str)
	for i := 0; i < add; i++ {
		str = append(str, ' ')
	}
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
