package main

import (
	"fmt"
	"net" //使用golang官方的net库
	"time"
)

func main() {
	var tcpAddr *net.TCPAddr
	var err error

	addr := "0.0.0.0:8012"
	if tcpAddr, err = net.ResolveTCPAddr("tcp", addr); err != nil {
		fmt.Println(err.Error())
		return
	}

	var tcpListener *net.TCPListener
	if tcpListener, err = net.ListenTCP("tcp", tcpAddr); err != nil {
		fmt.Println(err.Error())
		return
	}

	//循环接收连接
	for {
		conn, err := tcpListener.Accept()
		if err != nil {
			fmt.Println(err.Error())
		}

		go connHandler(conn) //开启一个协程进行连接处理
	}
}

//connHandler 处理终端发送的消息
func connHandler(conn net.Conn) {
	go ping(conn) //开启一个协程进行心跳
	for {
		msg := make([]byte, 5)
		_, err := conn.Read(msg)

		fmt.Println(string(msg))
		if err != nil {
			_ = conn.Close()
			return
		}
	}
}

//ping 每5秒发送消息，失败则关闭连接
func ping(conn net.Conn) {
	ticker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ticker.C:
			_, err := conn.Write([]byte("hello"))
			if err != nil {
				fmt.Println(err.Error())
				_ = conn.Close()
				return
			}
		}
	}
}
