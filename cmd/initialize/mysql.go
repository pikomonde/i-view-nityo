package initialize

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pikomonde/i-view-nityo/model"
)

func NewMySQL(config model.MySQLConfig) (*sqlx.DB, error) {
	return sqlx.Connect("mysql", fmt.Sprintf(
		"%s:%s@/%s",
		config.Username,
		config.Password,
		// config.Host,
		// config.Port,
		config.DBName,
	))
}
