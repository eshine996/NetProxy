package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//拨号,建立tcp连接
	conn, err := net.Dial("tcp", "127.0.0.1:8012")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//开启一个读消息协程
	go func() {
		for {
			msg := make([]byte, 5)
			_, err := conn.Read(msg)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(string(msg))
		}
	}()

	//主线程循环写
	for {
		_, err := conn.Write([]byte("hello"))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		time.Sleep(time.Second * 3)
	}
}
