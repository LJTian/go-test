package conf

import (
	"fmt"
	"github.com/Unknwon/goconfig"
)

var (
	CFG *Conf
)

//Conf 配置文件结构体
type Conf struct {
	Ip            string // 服务器ip地址 默认本机 127.0.0.1
	ServerPort    string // 服务器端口  默认 14333
	WebPort       string // 页面端口  默认24333
	DbIp          string // 数据库IP
	DbName        string // 数据库名称
	DbUserName    string // 数据用户名称
	DbPassWord    string // 数据库密码
	DbPoolNum     string // 数据库连接数
	LogLevel      int    // 日志级别
	LogFilePath   string // 日志文件
	PidFilePath   string // pid文件
	FIFOLen       int64  // 管道长度
	SendThreadNum int    //发送线程数
}

//单例工厂
func NewConf(confFileName string) *Conf {

	if CFG == nil {
		CFG = new(Conf)
		CFG.InitConf(confFileName)
	}

	return CFG
}

//InitConf 初始化配置文件
func (conf *Conf) InitConf(confFileName string) {
	cfg, err := goconfig.LoadConfigFile(confFileName)
	if err != nil {
		fmt.Printf("无法加载配置文件：%s", err)
	}
	conf.Ip = cfg.MustValue("Server", "ip")
	conf.ServerPort = cfg.MustValue("Server", "serverPort")
	conf.WebPort = cfg.MustValue("Server", "webPort")
	conf.DbIp = cfg.MustValue("Database", "dbIp")
	conf.DbName = cfg.MustValue("Database", "dbName")
	conf.DbUserName = cfg.MustValue("Database", "dbUserName")
	conf.DbPassWord = cfg.MustValue("Database", "dbPassWord")
	conf.DbPoolNum = cfg.MustValue("Database", "dbPoolNum")
	conf.LogLevel = cfg.MustInt("Log", "logLevel")
	conf.LogFilePath = cfg.MustValue("Log", "logFilePath")
	conf.PidFilePath = cfg.MustValue("Log", "pidFilePath")
	conf.FIFOLen = cfg.MustInt64("Server", "FIFOLEN")
	conf.SendThreadNum = cfg.MustInt("Server", "SendThreadNum")
}

//RestConf 返回配置文件
func (conf *Conf) RestConf() *Conf {
	return conf
}

//PrintConf 打印查看配置文件
func (conf *Conf) PrintConf() string {

	return fmt.Sprintf("\n Ip :[%s] \n "+
		"ServerPort :[%s] \n WebPort :[%s] \n "+
		"DbIp :[%s] \n DbName :[%s] \n "+
		"DbUserName :[%s] \n DbPassWord :[%s] \n "+
		"DbPoolNum :[%s] \n LogLevel :[%s] \n "+
		"LogFilePath :[%s] \n PidFilePath :[%s] \n"+
		"FIFOLen:[%d]\n",
		conf.Ip,
		conf.ServerPort, conf.WebPort,
		conf.DbIp, conf.DbName,
		conf.DbUserName, conf.DbPassWord,
		conf.DbPoolNum, conf.LogLevel,
		conf.LogFilePath, conf.PidFilePath,
		conf.FIFOLen)
}
