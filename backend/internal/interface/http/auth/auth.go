package auth

//
//import (
//	"github.com/golang-jwt/jwt/v5"
//	"time"
//)
//
//type Auth interface {
//	Generate(uid int, secret string) (string, error)
//	Check(token string) bool
//}
//
//type Token struct{}
//
//func (t *Token) Generate(uid int, secret string) (string, error) {
//	token := jwt.NewWithClaims(
//		jwt.SigningMethodHS256,
//		jwt.MapClaims{
//			"user": uid,
//			"exp":  time.Now().Add(24 * time.Hour).Unix(),
//		},
//	)
//
//	return token.SignedString([]byte(secret))
//}
