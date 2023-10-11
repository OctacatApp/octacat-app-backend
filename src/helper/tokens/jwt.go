package tokens

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrClaimsNotPointer     = errors.New("claims must be a pointer")
	ErrInvalidSigningMethod = errors.New("signing method is not valid")
	ErrClaimsTypeNotEquals  = errors.New("claims type is not equals")
)

type JWTResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func NewJWT[C jwt.Claims](claims C, secret []byte) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &tokenString, nil
}

func Validate[C jwt.Claims](tokenString string, secret []byte, claims C) (*jwt.Token, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if _, isOk := t.Method.(*jwt.SigningMethodHMAC); !isOk {
			return nil, ErrInvalidSigningMethod
		}
		return secret, nil
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, keyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetClaims[C jwt.Claims](token *jwt.Token) (C, error) {
	claims, isOk := token.Claims.(C)
	if !isOk {
		return claims, ErrClaimsTypeNotEquals
	}

	return claims, nil
}
