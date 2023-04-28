package models

type Course struct {
	ID      int    `sqlType:"serial" json:"id"`
	Name    string `sql:"NOT NULL CHECK (name <> '')" json:"name"`
	Credit  string `sql:"NOT NULL CHECK (description <> '')" json:"credit"`
	Teacher string `sql:"NOT NULL CHECK (Teacher <> '')" json:"teacher" `
}