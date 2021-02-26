package dto

type Vehicle struct {
	ID     int64  `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Make   string `db:"make" json:"make"`
	Model  string `db:"model" json:"model"`
	Year   int    `db:"year" json:"year"`
	UserID int64  `db:"user_id" json:"user_id"`
}
