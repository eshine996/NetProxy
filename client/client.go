package main

import (
	"NetProxy/lib/connect"
	"fmt"
	"io"
	"net"
)

func main() {

	//拨号,建立tcp连接
	conn, err := net.Dial("tcp", "127.0.0.1:8012")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	c := connect.Connect{
		Conn: conn,
	}

	//开启一个读消息协程
	go func() {
		for {
			msg, err := c.ReadMessage()
			if err != nil {
				if err != io.EOF {
					fmt.Println(err.Error())
					return
				}
				fmt.Println(err.Error())
			}
			fmt.Printf("%s:%s\n", msg.MsgType, msg.Content)
		}
	}()

	if _, err = c.SendVkey("123"); err != nil {
		fmt.Println(err.Error())
	}
	select {}

	////主线程循环写
	//for {
	//	_, err :=
	//	if err != nil {
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	time.Sleep(time.Second * 3)
	//}
}
