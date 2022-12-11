package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const PublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAlyK+cR8u27fuMn60btT+
Xy1Rxd4yitBzcMPkD/Y6FonFDBFvEkXaPx+zcQy1jdaBllnuJ7Ff7xwNIby7FOFc
UZN4tDiU8lUsoZjS3cR/OEW+qPnVHrIYa+sGpVwP2VdBEbpb7SHEbvT9hHOTtEwU
Zkj35Unoj5Lwa4WFA8asEpmxDs2G3C87HnhRtdwRWUNIJ7YTAIOMt4VQ1GaQCqaL
niuJ/h6VWSqipqGMRFhXBWzIlNlcVIBXyjgvlALtFCTC2z+H1cDRRzAff4WhUefx
laKPprVOgHnlXhQl66X+antHnW7GQ/TFTzFdUUoUzwpYbikK+5Gz3VMXYt+4tFYt
BQIDAQAB
-----END PUBLIC KEY-----`

func TestVerify(t *testing.T) {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(PublicKey))
	if err != nil {
		t.Fatalf("cannot parse public key:%v", err)
	}

	v := &JWTTokenVerifier{
		PublicKey: pubKey,
	}

	cases := []struct {
		name    string    //测试的name
		tkn     string    //token
		nowTime time.Time //测试时模拟的现在时间
		want    string    //模拟得到的accountID
		wantErr bool      //模拟测试时出错的场景
	}{
		{
			name:    "valid_name",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjM4NDYwOTdlOGY5NWZmZTBkNjMxMzM1In0.AJtXQw7pjBx0JWmTPaPdU1LD0Q3lJIA9s2D0ATQnWz4_rPho5BFZJV1Ulkfapl5VvHuiKihFkNxAW0FgMTDjVFVhMnVThTH8BaPa4KTxghrB1P8w-0Hoi4UO6ISgMUXfxESFGkWMgMk0T4SZSf1GPUl3q4LFkPCn-HaEnuhe-I3lzCR4GC2mjOwODON74LP85yKihGvskD1rdzWFZrffeYJg7yVst0_XTbxiCxXgREUOaisDJz3FeaVsbwmalDckiguQfpt-L8Eu2jZT9b44QI8WvY-2vXe94eIePAwCpZ_GPMEKGzoU9XOBtgs7O_QryrWrYCrg0X1uOHLg2mLzSA",
			nowTime: time.Unix(1516239122, 0),
			want:    "63846097e8f95ffe0d631335",
			wantErr: false,
		},
		{
			name:    "token_expired",
			tkn:     "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjM4NDYwOTdlOGY5NWZmZTBkNjMxMzM1In0.AJtXQw7pjBx0JWmTPaPdU1LD0Q3lJIA9s2D0ATQnWz4_rPho5BFZJV1Ulkfapl5VvHuiKihFkNxAW0FgMTDjVFVhMnVThTH8BaPa4KTxghrB1P8w-0Hoi4UO6ISgMUXfxESFGkWMgMk0T4SZSf1GPUl3q4LFkPCn-HaEnuhe-I3lzCR4GC2mjOwODON74LP85yKihGvskD1rdzWFZrffeYJg7yVst0_XTbxiCxXgREUOaisDJz3FeaVsbwmalDckiguQfpt-L8Eu2jZT9b44QI8WvY-2vXe94eIePAwCpZ_GPMEKGzoU9XOBtgs7O_QryrWrYCrg0X1uOHLg2mLzSA",
			nowTime: time.Unix(1517239122, 0),
			wantErr: true,
		},
		{
			name:    "bad_token",
			tkn:     "bad_token",
			nowTime: time.Unix(1516239122, 0),
			wantErr: true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			//控制JWT时间，让其匹配测试用例
			jwt.TimeFunc = func() time.Time {
				return c.nowTime
			}
			accountID, err := v.Verifier(c.tkn)

			if !c.wantErr && err != nil {
				t.Errorf("verification failed:%v", err)
			}

			if c.wantErr && err == nil {
				t.Errorf("want error;got no error")
			}

			if accountID != c.want {
				t.Errorf("wrong account id, want:%q,got:%q", c.want, accountID)
			}
		})
	}
}
