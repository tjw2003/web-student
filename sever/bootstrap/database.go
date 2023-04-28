package bootstrap

import (
	"log"
	"server/database"
)

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
	DB = database.OpenDatabase()
	InitializeTables()
}
