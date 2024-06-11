package main

import (
	"github.com/kjor99/golesson/dao"
	"github.com/kjor99/golesson/router"
)

func init() {
	dao.Conn()
}

func main() {

	router.StartGin()

}
