package gconf

type MongoConf struct {
	Enabled  bool   `ini:"enabled"`
	Node     string `ini:"nodes"`
	Pwd      string `ini:"pwd"`
	User     string `ini:"user"`
	Port     int    `ini:"port"`
	QuotaDB  string `ini:"quotaDB"`
	DB       string `ini:"db"`
	Sequence string `ini:"sequence"`
	Auth     string `ini:"auth"`
}

func (m *MongoConf) Verify() bool {
	if m.Node == "" || !m.Enabled {
		return false
	}
	return true
}

type JWTConf struct {
	Secrets string `ini:"secrets"`
	Exp     int64  `ini:"exp"`
	Header  string `ini:"header"` //
	Enabled bool   `ini:"enabled"`
}

type AlipayConf struct {
	SellerID string `ini:"sellerid"`
	PubKey   string `ini:"pubkey"`
	PriKey   string `ini:"prikey"`
	AppID    string `ini:"appid"`
}

type WXpayConf struct {
	AppID     string `ini:"appid"`
	YDAppID   string `ini:"ydappid"`
	MchID     string `ini:"mchid"`
	APIKey    string `ini:"key"`
	AppSecret string `ini:"appsecret"`
	IsProd    bool   `ini:"isprod"`
	// MPAppID   string `ini:"mpAppID"`
	// MBAppID string `ini:"mbAppID"`
}

type LogConf struct {
	Enabled bool   `ini:"enabled"`
	Path    string `ini:"path"`
}

type MysqlConf struct {
	Enabled      bool   `ini:"enabled"`
	Host         string `ini:"host"`
	Pwd          string `ini:"pwd"`
	User         string `ini:"user"`
	Port         int    `ini:"port"`
	DB           string `ini:"db"`
	MaxIdleConns int    `ini:"maxIdleConns"`
	MaxOpenConns int    `ini:"maxOpenConns"`
}

func (m *MysqlConf) Verify() bool {
	if m.Host == "" || !m.Enabled {
		return false
	}
	return true
}

type SqliteConf struct {
	Enabled bool   `ini:"enabled"`
	Pwd     string `ini:"pwd"`
	User    string `ini:"user"`
	DBPath  string `ini:"db_path"`
}

func (m *SqliteConf) Verify() bool {
	if m.DBPath == "" || !m.Enabled {
		return false
	}
	return true
}
