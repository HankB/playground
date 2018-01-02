package main

import (
	"os"
	"log"
	"crypto/md5"
	"fmt"
	"io"
)

func main() {
	file, err := os.Open("filelen.go") // For read access.
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	info, err := file.Stat();
	if err != nil {
		log.Fatal(err)
	}
	length := info.Size()

	print("length of file is ", length, "\n")

	h := md5.New()
	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%x\n", h.Sum(nil))

}