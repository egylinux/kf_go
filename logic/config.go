package logic

import "fmt"

const (
	host     = "localhost"
	port     = 5432
	user     = "ahmed"
	password = "123456"
	dbname   = "usersdb"
)

func ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
