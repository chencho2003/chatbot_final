package controller

import (
	"database/sql"
	"encoding/json"
	"myProject/model"
	"myProject/utils/httpResp"
	"net/http"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)


func Adduser(w http.ResponseWriter, r *http.Request) {
	var stud model.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&stud)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(stud.Password), bcrypt.DefaultCost)
	if err != nil {
		httpResp.RespondWithError(w, http.StatusInternalServerError, "failed to hash password")
		return
	}

	stud.Password = string(hashedPassword)

	saveErr := stud.Create()
	if saveErr != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, saveErr.Error())
		return
	}

	httpResp.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "User data added"})
}

// use capital function Name to be able call outside
var admin model.User

func Loginhandler(w http.ResponseWriter, r *http.Request) {
	const (
		StatusMyCustomCode = 480
	)

	var admin model.User
	if err := json.NewDecoder(r.Body).Decode(&admin); err != nil {
		httpResp.RespondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	defer r.Body.Close()

	email := admin.Email
	var admin2 model.User
	loginErr := admin2.Check(email)

	if loginErr != nil {
		switch loginErr {
		case sql.ErrNoRows:
			httpResp.RespondWithError(w, http.StatusUnauthorized, "invalid login")
		default:
			fmt.Println(loginErr)
			httpResp.RespondWithError(w, http.StatusBadRequest, "error in database")
		}
		return
	}

	// Compare the provided password with the stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(admin2.Password), []byte(admin.Password))
	if err != nil {
		httpResp.RespondWithError(w, http.StatusUnauthorized, "invalid login")
		return
	}

	cookie := http.Cookie{
		Name:    "admin-cookie",
		Value:   admin2.Avatar + "=" + admin2.Username,
		Expires: time.Now().Add(30 * time.Minute),
		Secure:  true,
	}
	// Set cookie and send it back to the client
	http.SetCookie(w, &cookie)

	if admin.Email == "admin@gmail.com" {
		httpResp.RespondWithError(w, StatusMyCustomCode, "admin")
		return
	}

	httpResp.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "successful"})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w,&http.Cookie{
		Name : "admin-cookie",
		Expires: time.Now(),
	})
	fmt.Println("logout successful")
	httpResp.RespondWithJSON(w,http.StatusOK,map[string]string{"message":"logout successful"})
}


//////////////////////////////////
