package services

import (
	"errors"
	"log"
	"web-student/internal/jwt"
	"web-student/server/models"
	"web-student/server/database"

	"github.com/gin-gonic/gin"
)

// Register
type RegisterService struct { //`form:"username"`进行绑定
	Username string `form:"username"`
	Password string `form:"password"`
}

func (s *RegisterService) Handle(c *gin.Context) (any, error) {
	// TODO: Register
	isUserExist, err := database.IsUserExist(s.Username)
	if err != nil {
		return nil, err
	}

	if isUserExist {
		log.Printf("[Register]: User %v already exists.\n", s.Username)
		return nil, errors.New("User already exists.")
	}
	_, err = database.InsertUser(models.User{
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
	isUserExist, err := database.IsUserExist(s.Username)
	if err != nil {
		return nil, err
	}

	if !isUserExist {
		log.Printf("[Login]: User %v not exists.\n", s.Username)
		return nil, errors.New("User not exist")
	}

	user, err := database.GetUserByUsername(s.Username)
	if err != nil {
		return nil, err
	}

	if user.Password != s.Password {
		log.Printf("[Login]:password is wrong\n")
		return nil, errors.New("password is not exist")
	}

	var jwtToken string
	jwtToken, err = jwt.CreateToken(user.ID, user.Username)

	res := make(map[string]any)
	res["token"] = jwtToken
	res["user"] = user

	return res, nil
}
