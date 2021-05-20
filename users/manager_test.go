package users

import (
	"errors"
	"github.com/egylinux/kf_go/users/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	// Arrange
	fakeConn := &mocks.DBConnector{}
	q := "Insert Into userstb(username,password,fullname) values('Ayman','aaaa','Ayman Hassan')"
	fakeConn.On("Exec", q).Return(nil, nil)
	q2 := "Insert Into userstb(username,password,fullname) values('Ali','aaaa','Ayman Hassan')"
	fakeConn.On("Exec", q2).Return(nil, errors.New("error saving"))

	mgr := NewManager(fakeConn)
	usr := &User{Username: "Ayman", Password: "aaaa", Fullname: "Ayman Hassan"}

	// Act
	ok, err := mgr.Add(usr)

	// Assert
	assert.Equal(t, true, ok)
	assert.Equal(t, nil, err)

	usr = &User{Username: "Ali", Password: "aaaa", Fullname: "Ayman Hassan"}
	// Act

	ok, err = mgr.Add(usr)

	// Assert
	assert.Equal(t, false, ok)
	assert.False(t, ok)
	assert.Error(t, err)
}
