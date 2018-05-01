package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dillonerhardt/rest-api-golang/models"
	"github.com/dillonerhardt/rest-api-golang/utils"
	"github.com/gorilla/context"
	"gopkg.in/mgo.v2/bson"

	jwt "github.com/dgrijalva/jwt-go"
)

// ValidateJWT verifies the JWT
func ValidateJWT(skip bool, handler http.Handler) http.Handler {
	// Skip the jwt verification
	if skip {
		return handler
	}
	// return the middleware function
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get authorization header
		authHeader := r.Header.Get("authorization")
		if authHeader == "" { // check it's not empty
			utils.WriteJSON(w, http.StatusBadRequest,
				[]byte(`{"success": false, "message": "Authorization header required"}`))
			return
		}
		// split header value
		bearer := strings.Split(authHeader, " ")
		if len(bearer) != 2 { // make sure its format looks correct
			utils.WriteJSON(w, http.StatusBadRequest,
				[]byte(`{"success": false, "message": "Malformatted authorization header"}`))
			return
		}
		// parse the token
		token, err := jwt.Parse(bearer[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return []byte(utils.Config.JWT.Secret), nil
		})
		if err != nil {
			utils.InternalServerErrorWriter(w)
			return
		}
		if !token.Valid { // make sure token is valid
			utils.WriteJSON(w, http.StatusUnauthorized,
				[]byte(`{"success": false, "message": "Authorization token invalid"}`))
			return
		}
		// get id from token
		id := token.Claims.(jwt.MapClaims)["id"].(string)
		// get user data
		u := models.User{ID: bson.ObjectIdHex(id)}
		err = u.GetUser()
		if err != nil {
			log.Panic(err)
			utils.InternalServerErrorWriter(w)
			return
		}
		// set user data on request
		context.Set(r, "user", u)
		// run handler
		handler.ServeHTTP(w, r)
	})
}
