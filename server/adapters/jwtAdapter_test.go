package adapters

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

func TestMakeJWTWithClaims(t *testing.T) {
	key := "testingSecret"
	claims := map[string]interface{}{"access_level": 1, "server": "rus_21"}
	token := MakeSignedJWTWithClaims(claims, key)
	t.Log(token)
	resultClaims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, resultClaims, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	for k, v := range resultClaims {
		t.Logf("%s: %s", k, v)
		f, ok := v.(float64)
		if ok {
			if int(f) != claims[k] {
				t.Error("Not valid claim")
			}
		} else if claims[k] != v {
			t.Error("Not valid claim")
		}
	}
}
