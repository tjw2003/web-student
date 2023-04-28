package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"server/models"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib" //数据库驱动
)

type DBHelper struct {
	DB *sql.DB
}

// Default 使用默认的连接参数连接数据库（用户名为 postgres，地址及端口为 localhost:5432）并返回一个 DBHelper
func Default() *DBHelper {
	db, err := sql.Open("pgx", "postgres://postgres:@localhost:5432/postgres")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	return &DBHelper{db}
}

// CreateTable 使用 tableName 作为表名，tableDef 中的每一项作为 SQL 语句中括号中的每一项创建一个表
func (db *DBHelper) CreateTable(tableName string, tableDefs ...string) error {
	_, err := db.DB.Exec(
		fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (%v)", tableName, strings.Join(tableDefs, ", ")))
	return err
}

// Insert 使用 target 作为 INTO 与 VALUES 中间的部分，values 中的每一项作为值的每一项插入记录
func (db *DBHelper) Insert(target string, values ...any) (int64, error) {
	valuesPlaceHolders := []string{}
	for i := 1; i <= len(values); i++ {
		valuesPlaceHolders = append(valuesPlaceHolders, fmt.Sprintf("$%v", i))
	}
	res, err := db.DB.Exec(
		fmt.Sprintf("INSERT INTO %v VALUES (%v)", target, strings.Join(valuesPlaceHolders, ", ")), values...)
	if err != nil {
		return -1, err
	}
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return -1, err
	}
	return lastInsertId, err
}

func (db *DBHelper) InsertUser(user models.User) (int64, error) {
	id, err := db.Insert("users(username,password)", user.Username, user.Password)
	return id, err
}

func (db *DBHelper) InsertCourse(course models.Course) (int64, error) {
	id, err := db.Insert("lessons(name,credit,teacher)", course.Name, course.Credit, course.Teacher)
	return id, err
}

func (db *DBHelper) InsertSelectCourse(userId int, course_id int) (int64, error) {
	id, err := db.Insert("select_courses(user_id, course_id)", userId, course_id)
	return id, err
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

		course.ID = id
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
	rows, _ := db.DB.Query("SELECT id,coursename,credit,teacher FROM  selectcourse WHERE username = $1", username)
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

		course.ID = id
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
