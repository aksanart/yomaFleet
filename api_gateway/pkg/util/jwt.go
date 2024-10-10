package util

import (
	"context"
	"fmt"
	"time"

	"github.com/aksan/weplus/apigw/pkg/config"
	"github.com/aksan/weplus/apigw/pkg/constant"
	uerror "github.com/aksan/weplus/apigw/pkg/error"
	"github.com/golang-jwt/jwt"
)

type MyClaims struct {
	Email string `json:"email"`
	Sub   string `json:"sub"`
	Name  string `json:"name"`
	Iat   int64  `json:"iat"`
	Exp   int64  `json:"exp"`
}

func (c MyClaims) Valid() error {
	if time.Unix(c.Exp, 0).Before(time.Now()) {
		return fmt.Errorf("token is expired")
	}
	return nil
}

func GenerateJwt(email string) (string, error) {
	// Define the signing key
	var signingKey = []byte(config.GetConfig("jwt_secret").GetString())
	claims := MyClaims{
		Email: email,
		Sub:   "auth",
		Iat:   time.Now().Unix(),
		Exp:   time.Now().Add(constant.JWT_EXPIRE).Unix(),
	}
	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign the token with the secret
	return token.SignedString(signingKey)
}

func ValidateToken(tokenString string) (*MyClaims, error) {
	var signingKey = []byte(config.GetConfig("jwt_secret").GetString())
	// Extract the token
	extractedClaims := &MyClaims{}
	parsedToken, err := jwt.ParseWithClaims(tokenString, extractedClaims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !parsedToken.Valid {
		return nil, uerror.ErrorUnauthorized.BuildError(context.Background())
	}
	return extractedClaims, nil
}
