package utils

import (
	"fmt"

	"github.com/adhyttungga/bri-life-testteknikbe/config"
	"github.com/golang-jwt/jwt/v5"
)

func ValidateToken(token string, agentId *string) bool {
	key, err := jwt.ParseECPublicKeyFromPEM([]byte(config.Config.PublicKey))
	if err != nil {
		return false
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}

		return key, nil
	})
	if err != nil {
		return false
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return false
	}

	id, _ := claims["dat"].(string)
	*agentId = id
	return true
}
