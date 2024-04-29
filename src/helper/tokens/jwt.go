package tokens

import (
	"reflect"

	"github.com/golang-jwt/jwt/v5"
	"github.com/irdaislakhuafa/go-sdk/codes"
	"github.com/irdaislakhuafa/go-sdk/errors"
)

var (
	ErrClaimsNotPointer     = "claims must be a pointer"
	ErrInvalidSigningMethod = "signing method is not valid"
	ErrClaimsTypeNotEquals  = "claims type is not equals"
)

type JWTResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func NewJWT[C jwt.Claims](claims C, secret []byte) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return nil, errors.NewWithCode(codes.CodeJWTSignedStringError, err.Error())
	}

	return &tokenString, nil
}

func Validate[C jwt.Claims](tokenString string, secret []byte, claims C) (*jwt.Token, error) {
	if t := reflect.TypeOf(claims); t.Kind() != reflect.Ptr {
		return nil, errors.NewWithCode(codes.CodeJWTInvalidClaimsType, ErrClaimsNotPointer)
	}

	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if _, isOk := t.Method.(*jwt.SigningMethodHMAC); !isOk {
			return nil, errors.NewWithCode(codes.CodeJWTInvalidMethod, ErrInvalidSigningMethod)
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
		return claims, errors.NewWithCode(codes.CodeJWTInvalidClaimsType, ErrClaimsTypeNotEquals)
	}

	return claims, nil
}
