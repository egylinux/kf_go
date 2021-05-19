package main

import (
	"fmt"
	"github.com/egylinux/kf_go/db"
	"github.com/egylinux/kf_go/users"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"net/http"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "ahmed"
	password = "123456"
	dbname   = "usersdb"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Users Management System")
	})
	e.GET("/users", getUser)
	e.GET("/users/:id", getUserById)
	e.Logger.Fatal(e.Start(":1323"))
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

func getUser(c echo.Context) error {
	uname := c.QueryParam("u")
	pass := c.QueryParam("p")

	dbConnector, err := db.NewConnector("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))

	if err != nil {
		fmt.Println("DB Error: ", err.Error())
		return c.String(http.StatusOK, "Error: "+err.Error())
	}

	usersMgr := users.NewManager(dbConnector)

	//exist,err:=usersMgr.GetAll()
	exist, err := usersMgr.Get(uname, pass)
	fmt.Println(exist)
	if err != nil {
		return c.String(http.StatusOK, "User not found, "+err.Error())
	}

	return c.String(http.StatusOK, "User is found")
}
func getUserById(c echo.Context) error {
	id := c.Param("id")

	dbConnector, err := db.NewConnector("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))

	if err != nil {
		fmt.Println("DB Error: ", err.Error())
		return c.String(http.StatusOK, "Error: "+err.Error())
	}
	usersMgr := users.NewManager(dbConnector)

	//exist,err:=usersMgr.GetAll()
	user, err := usersMgr.GetByID(strconv.Atoi(id))

	if err != nil {
		return c.String(http.StatusOK, "User not found, "+err.Error())
	}

	return c.String(http.StatusOK, "Welcome "+user.Fullname)
}
