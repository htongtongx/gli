package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/htongtongx/gli/gconf"
)

type Sqlite struct {
	c   *gconf.SqliteConf
	Cli *sql.DB
}

func NewSqlite(c *gconf.SqliteConf) (s *Sqlite, err error) {
	if !c.Verify() {
		fmt.Println("sqlite配置未启用.")
		return
	}
	s = new(Sqlite)
	s.c = c

	db, err := sql.Open("sqlite3", c.DBPath)
	if err != nil {
		return nil, err
	}
	s.Cli = db
	fmt.Println("sqlite连接成功!")
	return
}
