package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestClient(t *testing.T) {
	db := GetDB()

	CreateTable(db)
	//InsertData(db)
	//QueryOne(db)
	//QueryMulti(db)
	//UpdateData(db)
	//DeleteData(db)
}
