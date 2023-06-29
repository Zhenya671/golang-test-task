package service

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type IMiddleWere interface {
	GenerateToken(uuid string, name string) (string, error)
	ParseToken(token string) (string, string, error)
	generatePasswordHash(password string) string
}

const (
	TLLToken = 5 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserUUID string
	UserName string
}

func (s UserService) GenerateToken(uuid string, name string) (string, error) {
	var claims = tokenClaims{
		UserUUID: uuid,
		UserName: name,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TLLToken).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	return token.SignedString([]byte(s.AppSettings.TokenKey))
}

func (s UserService) ParseToken(token string) (string, string, error) {
	claims := tokenClaims{}
	TokenFunction := func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.AppSettings.TokenKey), nil
	}

	tkn, err := jwt.ParseWithClaims(token, &claims, TokenFunction)
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", "", err
		}
		fmt.Println(err)
		return "", "", err

	}

	if !tkn.Valid {
		return "", "", errors.New("invalid jwt token")
	}

	return claims.UserUUID, claims.UserName, nil
}

func (s UserService) generatePasswordHash(password string) string {
	hash := sha256.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(s.AppSettings.TokenKey)))
}
