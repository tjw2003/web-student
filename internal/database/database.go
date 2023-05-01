package database

import (
	"database/sql"
	// "errors"
	"fmt"
	"log"
	// "web-student/server/models"
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

// Query query data from table with tableName.
//
// When no conditions provided, it will not add the WHERE clause to the SQL.
// When conditions has the length of 1, it will be treated as the content of the WHERE clause
// When conditions has the length greater than 1, the first one will be used as the content of the WHERE clause, and others will be treated as the arguments. 
func (db *DBHelper) Query(tableName string, conditions ...interface{}) (rows *sql.Rows, err error) {
	if len(conditions) == 0 {
		rows, err = db.DB.Query(
			fmt.Sprintf("SELECT * FROM %v", tableName),
		)
	} else if len(conditions) == 1 {
		rows, err = db.DB.Query(
			fmt.Sprintf("SELECT * FROM %v WHERE %v", tableName, conditions[0]),
		)
	} else {
		rows, err = db.DB.Query(
			fmt.Sprintf("SELECT * FROM %v WHERE %v", tableName, conditions[0]), conditions[1:]...,
		)
	}
	return rows, err
}

func (db *DBHelper) QueryRow(tableName string, conditions ...interface{}) (row *sql.Row) {
	if len(conditions) == 0 {
		row = db.DB.QueryRow(
			fmt.Sprintf("SELECT * FROM %v", tableName),
		)
	} else if len(conditions) == 1 {
		row = db.DB.QueryRow(
			fmt.Sprintf("SELECT * FROM %v WHERE %v", tableName, conditions[0]),
		)
	} else {
		row = db.DB.QueryRow(
			fmt.Sprintf("SELECT * FROM %v WHERE %v", tableName, conditions[0]), conditions[1:]...,
		)
	}
	return row
}
