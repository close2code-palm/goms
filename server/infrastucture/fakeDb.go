package infrastucture

import "oauth2_api/domain"

type FakeDatabase struct {
	PassHashStorage map[domain.UserId][]byte
}

func (fd *FakeDatabase) StoreHash(uid domain.UserId, hash []byte) {
	fd.PassHashStorage[uid] = hash
}

func (fd *FakeDatabase) CheckHash(uid domain.UserId, hash []byte) bool {
	return string(fd.PassHashStorage[uid]) == string(hash)
}
