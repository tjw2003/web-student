package middlewars

import (
	"log"
	"net/http"
	"web-student/internal/serializer"
	"web-student/internal/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		
	log.Println("[middlewares/JWTAuth]")
		tokenStr := jwt.GetTokenStr(c)
		log.Printf("[middlewares/JWTAuth]: TokenStr: %v\n", tokenStr)

		token, err := jwt.DecodeTokenStr(tokenStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
			return
		}
		log.Printf("[middlewares/JWTAuth]: Token: %v\n", token)

		claims := token.Claims.(*jwt.MyCustomClaims)
		if err != nil {
			c.JSON(http.StatusBadRequest, serializer.ErrorResponse(err))
			return
		}
		log.Printf("[middlewares/JWTAuth]: claim: %v\n", claims)
		c.Set("id", claims.UserId)
		c.Set("username", claims.Username)
		c.Next()
	}
}

func Frontend(ctx *gin.Context) { //中间件
	path := ctx.Request.URL.Path

	if strings.HasPrefix(path, "/api") {
		ctx.Next()
	} else {
		log.Println("static")
		
		if strings.LastIndex(path, ".") == -1 {
			ctx.Request.URL.Path += ".html"
		}
		http.FileServer(http.Dir("./frontend/build")).ServeHTTP(ctx.Writer, ctx.Request)
		ctx.Abort()
	}
}
