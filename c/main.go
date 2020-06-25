package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://service-b:8022/")
	fmt.Printf("err %#v\n", err)
	fmt.Printf("err %#v\n", resp)
}
