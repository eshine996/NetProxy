package connect

import (
	"NetProxy/lib/myerr"
	"net"
)

/*
定义一个连接对象，主要是为conn提供一些自定义 读/写 的方法
*/

type Connect struct {
	net.Conn //继承net.Conn类
}

//通过字节流长度读取消息
func (c *Connect) readByLength(length int) ([]byte, error) {
	var err error
	buf := make([]byte, length)

	if _, err = c.Read(buf); err != nil {
		return nil, myerr.New(err.Error())
	}

	return buf, nil
}
