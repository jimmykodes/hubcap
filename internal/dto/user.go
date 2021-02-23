package dto

type User struct {
	ID        int64  `db:"id"`
	Email     string `db:"email"`
	ApiKey    string `db:"api_key"`
	SuperUser bool   `db:"super_user"`
}
