/*
处理拦截器的请求
*/
package sharedauth

import (
	"context"
	"coolcar/shared/auth/token"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	authorizationHeader = "authorization"
	bearerPrefix        = "Bearer " //要加空格
)

type tokenVerifier interface {
	Verifier(token string) (string, error)
}

//结构前面加*，interface前面不加
type interceptor struct {
	verifier tokenVerifier
}

//拦截器 创建 grpc 用户 拦截器，得到publicKey
func Interceptor(publicKeyFile string) (grpc.UnaryServerInterceptor, error) {
	//从文件中读publicKey
	f, err := os.Open(publicKeyFile)
	if err != nil {
		return nil, fmt.Errorf("cannot open public key file:%v", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("cannot read public key:%v", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil {
		return nil, fmt.Errorf("cannot parse public key:%v", err)
	}
	i := &interceptor{
		verifier: &token.JWTTokenVerifier{
			PublicKey: pubKey,
		},
	}
	return i.HandlerReq, nil
}

//ctx是请求携带信息
//req是请求
//info是一些其他信息
//handler是下面要做的事情
func (i *interceptor) HandlerReq(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	token, err := tokenFromContext(ctx)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "请求头部信息有误")
	}

	aid, err := i.verifier.Verifier(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token不合法:%v", err)
	}
	return handler(ContextWithAccountID(ctx, AccountID(aid)), req)
}

func tokenFromContext(c context.Context) (string, error) {
	//拿请求参数
	m, ok := metadata.FromIncomingContext(c)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "获取不到请求参数")
	}
	tkn := ""
	for _, v := range m[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) {
			tkn = v[len(bearerPrefix):]
		}
	}

	if tkn == "" {
		return "", status.Error(codes.Unauthenticated, "无法从请求参数中解析出token")
	}

	return tkn, nil
}

type accountIDKey struct{}

//限定AccountID为string类型，
type AccountID string

func (a AccountID) String() string {
	return string(a)
}

func ContextWithAccountID(c context.Context, aid AccountID) context.Context {
	return context.WithValue(c, accountIDKey{}, aid)
}

func AccountIDFromContext(c context.Context) (AccountID, error) {
	v := c.Value(accountIDKey{})
	aid, ok := v.(AccountID)
	if !ok {
		return "", status.Error(codes.Unauthenticated, "AccountIDFromContext error")
	}
	return aid, nil
}
