package dao

import (
	"fmt"

	"black-hole.com/modules/conf"
	"black-hole.com/modules/mongo"
	"black-hole.com/modules/mysql"
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
