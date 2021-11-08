package domain

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type Chaim struct {
	Id   string
	Name string
}

var hash = os.Getenv("APP_TOKEN")

func (c Chaim) GenerateJWT() string {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   c.Id,
		"aud":  c.Id,
		"name": c.Name,
		"nbf":  time.Now().UTC().Unix(),
		"exp":  time.Now().UTC().Add(time.Hour * 24).Unix(),
	})

	stringToken, _ := token.SignedString([]byte(hash))

	return stringToken
}

func Validate(tokenJwt string) (Chaim, error) {
	secret := []byte(hash)

	token, err := jwt.Parse(tokenJwt, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		return Chaim{}, errors.New("Invalid JWT Token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return Chaim{
			Id:   claims["aud"].(string),
			Name: claims["name"].(string),
		}, nil
	}

	return Chaim{}, errors.New("Invalid JWT Token")
}
