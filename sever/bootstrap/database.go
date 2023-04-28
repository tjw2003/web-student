package bootstrap

import (
	"server/database"
)

var DB *database.DBHelper

func init() {
	DB = database.OpenDatabase()
	DB.CreateUserTable()
	DB.CreateLessonTable()
	DB.CreateSelectCourseTable()
}
