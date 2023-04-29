package jwt

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


var signedString = []byte("prj")
type MyCustomClaims struct {
    UserId    int
	Username string
	jwt.RegisteredClaims
}

func CreateToken(userId int, username string) (string, error) {
	return jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), MyCustomClaims{
        userId,
		username,
        jwt.RegisteredClaims{
		},
	}).SignedString(signedString)
}

func GetTokenStr(c *gin.Context) string {
	tokenStr := ""
	tokenStr = strings.ReplaceAll(c.Request.Header.Get("Authorization"), "Bearer ", "")
	return tokenStr
}

func DecodeTokenStr(tokenStr string) (*jwt.Token, error) {
	var token *jwt.Token
	var err error
    token, err = jwt.ParseWithClaims(tokenStr, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
        return signedString, nil
    }, jwt.WithTimeFunc(func() time.Time {return time.Unix(0, 0)}))
	if err != nil {
		return token, err
	}
	return token, nil
}

