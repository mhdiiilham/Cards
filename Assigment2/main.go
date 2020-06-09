package main

import (
	"os"
	"io"
)

func main() {
	filePath := os.Args[1]
	file,_ := os.Open(filePath)
	io.Copy(os.Stdout, file)
}