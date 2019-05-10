package server

import (
	"fmt"
	"gitlab.com/adoontheway/goserver/util"
	"log"
	"os"
	"time"
)

var f os.File

func StartLog() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	result,err := util.CheckPathExist("./log")

	if err != nil {
		fmt.Println("Check Directory Failed:",err)
	}else{
		if !result {
			err = os.Mkdir("./log",os.ModePerm)
			if err != nil {
				fmt.Println("Make Directory failed:",err)
			}
		}
	}

	logname := time.Now().Format("20060102-150405")

	f, err := os.OpenFile("./log/"+logname+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)
	log.Println("Log file started")
	fmt.Print("Log file started")
	//daily backup
}

func CloseLog() {
	f.Close()
}
