package connect

import (
	"NetProxy/lib/crypt"
	"NetProxy/lib/myerr"
	"NetProxy/lib/tools"
	"net"
	"time"
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

//ReadMessage 根据我们的协议规则 读取并解析消息
func (c *Connect) ReadMessage() (*Message, error) {

	var err error
	var msgType []byte

	//设置读取超时，不能阻塞
	if err = c.SetReadDeadline(time.Now().Add(5 * time.Second)); err != nil {
		return nil, myerr.New(err.Error())
	}

	if msgType, err = c.readByLength(4); err != nil {
		return nil, myerr.New(err.Error())
	}

	var bLength []byte
	if bLength, err = c.readByLength(4); err != nil {
		return nil, myerr.New(err.Error())
	}

	var length int
	if length, err = tools.Bytes2Int(bLength); err != nil {
		return nil, myerr.New(err.Error())
	}

	var content []byte
	if content, err = c.readByLength(length); err != nil {
		return nil, myerr.New(err.Error())
	}

	msg := NewMessage(string(msgType), string(content))
	return &msg, nil

}

//SendError 发送错误消息
func (c *Connect) SendError(content string) (int, error) {
	msg := NewMessage(MSGERROR, content)
	return c.Write(msg.ToBytes())
}

//SendVkey 发送验证码
func (c *Connect) SendVkey(vkeyStr string) (int, error) {
	md5vkey := crypt.MD5(vkeyStr)
	msg := NewMessage(MSGVKEY, md5vkey)
	return c.Write(msg.ToBytes())
}

//SendHeartBeat 发送心跳
func (c *Connect) SendHeartBeat(msgStr string) (int, error) {
	msg := NewMessage(MSGHEART, msgStr)
	return c.Write(msg.ToBytes())
}

//SendConnInfo 发送连接信息
func (c *Connect) SendConnInfo(msgStr string) (int, error) {
	msg := NewMessage(MSGHEART, msgStr)
	return c.Write(msg.ToBytes())
}
