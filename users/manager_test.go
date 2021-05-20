package users

import (
	"fmt"
	import "github.com/stretchr/testify/mock"
)
type UsersTest struct {
	manager UserManager
	mock.Mock
}

type UserManager interface {
	GetAll() (bool, error)
	IsExist(userName, pasword string) (bool, error)
	Get(userName, pasword string) (*User, error)
	Add(user *User) (bool, error)
}

func (u *UsersTest) GetUser() {
	ret := u.Called()
	expected := User{Id: 1}

	actual,_ := u.manager.Get("Admin","123")
	if actual.Id != expected.Id {
		fmt.Printf("Test failed, expected: '%d', got:  '%d'", expected.Id, actual.Id)
	}
}