package services

import (
	"log"
	"server/database"

	"github.com/gin-gonic/gin"
)


type SelectCoursesService struct{
	CourseId int `form:"course_id"`
}

func (s *SelectCoursesService) Handle(c *gin.Context) (any, error) {
	id, err := database.InsertSelectCourse(c.GetInt("id"), s.CourseId)
	return id, err
}

type GetSelectCoursesService struct{
	// Username string `form:"Username"`
}

func (s *GetSelectCoursesService) Handle(c *gin.Context) (any, error) {
	courses, err := database.DB.QuerySelectCourse(c.GetString("username"))
	if err != nil {
		log.Printf("[GetSelectCoursesService]: Error %v\n", err)
		return nil, err
	}

	return courses, nil
}

