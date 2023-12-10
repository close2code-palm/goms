package application

import (
	"oauth2_api/domain"
)

type CreateToken func(domain.TokenData) int

func (ct CreateToken) AddToken(token domain.TokenData) int {
	return ct(token)
}
