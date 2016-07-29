package common

import (
	"fmt"
	"net/http"
	//"database/sql"
	"github.com/iHelos/goforum/database"
	_ "github.com/bmizerany/pq"
)

const (
	UriClear = "clear"
)

func Clear() {
	db := database.GetDB()
	rows, err := db.Query("truncate forumuser cascade")
	if err != nil {
		fmt.Printf("damn\n")
		panic(err)
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		count++
	}
	fmt.Printf("length: %d \n", count)
}

func ClearHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("GoForum", "A Go WebServer for TP database")
	w.Header().Set("Content-Type", "application/json")

	Clear()

	hostname := r.URL.Query()["hostname"]
	//w.Write([]byte(hostname[0]))
	fmt.Printf("%s", hostname)
}
