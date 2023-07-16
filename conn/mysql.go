package conn

import (
	"database/sql"
	"fmt"
)

func MysqlConnection(addr, dbname, username, password string) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		username, password, addr, dbname,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
