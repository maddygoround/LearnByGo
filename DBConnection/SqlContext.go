package SQLCtx

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var SQLdb *sql.DB

func InitDB() {
	SQLdb, _ = sql.Open("mysql", "YOUR_DATABASE_STRING")
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
