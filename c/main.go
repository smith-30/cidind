package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://service-a:8011/")
	fmt.Printf("err %#v\n", err)
	fmt.Printf("err %#v\n", resp)
}
