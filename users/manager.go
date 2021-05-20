package users

import (
	 "database/sql"
	"fmt"
	_ "github.com/lib/pq"
	_ "github.com/urfave/cli/v2"
)

// Manager to manage user
type Manager struct {
	connector DBConnector
}

// DBConnector contains dataAccess functionalities
type DBConnector interface {
	Get(dest interface{}, query string, args ...interface{}) error
	Select(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
}

// NewManager inistantiate user manager
func NewManager(connector DBConnector) *Manager {
	return &Manager{
		connector: connector,
	}
}

/*func init() {
	//connectionstring
	psqlconn := ConnectionString()
	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	// close database
	defer db.Close()
	// check db
	err = db.Ping()
	CheckError(err)
	//Check if user table exist , false = create table
	CheckUserTable(psqlconn)
}*/

func (m *Manager) GetAll() (bool, error) {
	q := fmt.Sprintf(`Select * from userstb `)

	usrs := make([]User, 0)

	if err := m.connector.Select(&usrs, q); err != nil {

		return false, err
	}
	fmt.Println(usrs)
	return true, nil
}
func (m *Manager) IsExist(userName, pasword string) (bool, error) {
	q := fmt.Sprintf(`Select * from userstb where username='%s' and password ='%s'`, userName, pasword)
	fmt.Println(q)
	usrs := make([]User, 0)

	if err := m.connector.Select(&usrs, q); err != nil {

		return false, err
	}
	fmt.Println(usrs)
	return true, nil
}

func (m *Manager) GetByID(Id int, err error) (*User, error) {
	q := fmt.Sprintf(`Select * from userstb where Id=%d `, Id)
	user := &User{}
	if err := m.connector.Get(user, q); err != nil {
		return &User{}, err
	}

	return user, nil
}
func (m *Manager) Get(userName, pasword string) (*User, error) {
	q := fmt.Sprintf(`Select * from userstb where username='%s' and password ='%s'`, userName, pasword)
	user := &User{}
	if err := m.connector.Get(user, q); err != nil {
		return &User{}, err
	}

	return user, nil
}
func (m *Manager) Add(user *User) (bool, error) {
	q := fmt.Sprintf(`Insert Into userstb(username,password,fullname) values('%s','%s','%s')`, user.Username, user.Password,user.Fullname)

	if _,err := m.connector.Exec(q); err != nil {
		return false, err
	}

	return true, nil
}

/*func CheckUserTable(psqlconn string) {

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	strsql := `SELECT count(*) as counter FROM pg_catalog.pg_tables WHERE schemaname != 'pg_catalog' AND  schemaname != 'information_schema' And tablename = 'userstb';`
	rows, e := db.Query(strsql)
	CheckError(e)
	counter := 0

	for rows.Next() {
		err = rows.Scan(&counter)
		CheckError(err)

	}

	if counter == 0 {
		strsql = `Create table userstb(id SERIAL PRIMARY KEY,fullName varchar(255),username Varchar(255),password varchar(50))`
		_, e = db.Exec(strsql)
		CheckError(err)
		fmt.Printf("Users table created\n")
	}
	defer db.Close()

}
*/
