package main

import (
	"NetProxy/bridge"
)

func main() {
	b := bridge.Bridge{
		Port: 9024,
	}
	go b.Run() //开启网桥服务器

	select {}
}
