package main

import (
	"fmt"
	"github.com/egylinux/kf_go/db"
	"github.com/egylinux/kf_go/users"
	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	//pass username,password from cli
	var uname, pass string
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
	dbConnector,err := db.NewConnector("", "")
	if err!=nil{
		fmt.Println("DB Error: ", err.Error())
		return
	}

	usersMgr := users.NewManager(dbConnector)

	exist,err:=usersMgr.IsExist("","")
	if err!=nil{
		fmt.Println("User not found", err.Error())
		return
	}

	if exist{
		fmt.Println("User is found")
	}
	/*result := CheckUser(uname, pass)
	fmt.Println(result)*/
}
