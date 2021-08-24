package db

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

var (
	Db  *sqlx.DB
	err error
)

//DBinfo 数据库信息
type DBinfo struct {
	Ip       string
	Post     string
	DBName   string
	UserName string
	PassWord string
}

//NewDBinfo 创建一个新的数据库信息
func NewDBinfo(ip string, post string, dbName string, userName string, passWord string) *DBinfo {

	dbinfo := &DBinfo{
		Ip:       ip,
		Post:     post,
		DBName:   dbName,
		UserName: userName,
		PassWord: passWord,
	}
	return dbinfo
}

//InitDB 初始化数据库
func (dbinfo *DBinfo) InitDB() {

	var sTmpBuf strings.Builder
	sTmpBuf.WriteString(dbinfo.UserName)
	sTmpBuf.WriteString(":")
	sTmpBuf.WriteString(dbinfo.PassWord)
	sTmpBuf.WriteString("@tcp(")
	sTmpBuf.WriteString(dbinfo.Ip)
	sTmpBuf.WriteString(":")
	sTmpBuf.WriteString(dbinfo.Post)
	sTmpBuf.WriteString(")/")
	sTmpBuf.WriteString(dbinfo.DBName)
	sTmpBuf.WriteString("?charset=utf8")

	DBUrl := sTmpBuf.String()
	Db, err = sqlx.Open("mysql", DBUrl)
	if err != nil {
		panic(err.Error())
	}

	err = Db.Ping()
	if err != nil {
		panic(err.Error())
	}
}

//SetConnPool 设置连接池
func (dbinfo *DBinfo) SetConnPool(ldle int, open int) error {

	if ldle > 0 {
		Db.SetMaxIdleConns(ldle)
	}

	if open > 0 {
		Db.SetMaxOpenConns(open)
	}
	return nil
}

//CloseDB 关闭数据库
func (dbinfo *DBinfo) CloseDB() {
	Db.Close()
}
