package main

import (
	"fmt"
	"net/http"
	//"./database"
	//"./handlers/forumuser"
	"./handlers/common"
)

const prefix = "/db/api/"

func main() {
	common.Clear()

	http.HandleFunc(prefix + common.UriClear, common.ClearHandler)

	http.ListenAndServe(":8032", nil)

	fmt.Printf("hey")
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "A Go WebServer")
	w.Header().Set("Content-Type", "text/html")

	common.Clear()

	hostname := r.URL.Query()["hostname"]
	//w.Write([]byte(hostname[0]))
	fmt.Printf("%s", hostname)
}