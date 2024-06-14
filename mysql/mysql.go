package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/htongtongx/gli/gconf"
	"github.com/jmoiron/sqlx"
)

var Client *Mysql

type Mysql struct {
	Cli  *sqlx.DB
	Node string
	Pwd  string
	User string
	c    *gconf.MysqlConf
}

func NewMysql(c *gconf.MysqlConf) (m *Mysql, err error) {
	if !c.Verify() {
		log.Println("mysql配置未启用.")
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", c.User, c.Pwd, c.Host, c.Port, c.DB)
	m = new(Mysql)
	m.Cli, err = sqlx.Open("mysql", dsn)
	if err != nil {
		log.Println("mysql连接失败：" + err.Error())
		return
	}
	if c.MaxIdleConns > 0 {
		m.Cli.SetMaxIdleConns(c.MaxIdleConns)
	}
	if c.MaxOpenConns > 0 {
		m.Cli.SetMaxOpenConns(c.MaxOpenConns)
	}
	log.Println("msyql连接成功!")
	return
}

func (m *Mysql) Get(dest interface{}, sql string, args ...interface{}) (has bool, err error) {
	err = m.Cli.Get(dest, sql, args...)
	if err != nil && err.Error() == "sql: no rows in result set" {
		has = false
		err = nil
		return
	}
	has = true
	return
}

func (m *Mysql) GetList(dest interface{}, sql string, args ...interface{}) (err error) {
	err = m.Cli.Select(dest, sql, args...)
	return
}

func (m *Mysql) Exec(sql string, args ...interface{}) (result sql.Result, err error) {
	result, err = m.Cli.Exec(sql, args...)
	return
}

type totalData struct {
	Total int64 `db:"total"`
}

func (m *Mysql) CountRows(table, filter string, args ...interface{}) (count int64, err error) {
	td := &totalData{}
	sql := fmt.Sprintf("select count(*) total from %s where %s", table, filter)
	_, err = m.Get(td, sql, args...)
	if err != nil {
		return
	}
	count = td.Total
	return
}
