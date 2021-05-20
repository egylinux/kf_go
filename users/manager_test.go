package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd(t *testing.T) {
	// Arrange
	fakeConn := &ConnectorMock{}
	mgr := NewManager(fakeConn)
	usr := &User{Username: "Ayman", Password: "aaaa", Fullname: "Ayman Hassan"}

	// Act
	ok, err := mgr.Add(usr)

	// Assert
	assert.Equal(t, true, ok)
	assert.Equal(t, nil, err)

	usr=&User{Username: "Ali", Password: "aaaa", Fullname: "Ayman Hassan"}
	// Act

	ok, err = mgr.Add(usr)

	// Assert
	assert.Equal(t, false, ok)
	assert.False(t, ok)
	assert.Error(t, err)
}