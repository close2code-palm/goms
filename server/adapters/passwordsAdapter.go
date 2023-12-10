package adapters

import (
	"fmt"
	"math/rand"
	"oauth2_api/domain"
	"time"

	"golang.org/x/crypto/argon2"
)

var SaltLength = 10

func EncryptPassword(password string, salt string) string {
	//RFC recommended parameters, memory kilo * kb
	key := argon2.IDKey([]byte(password), []byte(salt), 1, 64*domain.Kilo, 4, 32)
	return fmt.Sprintf("%s$%v", salt, key)
}

func GenerateSalt(saltLength int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	salt := make([]rune, saltLength)
	for i := range salt {
		salt[i] = chars[rand.Intn(len(chars))]
	}
	return string(salt)
}
