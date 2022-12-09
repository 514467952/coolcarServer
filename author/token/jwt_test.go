package token

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAlyK+cR8u27fuMn60btT+Xy1Rxd4yitBzcMPkD/Y6FonFDBFv
EkXaPx+zcQy1jdaBllnuJ7Ff7xwNIby7FOFcUZN4tDiU8lUsoZjS3cR/OEW+qPnV
HrIYa+sGpVwP2VdBEbpb7SHEbvT9hHOTtEwUZkj35Unoj5Lwa4WFA8asEpmxDs2G
3C87HnhRtdwRWUNIJ7YTAIOMt4VQ1GaQCqaLniuJ/h6VWSqipqGMRFhXBWzIlNlc
VIBXyjgvlALtFCTC2z+H1cDRRzAff4WhUefxlaKPprVOgHnlXhQl66X+antHnW7G
Q/TFTzFdUUoUzwpYbikK+5Gz3VMXYt+4tFYtBQIDAQABAoIBACj6yLyPgfsdkj6q
0G3E3vwbo7KNHxM4ljSbSPlzACkgTgeJqp4/xn+bLuKWWZwN77E+zS6CD8sbxgvB
ytBMkuWjoPTs6qYSpjli4Lo1D3HIp3vV0g2yqKCKZ5Dqt3ltCT2vVEGmCEc1so/5
brFSd62MV3saYS6c3weoenqdogagrq2GYdeANSnhCVn3P/mgkUmlHVqviliCRiDs
Byd7GgmHO+xv+bM6yJ5DMgYNZWvkvvWEfzrWDUvl8iPJJOZrQDVCUxdJ00dx1Xn+
M8q1R8i1TSsnyRIML6y5/03Rzhf8DARrRGYpX4DFPl1X55Q2g/kghHISQHIgN6HL
cGrbBfkCgYEA1iU4k4xo4r8JiHWf/6+W5QWsXm2RqZH+9ATHt7qKb8gx4Rfnb2Ae
nXU8qGH6BWxIWwO2oUwWek9WgZZxGTx8CFkadKQB7ixjM/AW+a7bp2lBAD4AYmVp
OvmzcjdnaHpLfdYyRRZxVXLKkRadIjsbEZmwCOhn60jpT5apYuwQ3p8CgYEAtKzU
YEhXJdWPwkOKGClrsJQDyM2m7kvdkbGYco3LPRr2jk9apLsfJ2+hy8AIO/hdlHXM
4xeM7Z2j9maD9XG3yduvzXkwDgvF58nPq4N0NbARYGFYhFqbi93MfU77IvqdZC2l
iye5sigcPkLTXsvby/MJKeoJnwbKKYBGjbQNZdsCgYAsQJ7TZtWZ3c5zg5MfvEJb
a3O2Q8AxIXllJcO7xPGsrdsY3960llFdSofRaAXJrxm4rSjFZjrS5ahDuTn/9A7R
jtA3wFihxkxtxDSLPkYn70k+apGbw8ceJ9GcTbSx02vcQjI7MqsS9FjF6L1qAHrU
hUTqYlRvO67zAnhYd934kwKBgQCoc4cjKe9O0BnTxsLHGj9Uh/wrSUp9XPB8+Tco
fvlhxHTiZCwf8HLwgs04OZezyjH9zAM2K+vyUwrfHd+khN0VcCUuDvJ19hYTsP5V
bB96OulIkhpEdHIX6rjQkEXn6/+4ujJhVq105J7Ikeeet7T7J2KA2LsVsO/l17pw
PdJMBwKBgE/BETGwpkKPte331HeOvaSKcuV6yw7y6f/LYi0qaMP5FHvxuwTblZP5
WSnu/WIRd0SqmEYtFULCc2uuFuXFJfxcxD1I7Pxfax57n3UfPxak2EbuvVYHcK61
ziJ6UF9CgUshBwSm7RtB8djA38XkkZR32GDf+H+SammgkEqwbPwc
-----END RSA PRIVATE KEY-----`

func TestGenerateToken(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot parse private key:%v", err)
	}

	g := NewJWTTokenGen("coolcar/auth", key)
	g.nowFunc = func() time.Time {
		return time.Unix(1516239022, 0)
	}
	tkn, err := g.GenerateToken("63846097e8f95ffe0d631335", 2*time.Hour)
	if err != nil {
		t.Errorf("cannot  generate token:%v", err)
	}

	want := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDYyMjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoiY29vbGNhci9hdXRoIiwic3ViIjoiNjM4NDYwOTdlOGY5NWZmZTBkNjMxMzM1In0.AJtXQw7pjBx0JWmTPaPdU1LD0Q3lJIA9s2D0ATQnWz4_rPho5BFZJV1Ulkfapl5VvHuiKihFkNxAW0FgMTDjVFVhMnVThTH8BaPa4KTxghrB1P8w-0Hoi4UO6ISgMUXfxESFGkWMgMk0T4SZSf1GPUl3q4LFkPCn-HaEnuhe-I3lzCR4GC2mjOwODON74LP85yKihGvskD1rdzWFZrffeYJg7yVst0_XTbxiCxXgREUOaisDJz3FeaVsbwmalDckiguQfpt-L8Eu2jZT9b44QI8WvY-2vXe94eIePAwCpZ_GPMEKGzoU9XOBtgs7O_QryrWrYCrg0X1uOHLg2mLzSA"
	if tkn != want {
		t.Errorf("wrong token generated, want:%q:got:%q", want, tkn)
	}
}
