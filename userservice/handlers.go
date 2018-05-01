package userservice

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dillonerhardt/rest-api-golang/utils"
	"gopkg.in/mgo.v2/bson"

	"github.com/dillonerhardt/rest-api-golang/models"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func getAuthorizedUserHandler(w http.ResponseWriter, r *http.Request) {
	u := context.Get(r, "user")
	data, err := json.Marshal(u)
	if err != nil {
		utils.InternalServerErrorWriter(w)
		return
	}
	utils.WriteJSON(w, http.StatusOK, data)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	u := models.User{ID: bson.ObjectIdHex(id)}
	err := u.GetUser()
	if err != nil {
		utils.WriteJSON(w, http.StatusOK, []byte(`{"success": false, "message": "User not found"}`))
		return
	}
	data, err := json.Marshal(u)
	if err != nil {
		utils.InternalServerErrorWriter(w)
		return
	}
	utils.WriteJSON(w, http.StatusOK, data)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "UpdateUserHandler not implemented")
}
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "DeleteUserHandler not implemented")
}
