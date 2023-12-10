package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"oauth2_api/adapters"
	"oauth2_api/infrastucture"
	"oauth2_api/presentation"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	config := infrastucture.ReadConfig("config.yaml")
	mainCtx := context.Background()

	mongoCtx, cancel := context.WithCancel(mainCtx)
	defer cancel()
	mongoUrl := fmt.Sprintf("mongodb://%s:%d", config.MongoConfig.Host, config.MongoConfig.Port)
	client, err := mongo.Connect(mongoCtx, options.Client().ApplyURI(mongoUrl))
	defer func() {
		if err = client.Disconnect(mongoCtx); err != nil {
			panic(err)
		}
	}()
	database := client.Database(config.MongoConfig.Database)

	tokenReader := adapters.Reader{}
	registrationAdapter := adapters.MongoAdapter{Col: database.Collection("users")}
	tokenHandler := presentation.GetJWTIdentityHandlerFactory(registrationAdapter, tokenReader)
	http.Handle("/register", presentation.RegisterIdentityHandlerFactory(registrationAdapter))
	http.Handle("/token", tokenHandler)
	log.Fatal(http.ListenAndServe(":8088", nil))
}
