package sql

import (
	"TYServer/utils/db"
	. "TYServer/utils/log"
	"errors"
	"github.com/jmoiron/sqlx"
)

type UserInfo struct {
	Id      int    `db:"Id"`
	Name    string `db:"Name"`
	Phone   string `db:"Phone"`
	Status  int    `db:"Status"`
	ProDate string `db:"Procdate"`
}

// NewUserInfo 创建一个用户信息
func NewUserInfo(name string, phone string, status int) *UserInfo {
	return &UserInfo{
		Name:   name,
		Phone:  phone,
		Status: status,
	}
}

// AddUsertoDB 添加用户信息到数据库
func (userinfo *UserInfo) AddUsertoDB() error {

	sqlbuf := `INSERT INTO t_user_info ( Name, Phone, Status, Procdate ) VALUES (?, ? , ? , sysdate() )`

	sql := db.Db.MustExec(sqlbuf, userinfo.Name, userinfo.Phone, userinfo.Status)
	if sql != nil {
		TlogPrintf(LOG_ERROR, "sql[%s]执行失败", sqlbuf)
		return errors.New("AddUsertoDB exec fail!")
	}
	return nil
}

// GetUserInfoById 根据id获取用户信息
func (userinfo *UserInfo) GetUserInfoById(id int) (UserInfo, error) {

	var userData UserInfo
	//err := db.Db.QueryRowx("SELECT Id, Name, Phone, Status, Procdate FROM t_user_info WHERE Id = ? ", id).StructScan(&userData)
	//if err != nil {
	//	//TlogPrintf(LOG_ERROR, "sql[%s]执行失败", sqlbuf)
	//	return userData, err
	//}

	//err := db.Db.Get(&userData, "SELECT Id, Name, Phone, Status, Procdate FROM t_user_info WHERE Id = ? ", id)
	//if err != nil {
	//	return userData, err
	//}

	stmt, err := db.Db.Preparex(`SELECT Id, Name, Phone, Status, Procdate FROM t_user_info WHERE Id = ? `)
	if err != nil {
		return userData, err
	}
	err = stmt.Get(&userData, 1)

	return userData, err
}

// GetUserInfoById 根据id获取用户信息
func (userinfo *UserInfo) GetUserInfoByIn() (UserInfo, error) {

	var userData UserInfo
	//err := db.Db.QueryRowx("SELECT Id, Name, Phone, Status, Procdate FROM t_user_info WHERE Id = ? ", id).StructScan(&userData)
	//if err != nil {
	//	//TlogPrintf(LOG_ERROR, "sql[%s]执行失败", sqlbuf)
	//	return userData, err
	//}

	//err := db.Db.Get(&userData, "SELECT Id, Name, Phone, Status, Procdate FROM t_user_info WHERE Id = ? ", id)
	//if err != nil {
	//	return userData, err
	//}
	var levels = []int{1, 1022, 10333}
	query, args, err := sqlx.In("SELECT * FROM t_user_info WHERE Id IN (?);", levels)
	query = db.Db.Rebind(query)
	rows, err := db.Db.Query(query, args...)
	if err != nil {
		return userData, err
	}

	//stmt, err := db.Db.Preparex(`SELECT Id, Name, Phone, Status, Procdate FROM t_user_info WHERE Id = ? `)
	//if err != nil {
	//	return userData, err
	//}
	err = rows.Scan(&userData.Id, &userData.Name, &userData.Phone, &userData.Status, &userData.ProDate)

	return userData, err
}

// GetUserInfoAll 获取全部用户信息
func (userinfo *UserInfo) GetUserInfoAll() ([]UserInfo, error) {

	var userDatas []UserInfo

	err := db.Db.Select(&userDatas, "SELECT Id, Name, Phone, Status, Procdate FROM t_user_info")
	if err != nil {
		return userDatas, err
	}
	//rows, err := db.Db.Queryx("SELECT Id, Name, Phone, Status, Procdate FROM t_user_info ")
	//if err != nil {
	//	return userDatas, err
	//}
	//for rows.Next() {
	//
	//	var userdata UserInfo
	//	err = rows.StructScan(&userdata)
	//	if err != nil {
	//		return userDatas, err
	//	}
	//	userDatas = append(userDatas, userdata)
	//}
	return userDatas, nil
}

//UpdateUserInfoById 修改用户信息
func (userinfo *UserInfo) UpdateUserInfoById(userinfoData *UserInfo, id int) error {

	defer func() {
		if err := recover(); err != nil {
			TlogPrintf(LOG_ERROR, "UpdateUserInfoById 执行失败")
		}
	}()
	sqlbuf := `update t_user_info 
				set 
				 Name =  ?
    			,Phone =  ?
				,Status =  ?
				where id = ?`

	db.Db.MustExec(sqlbuf, userinfoData.Name, userinfoData.Phone, userinfoData.Status, id)
	return nil
}
