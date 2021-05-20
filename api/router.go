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
}

type UsersRouter struct {
	manager UserManager
}

func NewRouter(manager UserManager) *echo.Echo {
	router := UsersRouter{manager: manager}
	r := echo.New()
	r.Add(http.MethodGet, "/users", router.GetUser)
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
