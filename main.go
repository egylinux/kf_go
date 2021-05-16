package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	Connect()
	//pass username,password from cli
	var uname, pass string
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		&cli.StringFlag{Destination: &uname, Name: "u", Value: "root", Usage: "Specify username. Default is root"},
		&cli.StringFlag{Destination: &pass, Name: "p", Value: "password", Usage: "Specify password. Default is password"},
	}
	error := app.Run(os.Args)
	CheckError(error)

	result := CheckUser(uname, pass)
	fmt.Println(result)
}
