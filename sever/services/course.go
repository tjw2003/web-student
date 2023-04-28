package services

import (
	"log"
	"server/bootstrap"
	"server/models"

	"github.com/gin-gonic/gin"
)

// Create Course
type CreateCourseService struct {
	Coursename string `form:"coursename"`
	Credit     string `form:"credit"`
	Teacher    string `form:"teacher"`
}

func (s *CreateCourseService) Handle(c *gin.Context) (any, error) {
	log.Println("Course Handle")

	err := bootstrap.DB.InsertCourse(models.Course{
		Name:    s.Coursename,
		Credit:  s.Credit,
		Teacher: s.Teacher,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}


type GetCoursesService struct{}

func (s *GetCoursesService) Handle(c *gin.Context) (any, error) {
	courses, err := bootstrap.DB.QueryAllCourse()
	if err != nil {
		log.Printf("[GetCoursesService]: Error %v\n", err)
		return nil, err
	}

	return courses, nil
}
