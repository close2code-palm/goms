package application

import (
	"oauth2_api/adapters"
	"oauth2_api/domain"
	"strings"
)

type IDManager interface {
	SaveUsersPass(domain.User) interface{}
	ReadUsersPass(domain.UserId) string
}

func RegisterUserPassword(r IDManager, u domain.User) interface{} {
	salt := adapters.GenerateSalt(adapters.SaltLength)
	saltedPass := adapters.EncryptPassword(u.Password, salt)
	dbUser := domain.User{UserId: u.UserId, Password: saltedPass}
	return r.SaveUsersPass(dbUser)
}

func checkUser(im IDManager, user domain.User) bool {
	passwordInDB := im.ReadUsersPass(user.UserId)
	if passwordInDB == domain.Null {
		return false
	}
	saltSlice := strings.Split(passwordInDB, "$")
	saltedPass := adapters.EncryptPassword(user.Password, saltSlice[0])
	return saltedPass == passwordInDB
}

func ProvideToken(user domain.User, im IDManager, tr TokenReader) *domain.TokenData {
	if checkUser(im, user) {
		token := MakeIntrospection(tr, user.UserId)
		return &token
	}
	return nil
}
