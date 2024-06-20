package gconf

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	// "github.com/htongtongx/gli/sms"
	"github.com/htongtongx/gli/util"
	"gopkg.in/ini.v1"
)

var runModeMap map[RunModeType]string

type RunModeType string

type Config struct {
	Mongo  MongoConf  `ini:"mongo"`
	Jwt    JWTConf    `ini:"jwt"`
	Alipay AlipayConf `ini:"alipay"`
	WXpay  WXpayConf  `ini:"wxpay"`
	Log    LogConf    `ini:"log"`
	Mysql  MysqlConf  `ini:"mysql"`
	// AliSMS    sms.AliSMSConf  `ini:"alisms"`
	// Redis     redis.RedisConf `ini:"redis"`
	IsProd    bool
	IsWindows bool
	IsTests   bool
	RunMode   RunModeType
	INICfg    *ini.File
}

// func (c *Config) NewJwt(key string) (jwt *JWTConf, err error) {
// 	jwt = &JWTConf{}
// 	err = c.INICfg.Section(key).MapTo(jwt)
// 	return
// }

// func (c *Config) NewAlipay(key string) (alipayCfg *AlipayConf, err error) {
// 	alipayCfg = &AlipayConf{}
// 	err = c.INICfg.Section(key).MapTo(alipayCfg)
// 	return
// }

// func (c *Config) NewWXpay(key string) (wxpayCfg *WXpayConf, err error) {
// 	wxpayCfg = &WXpayConf{}
// 	err = c.INICfg.Section(key).MapTo(wxpayCfg)
// 	return
// }

// func (c *Config) NewMysql(key string) (mysqlCfg *MysqlConf, err error) {
// 	mysqlCfg = &MysqlConf{}
// 	err = c.INICfg.Section(key).MapTo(mysqlCfg)
// 	return
// }

// func (c *Config) NewAliSMS(key string) (alismsCfg *sms.AliSMSConf, err error) {
// 	alismsCfg = &sms.AliSMSConf{}
// 	err = c.INICfg.Section(key).MapTo(alismsCfg)
// 	return
// }

// func (c *Config) NewRedis(key string) (redisCfg *redis.RedisConf, err error) {
// 	redisCfg = &redis.RedisConf{}
// 	err = c.INICfg.Section(key).MapTo(redisCfg)
// 	return
// }

const (
	RUNMODE_DEV  RunModeType = "dev"
	RUNMODE_PROD RunModeType = "prod"
	RUNMODE_TEST RunModeType = "testing"
)

func init() {
	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	dir := filepath.Dir(path)

	runModeMap = map[RunModeType]string{
		RUNMODE_DEV:  "./dev.ini",
		RUNMODE_PROD: dir + "/prod.ini",
		RUNMODE_TEST: "../cmd/test.ini",
	}
}

func New(iniConf *ini.File, v interface{}) (c *Config, err error) {
	c = &Config{}
	c.RunMode = ChcekRunMode()
	switch c.RunMode {
	case RUNMODE_PROD:
		c.IsProd = true
	case RUNMODE_TEST:
		c.IsTests = true
	}
	err = iniConf.MapTo(v)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c.INICfg = iniConf
	// c.enabledLog()
	return
}

func ChcekRunMode() RunModeType {
	rn := os.Getenv("RunMode")
	if rn == "prod" {
		return RUNMODE_PROD
	}
	if util.IsInTests() {
		return RUNMODE_TEST
	}
	return RUNMODE_DEV
}

func loadINIByMode(pathMap map[RunModeType]string) (*ini.File, error) {
	runMode := ChcekRunMode()
	fmt.Printf("当前加载%s.ini配置文件...", runMode)
	path, ok := pathMap[runMode]
	if !ok {
		return nil, errors.New("Fail to runMode: " + string(runMode))
	}
	c, err := ini.Load(path)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return nil, err
	}
	return c, nil
}

// path: ini 文件路径  v:需要加载配置的结构体指针
func LoadINI(path string, v interface{}) (c *Config, err error) {
	var iniCfg *ini.File
	if path == "" {
		iniCfg, err = loadINIByMode(runModeMap)
	} else {
		fmt.Println("加载自定义文件：" + path)
		iniCfg, err = ini.Load(path[0])
	}
	if err != nil {
		return
	}
	return New(iniCfg, v)
}

// 启用log
// func (c *Config) enabledLog() {
// 	if !c.Log.Enabled {
// 		return
// 	}
// 	if c.Log.Path == "" {
// 		year := time.Now().Year()
// 		month := time.Now().Format("01")
// 		day := time.Now().Day()
// 		c.Log.Path = "./log" + parse.ToString(year) + "-" + month + "-" + parse.ToString(day) + ".log"
// 	}
// 	logFile, err := os.OpenFile(c.Log.Path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
// 	if err != nil {
// 		log.Println("启用日志文件失败:" + err.Error())
// 	} else {
// 		log.SetOutput(logFile) // 将文件设置为log输出的文件
// 		// log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
// 		log.SetFlags(log.LstdFlags | log.Lshortfile)
// 	}
// }

func (c *Config) Get(section, key string) *ini.Key {
	return c.INICfg.Section(section).Key(key)
}

func (c *Config) GetString(section, key string) string {
	return c.Get(section, key).String()
}

// func (c *Config) NewMongo(key string) (mongoCfg *MongoConf, err error) {
// 	mongoCfg = &MongoConf{}
// 	err = c.INICfg.Section(key).MapTo(mongoCfg)
// 	return
// }
