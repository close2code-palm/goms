package adapters

import (
	"context"
	"log"
	"oauth2_api/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoAdapter struct {
	Col *mongo.Collection	
}

func (ma MongoAdapter) SaveUsersPass(u domain.User) interface{} {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := ma.Col.InsertOne(ctx, bson.M{
		"userId":         u.UserId,
		"saltedPassword": u.Password})
	if err != nil {
		log.Println(err)
	}
	return res.InsertedID
}

func (ma MongoAdapter) ReadUsersPass(uid domain.UserId) string {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var user domain.User
	err := ma.Col.FindOne(ctx, bson.M{"userId": uid}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Null
		}
	}
	return user.Password
}
