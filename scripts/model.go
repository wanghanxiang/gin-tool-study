package main

import (
	"fmt"

	"github.com/gohouse/converter"
)

func main() {
	err := converter.NewTable2Struct().
		SavePath("./model.go").
		Dsn("root:xiangzai@tcp(127.0.0.1:3306)/product_go?charset=utf8").
		TagKey("gorm").
		EnableJsonTag(true).
		//Table("product").
		Run()
	fmt.Println(err)
}
