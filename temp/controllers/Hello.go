package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/shubham1010/project/temp/models"
	"github.com/shubham1010/project/temp/common"
	"github.com/shubham1010/project/temp/database"
	"github.com/shubham1010/project/temp/password"

	log "github.com/sirupsen/logrus"
)


var t models.User


func Login(w http.ResponseWriter, r *http.Request) {
	if(r.Method!="POST") {
		log.Warn("Expected POST Request for Login")
		return
	}
	w.Header().Set("Context-type", "application/json")
	var checkUser models.User
	err := json.NewDecoder(r.Body).Decode(&checkUser)
	if err!=nil {
		log.Fatal("JSON Decoding Not Happened")
		return
	}

	var dbData models.User
	db := common.GetDBConnection()

	dbData  = database.CheckUserExist(db, checkUser)


	if (checkUser.Username==dbData.Username && password.CheckPasswordHash(checkUser.Pass, dbData.Pass)) {
		http.Redirect(w, r, "/info", http.StatusSeeOther)
	}else {
		var JErr models.JSONErr
		JErr.Error = "User Authentication Failed"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(JErr)
	}

}

func Info(w http.ResponseWriter, r *http.Request) {
	t.Id=1011
	t.Name="This is testing"
	t.Email="abc@gmail.com"
	HashPass, err := password.HashPassword("1010")
	if err!=nil {
		log.Fatal("Password is not Hashed...")
	}
	t.Pass = HashPass
	w.Header().Set("Context-type","application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(t)

}


func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-type","application/json")
	var newUser models.User

	err := json.NewDecoder(r.Body).Decode(&newUser)

	if err != nil {
		log.Warn("[CONTROLLERS]: SignUp Invalid User Data")
	}

	var dbData models.User
	db := common.GetDBConnection()

	dbData  = database.CheckUserExist(db, newUser)

	if (dbData.Username==newUser.Username) {
		var JSON models.JSONErr
		JSON.Error = "Username already exist please enter other username"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(JSON)
	}else {
		err := database.InsertNewUser(db, newUser)
		if err!=nil {
			log.Fatal("Insertion failed for new user")
		}
		var JSON models.JSONSuccess
		JSON.Success = "NewUser Created"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(JSON)
	}

}
