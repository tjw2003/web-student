package database

import (
	"database/sql"
	"web-student/server/models"
)

func InsertUser(user models.User) (int64, error) {
	return DB.Insert("users(username,password)", user.Username, user.Password)
}
func GetUserById(id int) (user models.User, err error) {
	err = DB.QueryRow("users", "id = $1", id).Scan(&user.ID, &user.Password, &user.Username)
	return user, err
}

func GetUserByUsername(username string) (user models.User, err error) {
	err = DB.QueryRow("users", "username = $1", username).Scan(&user.ID, &user.Password, &user.Username)
	return user, err
}

func IsUserExist(username string) (bool, error) {
	_, err := GetUserByUsername(username)
	if err == sql.ErrNoRows {
		return false, nil //不存在
	}
	if err != nil {
		return false, err
	}
	return true, nil //存在
}