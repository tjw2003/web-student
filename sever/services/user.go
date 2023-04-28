package services

import (
	"errors"
	"log"
	"server/bootstrap"
	"server/jwt"
	"server/models"

	"github.com/gin-gonic/gin"
)

// Register
type RegisterService struct { //`form:"username"`进行绑定
	Username string `form:"username"`
	Password string `form:"password"`
}

func (s *RegisterService) Handle(c *gin.Context) (any, error) {
	// TODO: Register
	isUserExist, err := bootstrap.DB.IsUserExist(s.Username)
	if err != nil {
		return nil, err
	}

	if isUserExist {
		log.Printf("[Register]: User %v already exists.\n", s.Username)
		return nil, errors.New("User already exists.")
	}
	err = bootstrap.DB.InsertUser(models.User{
		Username: s.Username,
		Password: s.Password,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}


// Login
type Login struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func (s *Login) Handle(c *gin.Context) (any, error) {
	log.Printf("[login/Handle]: %v %v\n", s.Username, s.Password)
	//user := &User{}
	isUserExist, err := bootstrap.DB.IsUserExist(s.Username)
	if err != nil {
		return nil, err
	}

	if !isUserExist {
		log.Printf("[Login]: User %v not exists.\n", s.Username)
		return nil, errors.New("User not exist")
	}

	password, _ := bootstrap.DB.QueryPassword(s.Username)
	id, _ := bootstrap.DB.QueryID(s.Username)

	if password != s.Password {
		log.Printf("[Login]:password is wrong\n")
		return nil, errors.New("password is not exist")
	}

	user := &models.User{
		ID:       id,
		Username: s.Username,
		Password: s.Password,
	}
	var jwtToken string
	jwtToken, err = jwt.CreateToken(user.ID, user.Username)

	res := make(map[string]any)
	res["token"] = jwtToken
	res["user"] = user

	return res, nil
}
