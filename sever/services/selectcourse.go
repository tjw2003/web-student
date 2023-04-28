package services

import (
	"log"
	"server/bootstrap"
	"server/models"

	"github.com/gin-gonic/gin"
)


type SelectCoursesService struct{
	// Username string `form:"Username"`
	Coursename string `form:"coursename"`
	Credit     string `form:"credit"`
	Teacher    string `form:"teacher"`
}

func (s *SelectCoursesService) Handle(c *gin.Context) (any, error) {
	err := bootstrap.DB.InsertSelectCourse(c.GetString("username"),
		models.Course{
			Name:    s.Coursename,
			Credit:  s.Credit,
			Teacher: s.Teacher,
		},
	)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

type GetSelectCoursesService struct{
	// Username string `form:"Username"`
}

func (s *GetSelectCoursesService) Handle(c *gin.Context) (any, error) {
	courses, err := bootstrap.DB.QuerySelectCourse(c.GetString("username"))
	if err != nil {
		log.Printf("[GetSelectCoursesService]: Error %v\n", err)
		return nil, err
	}

	return courses, nil
}

