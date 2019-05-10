package test

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"testing"
	"time"
)

type Android struct {
	ConfigId uint32
	UserId uint32
	ServerId uint32
	EnterTime uint32
	LeaveTime uint32
	TakeMinScore uint32
	TakeMaxScore uint32
	Status uint32
	Active uint32
}

func TestMysql(t *testing.T) {
	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = "12345"
	config.Net = "tcp"
	config.Addr = "192.168.2.127"
	config.ReadTimeout = time.Second * 5
	config.WriteTimeout = time.Second * 5
	config.Timeout = time.Second * 5
	config.DBName = "db_account"
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

	rows,err := db.Query("SELECT * FROM android_config limit 0,10")
	if err != nil {
		fmt.Print("Query Error:",err)
		t.FailNow()
	}
	for {
		if rows.Next() {
			android := &Android{}
			err = rows.Scan(android)
			if err != nil {
				fmt.Println("Scan Error:",err)
				break
			}
			fmt.Println(android.UserId)
		}else {
			break
		}
	}

	//result,err := rows.Columns()
	//if err != nil {
	//	fmt.Println("Error Result:",err)
	//	t.FailNow()
	//	return
	//}
	//
	//fmt.Println("Result:")
	//for k,v := range result {
	//	fmt.Print(k," : ",v)
	//}
}
