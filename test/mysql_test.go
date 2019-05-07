package test

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

func TestMysql(t *testing.T) {
	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = "12345"
	config.Net = "tcp"
	config.Addr = "192.168.2.127"
	config.ReadTimeout = time.Second * 5
	config.WriteTimeout = time.Second * 5
	config.Timeout = time.Second * 5
	//mysqldriver := &mysql.MySQLDriver{}
	dnsstr := config.FormatDSN()
	//mysqldriver.Open(dnsstr)

	db, err := sql.Open("mysql", dnsstr)
	if err != nil {
		fmt.Println("Error on conntecting : ", err)
		t.FailNow()
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Error on ping : ", err)
		t.FailNow()
	}
}
