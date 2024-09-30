package dao

import (
	"fmt"

	"github.com/htongtongx/gli/gconf"
	"github.com/htongtongx/gli/mongo"
	"github.com/htongtongx/gli/mysql"
	"github.com/htongtongx/gli/sqlite"
)

type Dao struct {
	Mon  *mongo.Mongo
	My   *mysql.Mysql
	Lite *sqlite.Sqlite
}

func (d *Dao) InitMy(myConf *gconf.MysqlConf) (err error) {
	my, err := mysql.NewMysql(myConf)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	d.My = my
	return
}

func (d *Dao) InitLite(sqliteConf *gconf.SqliteConf) (err error) {
	lite, err := sqlite.NewSqlite(sqliteConf)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	d.Lite = lite
	return
}

func New(c *gconf.Config) (dao *Dao) {
	mon, err := mongo.NewMongo(&c.Mongo)
	if err != nil {
		fmt.Println("mongo初始化失败：" + err.Error())
	}
	my, err := mysql.NewMysql(&c.Mysql)
	if err != nil {
		fmt.Println("mysql初始化失败：" + err.Error())
	}
	dao = &Dao{
		Mon: mon,
		My:  my,
	}
	return
}
