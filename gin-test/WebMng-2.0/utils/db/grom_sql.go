package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

// 数据库链接信息
type MySqlDBInfo struct {
	IP        string
	Port      string
	User      string
	Password  string
	Name      string
	CharSet   string
	ParseTime bool
}

// 工厂初始化
func NewMySqlDBInfo(ip string, port string, user string, password string, name string, charSet string, parseTime bool) *MySqlDBInfo {
	return &MySqlDBInfo{
		IP:        ip,
		Port:      port,
		User:      user,
		Password:  password,
		Name:      name,
		CharSet:   charSet,
		ParseTime: parseTime,
	}
}

// 初始化链接数据库
func (db *MySqlDBInfo) InitMysql() (GromDb *gorm.DB, err error) {

	// "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	var sTmpBuf strings.Builder
	sTmpBuf.WriteString(db.User)
	sTmpBuf.WriteString(":")
	sTmpBuf.WriteString(db.Password)
	sTmpBuf.WriteString("@tcp(")
	sTmpBuf.WriteString(db.IP)
	sTmpBuf.WriteString(":")
	sTmpBuf.WriteString(db.Port)
	sTmpBuf.WriteString(")/")
	sTmpBuf.WriteString(db.Name)
	sTmpBuf.WriteString("?charset=")
	sTmpBuf.WriteString(db.CharSet)

	if db.ParseTime {
		sTmpBuf.WriteString("&parseTime=True&loc=Local")
	}
	// fmt.Println(sTmpBuf.String())
	GromDb, err = gorm.Open(mysql.Open(sTmpBuf.String()), &gorm.Config{})
	return
}
