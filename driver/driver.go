package driver

import "database/sql"

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectSQL(url string) (*DB, error) {
	conn, err := sql.Open("mysql", url)
	if err != nil {
		panic(err)
	}

	dbConn.SQL = conn
	return dbConn, err
}
