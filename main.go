package main

import (
	"fmt"
	"github.com/egylinux/kf_go/api"
	"github.com/egylinux/kf_go/db"
	"github.com/egylinux/kf_go/users"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "ahmed"
	password = "123456"
	dbname   = "usersdb"
)

func main() {
	// build services
	dbConnector, err := db.NewConnector("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))

	if err != nil {
		fmt.Println("DB Error: ", err.Error())
		return
	}

	usersMgr := users.NewManager(dbConnector)
	router := api.NewRouter(usersMgr)

	router.Logger.Fatal(router.Start(":1323"))
	//pass username,password from cli
	/*var uname, pass string
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Destination: &uname, Name: "u", Value: "root", Usage: "Specify username. Default is root"},
		&cli.StringFlag{Destination: &pass, Name: "p", Value: "password", Usage: "Specify password. Default is password"},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	// build services
	dbConnector,err := db.NewConnector("postgres",  fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))

	if err!=nil{
		fmt.Println("DB Error: ", err.Error())
		return
	}

	usersMgr := users.NewManager(dbConnector)

	//exist,err:=usersMgr.GetAll()
	exist,err:=usersMgr.IsExist(uname,pass)

	if err!=nil{
		fmt.Println("User not found", err.Error())
		return
	}

	if exist{
		fmt.Println("User is found")
	}*/
	/*result := CheckUser(uname, pass)
	fmt.Println(result)*/
}
