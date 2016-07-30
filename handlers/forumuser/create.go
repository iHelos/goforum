package forumuser

import (
	"net/http"
	_ "github.com/lib/pq"
	"github.com/iHelos/goforum/handlers"
	"encoding/json"
	"github.com/iHelos/goforum/database"
	"fmt"
	//"log"
	//"github.com/bmizerany/pq"
	"github.com/lib/pq"
)

const (
	UriCreate = "user/create"
)

type createParams struct {
	ID int
	Username string
	About string
	Name string
	Email string
	IsAnonymous bool `json:"isAnonymous, omitempty"`
}

func getParamsCreate(r *http.Request) (createParams, error){
	var params createParams
	err := json.NewDecoder(r.Body).Decode(&params)
	return params, err
}

func createUser(params *createParams) (interface {}, int){
	db := database.GetDB()
	var query = fmt.Sprintf(
		"INSERT INTO %s (%s, %s, %s, %s, %s) VALUES ($1,$2,$3,$4,$5) RETURNING id",
		database.TableUser,
		database.User_name,
		database.User_email,
		database.User_username,
		database.User_about,
		database.User_isAnonymous,
	)
	var id int
	err := db.QueryRow(
		query,
		params.Name,
		params.Email,
		params.Username,
		params.About,
		params.IsAnonymous,
	).Scan(&id)
	if err!=nil{
		switch err.(*pq.Error).Code {
		case "23505":
			return nil, handlers.StatusUserCreated
		default:
			print(err.(*pq.Error).Code)
			//panic(err)
			return nil, handlers.StatusUnknown
		}
	}
	params.ID = id
	return params, handlers.StatusOK
}

func UserCreateHandler(w http.ResponseWriter, r *http.Request) {
	params, err := getParamsCreate(r)
	if err != nil{
		handlers.WriteAnswer(w, "", handlers.StatusNotValid)
		return
	}
	result, status := createUser(&params)
	handlers.WriteAnswer(w, result, status)
}