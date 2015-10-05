package SQLCtx

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_"github.com/denisenkom/go-mssqldb"
)

var SQLdb *sql.DB

func InitDB() {
	SQLdb, _ = sql.Open("mssql", "server=localhost;user id=sa;password=sangreal;database=GOTest;port=1433")
	err := SQLdb.Ping()
	if err != nil {
		fmt.Println("Connection Failed",err.Error())
	} else {
		fmt.Println("Connection Started")
	}
}

