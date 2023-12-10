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

func RegisterIdentityHandlerFactory(regAdapter application.Registrator) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		r.ParseForm()
		password := r.Form.Get("password")
		idParam := r.Form.Get("userId")
		userId, err := strconv.ParseUint(idParam, 10, 32)
		if err != nil {
			log.Println(err)
		}
		idForDb := domain.UserId(userId)
		user := domain.User{UserId: idForDb, Password: password}
		res := application.RegisterUserPassword(regAdapter, user)
		json.NewEncoder(w).Encode(dtos.InsertedOne{InsertedId: res})
	}
}
