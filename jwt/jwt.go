// jwt package includes some examples about how to code with golang-jwt
package jwt

import (
	"encoding/json"
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

// GenerateToken to get a token string
func GenerateToken(data any, key any) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{
		"data": data,
	})

	// to get the signed string from token needs a slice
	bs, err := json.Marshal(key)
	if err != nil {
		return "", err
	}

	ts, err := token.SignedString(bs)
	return ts, err
}

// ParseToken to parse a token
func ParseToken(token string, key any) (string, error) {
	// jwt.MapClaims is base on map type, so no need to use pointer here
	t, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return json.Marshal(key)
	})
	if err != nil {
		return "", err
	}

	if claim, ok := t.Claims.(jwt.MapClaims); ok {
		if s, ok := claim["data"].(string); ok {
			return s, nil
		}
		return "", errors.New("parse token failed, claim data is not a string")
	}
	return "", errors.New("parse token failed, invalid claim")
}
