package main

import (
	"fmt"
	"net/http"
	//"./database"
	//"./handlers/forumuser"
	"github.com/iHelos/goforum/handlers/common"
	"github.com/iHelos/goforum/handlers/forumuser"
)

const prefix = "/db/api/"

func main() {
	//common functions
	http.HandleFunc(prefix + common.UriClear, common.ClearHandler)
	http.HandleFunc(prefix + common.UriStatus, common.StatusHandler)

	//user functions
	http.HandleFunc(prefix + forumuser.UriCreate, forumuser.UserCreateHandler)

	http.ListenAndServe(":8032", nil)

	fmt.Printf("hey")
}