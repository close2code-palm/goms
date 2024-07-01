package adapters

import (
	"fmt"

	"golang.org/x/exp/maps"

	"github.com/golang-jwt/jwt/v5"
)

func MakeSignedJWTWithClaims(attributes map[string]interface{}, signKey string) string {
	claimsMap := jwt.MapClaims{}
	maps.Copy(claimsMap, attributes)
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claimsMap)
	// var signInterface interface{} = signKey
	signed, err := token.SignedString([]byte(signKey))
	if err != nil {
		fmt.Printf("Got signing error!")
	}
	return signed
}
