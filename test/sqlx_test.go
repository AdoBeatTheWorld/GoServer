package test

import (
	"fmt"
	"testing"

	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	database, err := sqlx.Open("mysql", "root:12345@tcp(192.168.2.127:3306)/db_account")
	if err != nil {
		fmt.Println(err)
	}
	Db = database
}

func TestSqlx(t *testing.T) {
	r, err := Db.Exec("insert into android_config()")
}
