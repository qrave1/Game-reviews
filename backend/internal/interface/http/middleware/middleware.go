package middleware

//
//import (
//	"fmt"
//	"gameReview/internal/interface/http/auth"
//	"github.com/gofiber/fiber/v2"
//	"github.com/golang-jwt/jwt/v5"
//	"strconv"
//	"strings"
//)
//
//type AuthMiddleware struct {
//	tokenParser *jwt.Parser
//	tokenSecret string
//	auth        auth.Auth
//}
//
//func NewAuthMiddlewareImpl(tokenSecret string) *AuthMiddleware {
//	return &AuthMiddleware{
//		tokenParser: jwt.NewParser(
//			jwt.WithValidMethods(
//				[]string{"HS256"},
//			),
//		),
//		tokenSecret: tokenSecret,
//	}
//}
//
//func (t *AuthMiddleware) Handle(c *fiber.Ctx) error {
//
//	authHeader := ctx.Header("Authorization")
//	if authHeader == "" {
//		return fmt.Errorf("forbidden")
//	}
//
//	tokenParts := strings.Split(authHeader, " ")
//	if len(tokenParts) < 2 {
//		return fmt.Errorf("invalid token")
//	}
//
//	tokenType, tokenString := tokenParts[0], tokenParts[1]
//	if tokenType != "Bearer" {
//		return fmt.Errorf("unsupported token type")
//	}
//
//	var tokenDto dto.TransportAPIAuthToken
//	_, err := t.tokenParser.ParseWithClaims(tokenString, &tokenDto, func(token *jwt.Token) (interface{}, error) {
//		return []byte(t.serviceSecret), nil
//	})
//	if err != nil {
//		return fmt.Errorf("error validate token. %w", err)
//	}
//
//	acc, err := t.auth.Generate(strconv.Itoa(tokenDto.AccountId), auth.Metadata(map[string]string{
//		"user_id":  strconv.Itoa(tokenDto.UserId),
//		"is_admin": boolStr(tokenDto.IsAdmin),
//	}))
//	if err != nil {
//		return fmt.Errorf("error generate auth params. %w", err)
//	}
//
//	ctx.SetUserContext(auth.ContextWithAccount(ctx.UserContext(), acc))
//
//	return nil
//}
//
//func boolStr(b bool) (v string) {
//	v = "false"
//	if b {
//		v = "true"
//	}
//	return
//}
