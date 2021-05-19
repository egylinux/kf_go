package logic

import (
	"database/sql"
	_ "github.com/lib/pq"
)
func Connect()  {
	psqlconn := ConnectionString()
	// open database
	_, err := sql.Open("postgres", psqlconn)
	CheckError(err)
}