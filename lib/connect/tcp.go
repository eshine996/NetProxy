package connect

import (
	"NetProxy/lib/myerr"
	"fmt"
	"net"
	"strings"
)

//封装一些tcp连接的方法

func TcpListenerAndProcess(addr string, f func(c Connect)) error {
	tcpListener, err := NewTCPListener(addr)
	if err != nil {
		return myerr.New(err.Error())
	}

	Accept(tcpListener, f)
	return nil
}

//NewTCPListener 获取一个TCP的监听
func NewTCPListener(addr string) (*net.TCPListener, error) {
	var tcpAddr *net.TCPAddr
	var err error

	if tcpAddr, err = net.ResolveTCPAddr("tcp", addr); err != nil {
		return nil, myerr.New(err.Error())
	}

	var tcpListener *net.TCPListener
	if tcpListener, err = net.ListenTCP("tcp", tcpAddr); err != nil {
		return nil, myerr.New(err.Error())
	}

	return tcpListener, nil
}

//Accept 接收消息然后处理
//f 为函数消息的处理函数
func Accept(l net.Listener, f func(c Connect)) {
	for {
		c, err := l.Accept()

		if err != nil {
			if strings.Contains(err.Error(), "use of closed network connection") {
				break
			}

			if strings.Contains(err.Error(), "the mux has closed") {
				break
			}

			fmt.Println(err.Error())
			continue
		}
		fmt.Println(c.RemoteAddr().String(), "connect in")

		if c == nil {
			//logs.Warn("nil connection")
			break
		}
		go f(Connect{c})
	}
}
