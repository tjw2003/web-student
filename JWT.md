# JWT

创建jwt

```go
type MyCustomClaims struct {
    Userid    int
	jwt.RegisteredClaims
}

func CreateToken(userid int) (string, error) {
	return jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), MyCustomClaims{
        userid,
        jwt.RegisteredClaims{
		},
	}).SignedString(signedString)
}


//token := jwt.New(jwt.GetSigningMethod("HS256"));//使用HS256方法的token
//jwt.RegisteredClaims{
//		},jwt中间部分
//token.SignedString()转换成字符串
```

解析jwt

```go
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

//jwt.parse将字符串转换成token结构体
```

```go
func GetTokenStr(c *gin.Context) string {
	tokenStr := ""
	tokenStr = strings.ReplaceAll(c.Request.Header.Get("Authorization"), "Bearer ", "")
	return tokenStr
}


func MustGetClaims(c *gin.Context) (*MyCustomClaims, error) {
	log.Println("[MustGetClaims]")
	tokenStr := GetTokenStr(c)
	log.Printf("[MustGetClaims]: TokenStr: %v\n", tokenStr)

	token, err := DecodeTokenStr(tokenStr)
    if err != nil {
		return nil, err
    }
	log.Printf("[MustGetClaims]: Token: %v\n", token)

	return token.Claims.(*MyCustomClaims), nil
}


```

