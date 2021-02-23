package dto

type Vehicle struct {
	ID     int64  `db:"id"`
	Name   string `db:"name"`
	Make   string `db:"make"`
	Model  string `db:"model"`
	Year   int    `db:"year"`
	UserID int64  `db:"user_id"`
}
