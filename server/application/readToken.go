package application

import (
	"oauth2_api/domain"
)

type TokenReader interface {
	ReadToken(domain.UserId) domain.TokenData
}

func MakeIntrospection(tr TokenReader, id domain.UserId) domain.TokenData {
	return tr.ReadToken(id)
}
