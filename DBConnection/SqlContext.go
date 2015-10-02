package SQLCtx

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var SQLdb *sql.DB

func InitDB() {
	SQLdb, _ = sql.Open("mysql", "root:root123@tcp(172.16.1.201:3306)/gotest") //sql.Open("mssql", cObj.ConnectionString)
	err := SQLdb.Ping()
	if err != nil {
		fmt.Println("Connection Failed")
	} else {
		fmt.Println("Connection Started")
	}
}

func DB() *sql.DB {
	return SQLdb
}
