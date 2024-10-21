package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"math/rand"
	"time"
	"voxcon/constant"
)

type Claims struct {
	GameId string `json:"game_id"`
	jwt.RegisteredClaims
}

func GenerateRandomID() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	id := make([]byte, 10)
	for i := range id {
		id[i] = constant.Charset[rand.Intn(len(constant.Charset))]
	}
	return string(id)
}

func CreateToken(gameId string) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(constant.JwtSecretKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
