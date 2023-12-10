package adapters

import (
	"math/rand"
	"oauth2_api/domain"
	"time"
)

type Reader struct{}

func (r Reader) ReadToken(id domain.UserId) domain.TokenData {
	expires := time.Now().Unix()
	return domain.TokenData{
		Uid: id, Username: "Neo",
		Roles: []string{}, Expires: int(expires),
	}
}

func StoreToken(token domain.TokenData) int {
	id := rand.Uint32()
	return int(id)
}
