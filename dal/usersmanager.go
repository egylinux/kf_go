package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/urfave/cli/v2"
)

func init() {
	//connectionstring
	psqlconn := ConnectionString()
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()
	// check db
	err = db.Ping()
	CheckError(err)
	//Check if user table exist , false = create table
	CheckUserTable(psqlconn)
}

func CheckUser(userName, pasword string) string {

	db, err := sql.Open("postgres", ConnectionString())
	row, e := db.Query(fmt.Sprintf(`Select * from userstb where username='%s' and password ='%s'`, userName, pasword))
	CheckError(e)

	if row.Next() {
		id, username, password, fullName := 0, "", "", ""
		err = row.Scan(&id, &fullName, &username, &password)
		CheckError(err)
		return fmt.Sprintf("Welcome %s", fullName)
	} else {
		return "Invalid login info"
	}
}

func CheckUserTable(psqlconn string) {

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	strsql := `SELECT count(*) as counter FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND  schemaname != 'information_schema' And tablename = 'userstb';`
	rows, e := db.Query(strsql)
	CheckError(e)
	counter := 0

	for rows.Next() {
		err = rows.Scan(&counter)
		CheckError(err)

	}

	if counter == 0 {
		strsql = `Create table userstb(id SERIAL PRIMARY KEY,fullName varchar(255),username Varchar(255),password varchar(50))`
		_, e = db.Exec(strsql)
		CheckError(err)
		fmt.Printf("Users table created\n")
	}
	defer db.Close()

}
