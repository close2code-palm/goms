package application

import (
	"oauth2_api/adapters"
	"oauth2_api/domain"
)

type Registrator interface {
	Register(domain.User) interface{}
}

func RegisterUserPassword(r Registrator, u domain.User) interface{} {
	saltedPass := adapters.EncryptPassword(u.Password, 10)
	dbUser := domain.User{UserId: u.UserId, Password: saltedPass}
	return r.Register(dbUser)
}
