package common

import (
	"net/http"
	"database/sql"
	"github.com/iHelos/goforum/database"
//	"github.com/iHelos/goforum/helper"
	_ "github.com/bmizerany/pq"
	"log"
	"github.com/iHelos/goforum/handlers"
)

const (
	UriClear = "clear"
)

func Clear() (interface{}, int){
	db := database.GetDB()
	txn, err := db.Begin()
	if err != nil {
		log.Print(err)
		return nil, handlers.StatusUnknown
	}
	allTables := []database.Tablename{
		database.TableFollow,
		database.TableUser,
		database.TablePost,
		database.TableThread,
		database.TableThread,
	}
	for _, table := range allTables{
		err = clearTable(txn, table)
		if err != nil{
			log.Print(err)
			txn.Rollback()
			return nil, handlers.StatusUnknown
		}
	}
	txn.Commit()
	return "OK", handlers.StatusOK
}

func ClearHandler(w http.ResponseWriter, r *http.Request) {
	body, status := Clear()
	answer, err := handlers.MakeAnswer(status, body)
	if err != nil{
		log.Print(err)
	}
	handlers.WriteAnswer(w, answer)
}

func clearTable(tr *sql.Tx, table database.Tablename) error{
	query := string("truncate " + table +  " cascade")
	row, err := tr.Query(query)
	defer row.Close()
	return err
}