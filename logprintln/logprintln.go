package main

import (
	"fmt"
	"log"
)

const requiredPri = 5

// LogPrintln outputs remaining arguments according to the value
// of the first argument.
func LogPrintln(pri int, a ...interface{}) {
	if pri < requiredPri {
		log.Println(a...)
	}
	return
}

// FmtPrintln outputs remaining arguments according to the value
// of the first argument.
func FmtPrintln(pri int, a ...interface{}) (n int, err error) {
	if pri < requiredPri {
		return fmt.Println(a...)
	}
	return 0, nil
}

func main() {
	fmt.Println("This is a message")
	LogPrintln(3, "This is a message")
	FmtPrintln(3, "This is a message")

	fmt.Println("message", "with", "multiple", "arguments")
	LogPrintln(3, "message", "with", "multiple", "arguments")
	FmtPrintln(3, "message", "with", "multiple", "arguments")
}
