package model

import (
	"fmt"
	"os"
)

func migration() {
	//自动迁移模式
	err := DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&User{},
			&Notice{},
			&Product{},
			&ProductImg{})
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
}
