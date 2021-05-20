package api

import (
	"fmt"
	"github.com/egylinux/kf_go/users"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserManager interface {
	GetAll() (bool, error)
	IsExist(userName, pasword string) (bool, error)
	Get(userName, pasword string) (*users.User, error)
	Add(user *users.User) (bool, error)
}

type UsersRouter struct {
	manager UserManager
}

func NewRouter(manager UserManager) *echo.Echo {
	router := UsersRouter{manager: manager}
	r := echo.New()
	r.Add(http.MethodGet, "/users", router.GetUser)
	r.Add(http.MethodPost, "/adduser", router.AddUser)
	return r
}

func (u *UsersRouter) GetUser(c echo.Context) error {
	uname := c.QueryParam("u")
	pass := c.QueryParam("p")
c.Bind()
	exist, err := u.manager.Get(uname, pass)
	fmt.Println(exist)
	if err != nil {
		return c.String(http.StatusOK, "User not found, "+err.Error())
	}

	return c.String(http.StatusOK, "User is found")
}

func (u *UsersRouter) AddUser(c echo.Context) error {

	user:= new(users.User)
	if err := c.Bind(user); err != nil {
		fmt.Println("Error",err.Error())
		return err
	}


	_, err := u.manager.Add(user)
	//fmt.Println(exist)
	if err != nil {
		return c.String(http.StatusOK, "Error while saving , "+err.Error())
	}

	return c.String(http.StatusOK, "Saved Successfully")
}

/*func getUserById(c echo.Context) error {
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
*/
