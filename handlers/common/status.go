package common

import (
	"github.com/iHelos/goforum/handlers"
	"net/http"
	"database/sql"
	"github.com/iHelos/goforum/database"
	"fmt"
	"log"
)

const (
	UriStatus = "status"
)

type StatusAnswer struct {
	UserCount   int `json:"user"`
	ThreadCount int `json:"thread"`
	ForumCount  int `json:"forum"`
	PostCount   int `json:"post"`
}

func (sa *StatusAnswer) setCount(table database.Tablename, count int) (ok bool)  {
	switch table {
	case database.TableUser:
		sa.UserCount = count
		return true
	case database.TableThread:
		sa.ThreadCount = count
		return true
	case database.TablePost:
		sa.PostCount = count
		return true
	case database.TableForum:
		sa.ForumCount = count
		return true
	default:
		return false
	}
}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	body, status := status()
	handlers.WriteAnswer(w, body, status)
}

func status() (interface{}, int) {

	db := database.GetDB()
	txn, err := db.Begin()
	if err != nil {
		log.Print(err)
		return nil, handlers.StatusUnknown
	}
	result := StatusAnswer{}
	queryTables := []database.Tablename{
		database.TableUser,
		database.TablePost,
		database.TableForum,
		database.TableThread,
	}
	for _, table := range queryTables{
		count, err := statustable(txn, table)
		if err != nil{
			return nil, 5
		}
		result.setCount(table, count)
	}
	return result, 0
}

func statustable(tr *sql.Tx, table database.Tablename) (int, error){
	alias := "tablecount"
	query := fmt.Sprintf("select count(*) as %s from %s", alias, table)
	var id int
	err := tr.QueryRow(query).Scan(&id)
	return id,err
}