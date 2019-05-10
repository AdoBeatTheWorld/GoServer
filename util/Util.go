package util

import (
	"fmt"
	"os"
)

func CheckPathExist(path string) (bool, error) {
	_,err := os.Stat(path)
	if err == nil {
		 return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	fmt.Println("Check Directory failed:",err)
	return false, err
}