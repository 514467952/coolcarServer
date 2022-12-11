package token

import (
	"crypto/rsa"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type JWTTokenVerifier struct {
	PublicKey *rsa.PublicKey
}

func (v *JWTTokenVerifier) Verifier(token string) (string, error) {
	t, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(*jwt.Token) (interface{}, error) {
		return v.PublicKey, nil
	})
	//验证是否能够通过公钥解出token
	if err != nil {
		return "", fmt.Errorf("cannot parse%v", err)
	}
	//验证token语法是否合法
	if !t.Valid {
		return "", fmt.Errorf("token not valid")
	}

	//检验token的声明类型是否是StandardClaims
	clm, ok := t.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", fmt.Errorf("token claim is not vStandardClaims")
	}
	//验证token里面携带信息是否合法，例如是否已经到过期时间
	if err := clm.Valid(); err != nil {
		return "", fmt.Errorf("claim not valid:%v", err)
	}

	return clm.Subject, nil
}
