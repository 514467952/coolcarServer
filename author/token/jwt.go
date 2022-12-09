package token

import (
	"crypto/rsa"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTTokenGen struct {
	issuer string
	//时间由外界传入，方便单元测试
	nowFunc    func() time.Time
	privateKey *rsa.PrivateKey
}

//构造函数
func NewJWTTokenGen(issuer string, privateKey *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{
		issuer:     issuer,
		nowFunc:    time.Now,
		privateKey: privateKey,
	}
}

func (t *JWTTokenGen) GenerateToken(accountID string, expire time.Duration) (string, error) {
	nowSec := t.nowFunc().Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Issuer:    t.issuer,
		IssuedAt:  nowSec,
		ExpiresAt: nowSec + int64(expire.Seconds()),
		Subject:   accountID,
	})

	return tkn.SignedString(t.privateKey)
}
