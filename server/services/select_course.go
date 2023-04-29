package services

import (
	"log"
	"web-student/server/database"

	"github.com/gin-gonic/gin"
)


type SelectCourseService struct{
	CourseId int `form:"course_id"`
}

func (s *SelectCourseService) Handle(c *gin.Context) (any, error) {
	id, err := database.InsertSelectCourse(c.GetInt("id"), s.CourseId)
	return id, err
}

type GetSelectedCoursesService struct{
	// Username string `form:"Username"`
}

func (s *GetSelectedCoursesService) Handle(c *gin.Context) (any, error) {
	courses, err := database.GetSelectedCourses(c.GetInt("id"))
	if err != nil {
		log.Printf("[GetSelectedCoursesService]: Error %v\n", err)
		return nil, err
	}

	return courses, nil
}

