package server

import (
	"log"
	"os"
	"time"
)

func StartLog() {
	logname := time.Now().Format("20060102150405")
	f, err := os.OpenFile("./log/"+logname+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	nowTime := time.Now().AddDate(0, 0, 1)
	timeStr := nowTime.Format("2006-01-02") + " 03:00"
	t, _ := time.ParseInLocation("2006-01-01 01:01:01", timeStr, time.Local)
	a := t.Sub(time.Now())
	time.AfterFunc(a, func() {
		f.Close()
		logname := time.Now().Format("20060102150405")
		f, err = os.OpenFile("./log/"+logname+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("error opening file: %v", err)
		}
		defer f.Close()
	})
}
