package models

type Course struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Credit  string `json:"credit"`
	Teacher string `json:"teacher" `
}