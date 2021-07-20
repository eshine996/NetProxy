package bridge

import (
	"NetProxy/lib/connect"
)

// Terminal 终端类
type Terminal struct {
	id        int              //在数据库中的id
	Conn      *connect.Connect //与终端设备的来连接
	retryTime int              //重试次数，当我们在写消息时3此不成功则关闭
	Vkey      string           //验证码
	Host      string           //连接终端的地址
}
