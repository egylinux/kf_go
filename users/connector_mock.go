package users

import (
	"database/sql"
	"errors"
	"strings"
)

type ConnectorMock struct {
}

func (c *ConnectorMock) Get(dest interface{}, query string, args ...interface{}) error {
	if strings.Contains(query, "admin") {
		return errors.New("reading user expection ")
	}
	return nil
}

func (c *ConnectorMock) Select(dest interface{}, query string, args ...interface{}) error {
	return nil
}

func (c *ConnectorMock) Exec(q string, args ...interface{}) (sql.Result, error) {
	if strings.Contains(q, "Ali") {
		return nil, errors.New("saving expection ")
	}
	return nil, nil
}
