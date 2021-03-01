package dto

type User struct {
	ID        int64  `db:"id" json:"id"`
	Username  string `db:"username" json:"username"`
	ApiKey    string `db:"api_key" json:"api_key"`
	SuperUser bool   `db:"super_user" json:"super_user"`
}
