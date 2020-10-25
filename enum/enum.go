package enum

type UserNameType int8

const (
	UserName_NONEXISTS UserNameType = 0   //用户名不存在，并且格式正确
	UserName_INVALID   UserNameType = 1   //用户名格式无效
	UserName_EXISTS    UserNameType = 2   //用户名已存在
	UserName_UNKNOWN   UserNameType = 100 //未知错误
)

type RegisterInfoType int8

const (
	RegisterInfo_SUCCESS RegisterInfoType = 0 //注册用户成功

	RegisterInfo_NameInvalID      RegisterInfoType = 3   //用户名不合法
	RegisterInfo_PwdFail          RegisterInfoType = 4   //用户密码格式错误
	RegisterInfo_EMAILFAIL        RegisterInfoType = 5   //注册邮箱地址不是有效的邮箱格式
	RegisterInfo_NAMEEXIST        RegisterInfoType = 6   //用户名已存在
	RegisterInfo_EMAIlEXIST       RegisterInfoType = 7   //注册邮箱地址已存在
	RegisterInfo_ReferrerNonExist RegisterInfoType = 8   //推荐人用户不存在
	RegisterInfo_QQFail           RegisterInfoType = 9   //请输入正确格式的QQ号码
	RegisterInfo_PhoneFail        RegisterInfoType = 10  //请输入正确格式的电话号码
	RegisterInfo_AgentNonExist    RegisterInfoType = 11  //指定用户代理商管理员不存在
	RegisterInfo_AgentFail        RegisterInfoType = 12  //指定用户代理商不是正确的代理商类型管理员
	RegisterInfo_TypeFail         RegisterInfoType = 13  //用户所属软件注册类型参数错误
	RegisterInfo_UNKNOWN          RegisterInfoType = 100 //未知错误
	RegisterInfo_DBFail           RegisterInfoType = 101 //数据库更新数据时发生错误
)

type LoginType int8

const (
	LoginType_SUCCESS      LoginType = 0   //用户登录成功
	LoginType_NameFail     LoginType = 1   //用户名不合法
	LoginType_VerifyEmpty  LoginType = 2   //验证密码为空字符串
	LoginType_UserNonExist LoginType = 3   //用户名不存在
	LoginType_PwdFail      LoginType = 4   //登录密码不正确
	LoginType_AgentFail    LoginType = 5   //用户登录失败，此用户只能通过代理商发布的特定客户端登录
	LoginType_UNKNOWN      LoginType = 100 //未知错误
	LoginType_DBFail       LoginType = 101 //读取数据库用户信息发生错误
)

type VersionType int8

const (
	VersionType_SUCCESS  VersionType = 0   //获取版本信息成功
	VersionType_TypeFail VersionType = 1   //客户端软件类型错误
	VersionType_NonExist VersionType = 2   //当前软件不存在版本信息
	VersionType_UNKNOWN  VersionType = 100 //未知错误
)

type UserType int8

const (
	UserType_NoneUserType UserType = 0 //无类型
	UserType_Normal       UserType = 1 //普通用户
	UserType_Gold         UserType = 2 //静态会员
	UserType_WhiteGold    UserType = 3 //动态会员
	UserType_Diamond      UserType = 4 //钻石会员
	UserType_Platinum     UserType = 5 //铂金会员
	UserType_Bronze       UserType = 6 //青铜会员
)
