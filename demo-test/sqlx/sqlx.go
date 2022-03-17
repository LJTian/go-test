package sqlx

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
)

func SelectAndLinkDB() *DBinfo {
	fmt.Println("开始测试~")
	newDB := NewDBinfo("192.168.0.182", "3306", "go_mis", "root", "p@ssw0rd")
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

/* 商品基本信息表 */
type Commodity struct {
	Id           sql.NullInt64   `db:"id"`            // 主键ID
	CommodNo     sql.NullString  `db:"commod_no"`     // 商品编码
	CommodName   sql.NullString  `db:"commod_name"`   // 商品名称
	TypeName     sql.NullString  `db:"type_name"`     // 商品分类名称
	BrandName    sql.NullString  `db:"brand_name"`    // 品牌名称
	MeasureUnit  sql.NullString  `db:"measure_unit"`  // 计量单位
	Norms        sql.NullString  `db:"norms"`         // 商品规格
	Price        sql.NullFloat64 `db:"price"`         // 标准售价
	SimpleCode   sql.NullString  `db:"simple_code"`   // 助记码
	OutsideSale  sql.NullString  `db:"outside_sale"`  // 室外销售标志
	SafeFlag     sql.NullString  `db:"sale_flag"`     // 停售标识
	RedeemFlag   sql.NullString  `db:"redeem_flag"`   // 积分换购标志
	RedeemPoints sql.NullInt32   `db:"redeem_points"` // 换购积分值
}

func DbGetCommodityById(pstDb *sqlx.DB) (object *Commodity, err error) {

	var commodity Commodity
	object = &commodity

	sqlBuff := fmt.Sprintf(
		"select * from COMMODITY where id = %d", 10000)

	err = pstDb.QueryRowx(sqlBuff).StructScan(object)
	fmt.Println(err)

	return

}

func DbGetOutsideCommodityList(pstDb *sqlx.DB) (object *[]Commodity, err error) {

	var commodity Commodity
	commoditys := make([]Commodity, 0)
	sqlBuff := fmt.Sprintf("SELECT * " +
		"FROM COMMODITY " +
		"WHERE SALE_FLAG  = '0' " +
		"AND ID IN " +
		"(SELECT COMMOD_ID FROM STATION_COMMODITY_CTRL WHERE OUT_SALE_FLAG = 'Y' AND STOP_FLAG = 0) ")

	rows, err := pstDb.Queryx(sqlBuff)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		err = rows.StructScan(&commodity)
		if err != nil {
			rows.Close()
			return nil, err
		}

		commoditys = append(commoditys, commodity)
	}
	object = &commoditys
	rows.Close()
	return
}
