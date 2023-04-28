package middlewars

import (
	"log"
	"net/http"
	"server/jwt"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwt.MustGetClaims(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, "token invalid")
			c.Abort()
			return
		}
		log.Printf("[middlewares/JWTAuth]: claim: %v\n", claims)
		c.Set("id", claims.Userid)
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
