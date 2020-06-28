package main

import (
	"fmt"
	"net/http"

	"github.com/smith-30/cidind/rdb"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "a")
}

func main() {
	_, err := rdb.Mysql("root", "MYSQL_ADMIN", "0.0.0.0", "3306", "testdb", true)
	if err != nil {
		panic(err)
	}
	http.HandleFunc("/", handler) // ハンドラを登録してウェブページを表示させる
	http.ListenAndServe("0.0.0.0:8011", nil)
}
