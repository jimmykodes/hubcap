package dto

type User struct {
	ID        int64  `db:"id" json:"id"`
	Email     string `db:"email" json:"email"`
	ApiKey    string `db:"api_key" json:"api_key"`
	SuperUser bool   `db:"super_user" json:"super_user"`
}
