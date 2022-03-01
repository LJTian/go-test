package sqlx

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
)

func SelectAndLinkDB() *DBinfo {
	fmt.Println("开始测试~")
	newDB := NewDBinfo("192.168.0.87", "3306", "mis_go_new", "root", "p@ssw0rd")
	newDB.InitDB()
	fmt.Println("连接数据库成功！")
	return newDB
}

func CloseLinkDB(newDB *DBinfo) {
	newDB.CloseDB()
}

/* 系统序列表 */
type Sequence struct {
	Name         string `db:"name"`          // 名称
	CurrentValue string `db:"current_value"` // 当前值
	Increment    int64  `db:"increment"`     // 每次自增值
	MaxValue     int64  `db:"max_value"`     // 最大值
	RecycleFlag  string `db:"recycle_flag"`  // 是否循环 Y-循环, N-不循环
}

// 根据Key取序列号
func DbFetchSequence(pstDb *sqlx.Tx, sName string) (object int64, err error) {

	var nextSeqVal int64

	sequence := new(Sequence)

	//sqlStr := "select * from sequence where name = ? "
	//err = pstDb.QueryRowx(sqlStr, sName).StructScan(sequence)
	//if err != nil {
	//	return
	//}

	sqlStr := "select * from `sequence` where name  = ? "
	err = pstDb.Get(sequence, sqlStr, sName)
	if err != nil {
		return
	}
	// 获取当前值
	lCurSeqVal, _ := strconv.ParseInt(sequence.CurrentValue, 10, 64)
	object = lCurSeqVal

	/* 序列达到最大值: 重置、或报错 */
	if lCurSeqVal == sequence.MaxValue {
		if sequence.RecycleFlag == "Y" {
			nextSeqVal = 1
		} else {
			err = fmt.Errorf("非循环序列[%s]已达最大值[%d]!", sName, sequence.MaxValue)
			return
		}
	} else {
		nextSeqVal = lCurSeqVal + sequence.Increment
		var result sql.Result
		result, err = pstDb.Exec("UPDATE `SEQUENCE` "+
			"SET CURRENT_VALUE = ? "+
			"WHERE NAME = ? ",
			nextSeqVal, sName)
		if err != nil {
			return
		}
		_, err = result.RowsAffected()
		if err != nil {
			return
		}
	}

	return
}
