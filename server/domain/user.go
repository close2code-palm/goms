package domain

type User struct {
	UserId   UserId
	Password string `bson:"saltedPassword"`
}
