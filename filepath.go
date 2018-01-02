package main

import (
//	"log"
	"os"
	"path/filepath"
)

func myWalkFunc(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		print("dir: ", path, "\n")
	} else if info.Mode() & os.ModeType != 0 {
		print("other: ", path, "\n")
	} else {
		print("len: ", info.Size(), " ", path, "\n")
		}
	return nil
}

func main() {
// func Walk(root string, walkFn WalkFunc) error
	filepath.Walk("/home/hbarta/Downloads", myWalkFunc)
}