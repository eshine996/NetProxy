package main

import (
	"NetProxy/bridge"
)

func main() {
	//db.Sql.AutoMigrate(&db.BridgeInfo{})
	b := bridge.Bridge{
		Port: 8012,
	}
	go b.Run() //开启网桥服务器

	select {}
}
