package database

import (
	"database/sql"
	"errors"
	"log"
	"server/models"

	_ "github.com/jackc/pgx/v5/stdlib" //数据库驱动
)

type DBHelper struct {
	DB *sql.DB
}

func OpenDatabase() *DBHelper {
	db, err := sql.Open("pgx", "postgres://postgres:@localhost:5432/postgres")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return &DBHelper{db}
}

func (db *DBHelper) CreateUserTable() {
	result, err := db.DB.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL,username VARCHAR(20) NOT NULL UNIQUE,password VARCHAR(20) NOT NULL)")
	if err != nil {
		log.Println(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	if rows != 1 {
		log.Printf("expected to affect 1 row, affected %d\n", rows)
	}
}

func (db *DBHelper) CreateLessonTable() {
	result, err := db.DB.Exec("CREATE TABLE IF NOT EXISTS lesson (id SERIAL,coursename VARCHAR(20) NOT NULL,credit VARCHAR(20) NOT NULL,teacher VARCHAR(20) NOT NULL)")
	if err != nil {
		log.Println(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	if rows != 1 {
		log.Printf("expected to affect 1 row, affected %d\n", rows)
	}
}

func (db *DBHelper) CreateSelectCourseTable() {
	result, err := db.DB.Exec("CREATE TABLE IF NOT EXISTS selectcourse (id SERIAL,username Varchar(20),coursename VARCHAR(20) NOT NULL,credit VARCHAR(20) NOT NULL,teacher VARCHAR(20) NOT NULL)")
	if err != nil {
		log.Println(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	if rows != 1 {
		log.Printf("expected to affect 1 row, affected %d\n", rows)
	}
}

func (db *DBHelper) InsertUser(user models.User) error {
	_, err := db.DB.Exec("INSERT INTO users(username,password) VALUES ($1,$2)", user.Username, user.Password)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (db *DBHelper) InsertCourse(course models.Course) error {
	res, err := db.DB.Exec("INSERT INTO lesson(coursename,credit,teacher) VALUES ($1,$2,$3)", course.Name, course.Credit, course.Teacher)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(res.LastInsertId())
	return nil
}

func (db *DBHelper) InsertSelectCourse(username string,course models.Course) error {
	res, err := db.DB.Exec("INSERT INTO selectcourse(username,coursename,credit,teacher) VALUES ($1,$2,$3,$4)",username, course.Name, course.Credit, course.Teacher)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(res.LastInsertId())
	return nil
}

func (db *DBHelper) QueryPassword(username string) (string, error) {
	var password string
	err := db.DB.QueryRow("SELECT password FROM users WHERE username = $1", username).Scan(&password)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with username %s\n", username)
		return "", errors.New("no user with username")
	case err != nil:
		log.Printf("query error: %v\n", err)
		return "", errors.New("query error")
	default:
		log.Printf("password is %s\n", password)
		return password, nil
	}
}

func (db *DBHelper) QueryID(username string) (int, error) {
	var id int
	err := db.DB.QueryRow("SELECT id FROM users WHERE username = $1", username).Scan(&id)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with username %s\n", username)
		return 0, errors.New("no user with username")
	case err != nil:
		log.Printf("query error: %v\n", err)
		return 0, errors.New("query error")
	default:
		log.Printf("password is %d\n", id)
		return id, nil
	}
}

func (db *DBHelper) QueryAllCourse() ([]models.Course, error) {
	rows, _ := db.DB.Query("SELECT id,coursename,credit,teacher FROM lesson")
	courses := make([]models.Course, 0)

	for rows.Next() {
		var id int
		var coursename string
		var credit string
		var teacher string

		var course models.Course

		if err := rows.Scan(&id, &coursename, &credit, &teacher); err != nil {
			return nil, nil
		}

		course.ID = id;
		course.Name = coursename
		course.Credit = credit
		course.Teacher = teacher

		courses = append(courses, course)
	}

	err := rows.Close()
	if err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}
// WHERE username = $1",username

func (db *DBHelper) QuerySelectCourse(username string) ([]models.Course, error) {
	rows, _ := db.DB.Query("SELECT id,coursename,credit,teacher FROM  selectcourse WHERE username = $1",username)
	courses := make([]models.Course, 0)

	for rows.Next() {
		var id int
		var coursename string
		var credit string
		var teacher string

		var course models.Course

		if err := rows.Scan(&id, &coursename, &credit, &teacher); err != nil {
			return nil, nil
		}

		course.ID = id;
		course.Name = coursename
		course.Credit = credit
		course.Teacher = teacher

		courses = append(courses, course)
	}

	err := rows.Close()
	if err != nil {
		return nil, err
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}

func (db *DBHelper) IsUserExist(username string) (bool, error) {
	var user string
	err := db.DB.QueryRow("SELECT username FROM users WHERE username=$1", username).Scan(&user)
	if err == sql.ErrNoRows {
		return false, nil //不存在
	}
	if err != nil {
		return false, err
	}
	return true, nil //存在
}
