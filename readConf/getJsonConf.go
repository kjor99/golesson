package readConf

import (
	"fmt"
	"os"
)

// v JSON 结构体
func GetJsonConf(url string) *os.File {

	file, err := os.Open(url)
	//defer file.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil
	}

	return file
}
