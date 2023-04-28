package models

type User struct {
	ID       int    `sqlType:"serial" sql:"PRIMARY KEY" json:"id"`
	Username string `sql:"NOT NULL CHECK (username <> '')" json:"username"`
	Password string `sql:"NOT NULL" json:"-"`
}
