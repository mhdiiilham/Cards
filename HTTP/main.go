package main

import (
	"os"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	bs := make([]byte, 9999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}