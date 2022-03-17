package sqlx

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
)

func TestDbFetchSequence(t *testing.T) {

	newDB := SelectAndLinkDB()

	Tx, err := Db.Beginx()
	key := "TERM_ID_SEQ"
	seqNum, err := DbFetchSequence(Tx, key)
	if err != nil {
		fmt.Println("DbFetchSequence err is :", err)
	}
	sSeqNum := fmt.Sprintf("%06d", seqNum)
	fmt.Printf("key is [%s], value is [%s]", key, sSeqNum)
	Tx.Commit()

	CloseLinkDB(newDB)
}

func TestDbGetOutsideCommodityList(t *testing.T) {

	newDB := SelectAndLinkDB()

	CommodityLists, err := DbGetOutsideCommodityList(Db)
	if err != nil {
		fmt.Println("DbFetchSequence err is :", err)
	}
	fmt.Println(CommodityLists)

	CloseLinkDB(newDB)
}

func TestDbGetCommodityById(t *testing.T) {

	newDB := SelectAndLinkDB()

	CommodityLists, err := DbGetCommodityById(Db)
	if err != nil {
		if !errors.As(err, &sql.ErrNoRows) {
			fmt.Println("DbGetCommodityById err is :", err)
		} else {
			fmt.Println(CommodityLists)
		}
	}

	CloseLinkDB(newDB)

}
