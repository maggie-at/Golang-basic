package go_SQL

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

const (
	USERNAME = "root"
	PASSWORD = "alantam."
	HOST     = "127.0.0.1"
	PORT     = "3306"
	DBNAME   = "golang"
)

func InitDB_() *sql.DB {
	dsn := strings.Join([]string{USERNAME, ":", PASSWORD, "@tcp(", HOST, ":", PORT, ")/", DBNAME, "?charset=utf8"}, "")

	// sql.Open(driverName, dataSourceName): 检查连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(10)
	db.SetMaxIdleConns(5)

	// Ping: 建立连接
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection established.")
	return db
}
