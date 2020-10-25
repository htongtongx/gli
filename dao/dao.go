package dao

import (
	"fmt"

	"github.com/htongtongx/gli/conf"
	"github.com/htongtongx/gli/mongo"
	"github.com/htongtongx/gli/mysql"
)

type Dao struct {
	Mon *mongo.Mongo
	My  *mysql.Mysql
}

func New(c *conf.Config) (dao *Dao) {
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
