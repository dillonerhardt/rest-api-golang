package authservice

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dillonerhardt/rest-api-golang/utils"

	"github.com/dillonerhardt/rest-api-golang/models"
	"gopkg.in/mgo.v2/bson"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	// get user data from body
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		utils.InternalServerErrorWriter(w)
		return
	}
	// store submitted password seperatley
	password := u.Password
	// get user from database
	err = u.GetUserByIdentifier()
	if err != nil {
		utils.WriteJSON(w, http.StatusUnauthorized,
			[]byte(`{"success": false, "message": "Incorrect credentials"}`))
	}
	// check passoword is correct
	if !utils.VerifyPassword(password, u.Password) {
		utils.WriteJSON(w, http.StatusUnauthorized,
			[]byte(`{"success": false, "message": "Incorrect credentials"}`))
		return
	}
	// create token
	token, err := createJWT(u)
	if err != nil {
		utils.InternalServerErrorWriter(w)
		return
	}
	utils.WriteJSON(w, http.StatusOK,
		[]byte(fmt.Sprintf(`{"success": true, "message": "User created", "token": "%s"}`, token)))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	// get user data from body
	var u models.User
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		utils.InternalServerErrorWriter(w)
		return
	}
	// Set objectID
	u.ID = bson.NewObjectId()
	// Create password hash
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		utils.InternalServerErrorWriter(w)
		return
	}
	// Set created data
	u.Created = time.Now()
	err = u.CreateUser()
	if err != nil {
		utils.InternalServerErrorWriter(w)
		return
	}
	token, err := createJWT(u)
	if err != nil {
		utils.InternalServerErrorWriter(w)
		return
	}
	utils.WriteJSON(w, http.StatusOK,
		[]byte(fmt.Sprintf(`{"success": true, "message": "User created", "token": "%s"}`, token)))
}

// create a JWT
func createJWT(u models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": u.ID.Hex(),
	})
	tokenStr, err := token.SignedString([]byte(utils.Config.JWT.Secret))
	return tokenStr, err
}
