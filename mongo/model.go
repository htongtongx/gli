package mongo

type IMongo interface {
	Table() string
	SetID(int64)
	UniqueKey() string
	UniqueValue() interface{}
	IsAI() bool
}

type IUser interface {
	IMongo
	AuthPwd(string, string) bool
}
