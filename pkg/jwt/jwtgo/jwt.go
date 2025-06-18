package jwtgo

import (
	"context"
	"encoding/hex"
	"fmt"
	goJwtPkg "github.com/golang-jwt/jwt/v5"
	"github.com/sreekar2307/katha/pkg/jwt"
)

type gojwt struct {
	secret []byte
}

func NewGoJWT(secretAsHex string) (jwt.JWT, error) {
	secret, err := hex.DecodeString(secretAsHex)
	if err != nil {
		return nil, fmt.Errorf("failed to decode secret from hex: %w", err)
	}
	return gojwt{secret: secret}, nil
}

func (g gojwt) Token(_ context.Context, m map[string]any) (string, error) {
	token := goJwtPkg.NewWithClaims(goJwtPkg.SigningMethodHS256, goJwtPkg.MapClaims(m))
	signedToken, err := token.SignedString(g.secret)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return signedToken, nil
}

func (g gojwt) Validate(_ context.Context, signedToken string) (map[string]any, error) {
	token, err := goJwtPkg.Parse(signedToken, func(token *goJwtPkg.Token) (interface{}, error) {
		return g.secret, nil
	}, goJwtPkg.WithValidMethods([]string{goJwtPkg.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return token.Claims.(goJwtPkg.MapClaims), nil
}
