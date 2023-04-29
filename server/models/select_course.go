package models

type SelectCourse struct {
	ID      int    `json:"id"`
	UserId    string `json:"user_id"`
	CourseId  string `json:"course_id"`
}