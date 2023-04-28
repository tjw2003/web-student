package database

import (
"server/models"
	"log"
	"server/internal/database"
)

// DB is the DBHelper
var DB *database.DBHelper

func InitializeTables() {
	var err error
	err = DB.CreateTable("users",
		"id SERIAL",
		"username TEXT NOT NULL UNIQUE",
		"password TEXT NOT NULL",
	)
	if err != nil {
		log.Panicln("Failed to create table users")
	}
	err = DB.CreateTable("lessons",
		"id SERIAL",
		"name TEXT NOT NULL",
		"credit TEXT NOT NULL",
		"teacher TEXT NOT NULL",
	)
	if err != nil {
		log.Panicln("Failed to create table lessons")
	}
	err = DB.CreateTable("select_courses",
		"id SERIAL",
		"user_id REFERENCES users",
		"course_id REFERENCES courses",
	)
	if err != nil {
		log.Panicln("Failed to create table select_courses")
	}
}

func init() {
	DB = database.Default()
	InitializeTables()
}

//--- 数据库操作 ---

func InsertUser(user models.User) (int64, error) {
	return DB.Insert("users(username,password)", user.Username, user.Password)
}

func InsertCourse(course models.Course) (int64, error) {
	return DB.Insert("lessons(name,credit,teacher)", course.Name, course.Credit, course.Teacher)
}

func InsertSelectCourse(userId int, course_id int) (int64, error) {
	return DB.Insert("select_courses(user_id, course_id)", userId, course_id)
}