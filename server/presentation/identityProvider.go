package presentation

import (
	"encoding/json"
	"log"
	"net/http"
	"oauth2_api/application"
	"oauth2_api/domain"
	"oauth2_api/presentation/dtos"
	"strconv"

	"go.mongodb.org/mongo-driver/mongo"
)

func IDPWithDBHandlerFactory(database *mongo.Database) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}
}

func getUserFromRequest(r *http.Request) domain.User {
	r.ParseForm()
	password := r.Form.Get("password")
	idParam := r.Form.Get("userId")
	userId, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		log.Println(err)
	}
	idForDb := domain.UserId(userId)
	user := domain.User{UserId: idForDb, Password: password}
	return user
}

func RegisterIdentityHandlerFactory(regAdapter application.IDManager) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		user := getUserFromRequest(r)
		res := application.RegisterUserPassword(regAdapter, user)
		json.NewEncoder(w).Encode(dtos.InsertedOne{InsertedId: res})
	}
}

func GetJWTIdentityHandlerFactory(im application.IDManager, tr application.TokenReader) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		user := getUserFromRequest(r)
		token := application.ProvideToken(user, im, tr)
		if token == nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		tokenDTO := dtos.Token{AccessToken: *token}
		json.NewEncoder(w).Encode(tokenDTO)
	}
}
