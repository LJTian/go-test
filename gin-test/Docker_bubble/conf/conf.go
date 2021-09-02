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
	Ip       string // 数据库IP
	Port     string // 数据库端口号
	Name     string // 数据库名称
	UserName string // 数据用户名称
	PassWord string // 数据库密码
	PoolNum  string // 数据库连接数
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

	conf.Ip = cfg.MustValue("Database", "dbIp")
	conf.Port = cfg.MustValue("Database", "dbPort")
	conf.Name = cfg.MustValue("Database", "dbName")
	conf.UserName = cfg.MustValue("Database", "dbUserName")
	conf.PassWord = cfg.MustValue("Database", "dbPassWord")
	conf.PoolNum = cfg.MustValue("Database", "dbPoolNum")
}

//RestConf 返回配置文件
func (conf *Conf) RestConf() *Conf {
	return conf
}

//PrintConf 打印查看配置文件
func (conf *Conf) PrintConf() string {

	return fmt.Sprintf(
		" DbIp :[%s] \n DbPort :[%s] \n DbName :[%s] \n "+
			"DbUserName :[%s] \n DbPassWord :[%s] \n "+
			"DbPoolNum :[%s] \n",
		conf.Ip, conf.Port, conf.Name,
		conf.UserName, conf.PassWord,
		conf.PoolNum)
}
