package main

import (
	"product-mall/conf"
	"product-mall/internal/routes"
	"product-mall/pkg/db"
)

func main() {
	conf.Init()
	db.Database(conf.MysqlpathRead, conf.MysqlpathWrite)
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
