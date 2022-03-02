package sqlx

import (
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
