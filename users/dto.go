package users

type User struct {
	Id       int
	Username string `db:"username"`
	Password string `db:"password"`
	Fullname string `db:"fullname"`
}
