package main

import (
	"gin-tool-study/conf"
	"gin-tool-study/routes"
)

func main() {
	// Ek1+Ep1==Ek2+Ep2
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
