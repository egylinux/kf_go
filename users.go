package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"github.com/urfave/cli/v2"
)

const (
	host	="localhost"
	port	= 5432
	user	= "ahmed"
	password =	"123456"
	dbname = "usersdb"
)

func main()  {
	//connectionstring
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
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
	//pass username,password from cli
	var uname,pass string
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Destination:&uname, Name:"u", Value:"root", Usage:  "Specify username. Default is root"},
		&cli.StringFlag{Destination:&pass, Name:"p", Value:"password", Usage:  "Specify password. Default is password"},
	}
	error := app.Run(os.Args)
	CheckError(error)

row,e:= db.Query(fmt.Sprintf(`Select * from userstb where username='%s' and password ='%s'`,uname,pass))
CheckError(e)

	if row.Next() {
		id,username,password,fullName:=0,"","",""
		err = row.Scan(&id,&fullName,&username,&password)
		CheckError(err)
		fmt.Println("Welcome ",fullName)
	}else{
		fmt.Println("Invalid login info")
	}




}
func CheckError(err error) {
	if err != nil {
		panic(err)
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