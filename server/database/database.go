package database

import (
	"log"
	"web-student/internal/database"
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
	err = DB.CreateTable("courses",
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
			     "user_id INT REFERENCES users(id)",
			     "course_id INT REFERENCES courses(id)",
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
