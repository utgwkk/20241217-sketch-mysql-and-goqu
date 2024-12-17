package internal

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

func ConnectDB() (*sql.DB, error) {
	mysqlCfg := mysql.NewConfig()
	mysqlCfg.Addr = fmt.Sprintf(
		"%s:%s",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
	)
	mysqlCfg.Net = "tcp"
	mysqlCfg.User = os.Getenv("DATABASE_USER")
	mysqlCfg.Passwd = os.Getenv("DATABASE_PASSWORD")
	mysqlCfg.DBName = os.Getenv("DATABASE_DBNAME")

	return sql.Open("mysql", mysqlCfg.FormatDSN())
}
