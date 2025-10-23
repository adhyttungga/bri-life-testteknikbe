package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/adhyttungga/bri-life-testteknikbe/config"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(agentId string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(config.Config.PrivateKey))
	if err != nil {
		return "", fmt.Errorf("jwt.ParseRSAPrivateKeyFromPEM: %w", err)
	}

	ct := time.Now().UTC()
	claims := make(jwt.MapClaims)
	claims["dat"] = agentId
	claims["exp"] = ct.Add(7 * time.Hour).Unix()
	claims["iat"] = ct.Unix()
	claims["nbf"] = ct.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	if err != nil {
		log.Printf("error generate token: %v", err)
		return "", err
	}

	return token, nil
}
