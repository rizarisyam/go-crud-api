package handlers

import (
	"encoding/json"
	"strconv"

	// "fmt"
	"go-crud-article/connection"
	"go-crud-article/structs"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	// "time"
	// "github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	payloads, _ := ioutil.ReadAll(r.Body)

	var user structs.User
	var riskProfile structs.RiskProfile

	json.Unmarshal(payloads, &user)

	connection.DB.Create(&user)
	if user.Age >= 30 {
		riskProfile.Userid = user.ID
		riskProfile.Stock = 72.5
		riskProfile.Bond = 21.5
		riskProfile.MM = 100 - (72.5 + 21.5)
	}

	if user.Age >= 20 {
		riskProfile.Userid = user.ID
		riskProfile.Stock = 54.5
		riskProfile.Bond = 24.5
		riskProfile.MM = 100 - (54.5 + 24.5)
	}

	if user.Age < 20 {
		riskProfile.Userid = user.ID
		riskProfile.Stock = 34.5
		riskProfile.Bond = 45.5
		riskProfile.MM = 100 - (34.5 + 45.5)
	}
	connection.DB.Create(&riskProfile)
	res := structs.Result{Code: 200, Data: user, Message: "Success create user"}

	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	page := 1
	take := 1

	query := r.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "take":
			take, _ = strconv.Atoi(queryValue)
			break

		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		}
	}
	// vars := mux.Vars(r)
	// articleID := vars["id"]
	offset := (page - 1) * take
	// log.Println("Connection failed")

	user := []structs.User{}
	// // connection.DB.First(&article, articleID)
	connection.DB.Limit(take).Offset(offset).Find(&user)

	res := structs.Result{Code: 200, Data: user, Message: "Success get article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

func getUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userid := vars["id"]

	// var user structs.User
	var riskProfile structs.RiskProfile

	// connection.DB.First(&user, userid)
	connection.DB.Where("userid = ?", userid).Find(&riskProfile)

	res := structs.Result{Code: 200, Data: riskProfile, Message: "Success get article"}
	result, err := json.Marshal(res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
