package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/pkg/errors"
	"strings"
)

// NewConn 创建一个数据库连接
func NewConn(dbURL string) (*gorm.DB, error) {
	dialect := "mysql"
	if strings.HasSuffix(dbURL, ".db") {
		dialect = "sqlite3"
	}
	conn, err := gorm.Open(dialect, dbURL)
	if err != nil {
		return nil, errors.Wrap(err, "database connection error")
	}

	return conn, nil

}
