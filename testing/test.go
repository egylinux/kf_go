package testing

import (
	"fmt"
	"github.com/egylinux/kf_go/users"
	"testing"
)
type UsersTest struct {
	manager UserManager
}

type UserManager interface {
	GetAll() (bool, error)
	IsExist(userName, pasword string) (bool, error)
	Get(userName, pasword string) (*users.User, error)
	Add(user *users.User) (bool, error)
}

func (u *UsersTest) getUser(t *testing.T) {
	expected := users.User{Id: 1}

	actual,_ := u.manager.Get("Admin","123")
	if actual.Id != expected.Id {
		fmt.Printf("Test failed, expected: '%d', got:  '%d'", expected.Id, actual.Id)
	}
}