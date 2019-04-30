package test

import (
	"fmt"
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	ti := time.Now().Format("2006-01-02 15:04:05.9999")
	fmt.Println("time:", ti)
}
