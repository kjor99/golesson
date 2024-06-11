package utils

import (
	"crypto/md5"
	"fmt"
)

func ToMd5(str string) string {

	b := []byte(str)
	str = fmt.Sprintf("%x", md5.Sum(b))
	return str
}
