package infrastucture

import (
	"oauth2_api/domain"
	"testing"
)

func TestFakeDb(t *testing.T) {
	db := FakeDatabase{map[domain.UserId][]byte{}}
	passHash := []byte{2, 3}
	db.StoreHash(1, passHash)
	if !db.CheckHash(1, passHash) {
		t.Error("Hash comparison missed!")
	}
}
