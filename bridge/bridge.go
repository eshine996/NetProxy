package bridge

import (
	"NetProxy/db"
	"NetProxy/lib/connect"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

//Bridge 基于自定义的库，开始改造bridge
type Bridge struct {
	Port      int
	Terminals sync.Map //key=uuid value=Conn 多个终端设备接入,用来储存连接，sync.map保证线程安全
}

//Run 启动网桥
func (b *Bridge) Run() {
	//开启ping
	go b.ping()

	err := connect.TcpListenerAndProcess("0.0.0.0:"+strconv.Itoa(b.Port), b.process)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

}

func (b *Bridge) process(c connect.Connect) {
	message, err := c.ReadMessage()
	if err != nil {
		fmt.Println(err.Error())
		_, _ = c.SendError(err.Error())
		_ = c.Close()
		return
	}

	var id string

	switch message.MsgType {
	case connect.MSGVKEY:

		//验证vkey是否正确
		var ret bool
		bridgeInfo := db.BridgeInfo{Vkey: message.Content}

		if ret, err = bridgeInfo.IsExist(); err != nil {
			fmt.Println(err.Error())
			_, _ = c.SendError(err.Error())
			_ = c.Close()
			return
		}

		if !ret {
			_, _ = c.SendError("vkey is not register")
			_ = c.Close()
			return
		}

		id = message.Content

	default:
		_, _ = c.SendError("Bad protocol message")
		_ = c.Close()
		return
	}

	t := Terminal{
		Conn: &c,
		Vkey: id,
	}

	b.Terminals.Store(id, &t)

}

func (b *Bridge) ping() {
	ticker := time.NewTicker(5 * time.Second) //5秒定时器
	defer ticker.Stop()
	for {
		select {
		case ti := <-ticker.C:

			b.Terminals.Range(func(key, value interface{}) bool {
				t := value.(*Terminal)

				go func() {
					if t.retryTime >= 3 {
						b.Terminals.Delete(key.(string)) //发送失败3次则关闭连接
					}

					timeStr := ti.Format("2006-01-02 15:04:05")
					if _, err := t.Conn.SendHeartBeat(timeStr); err != nil {
						fmt.Println(err.Error())
						t.retryTime++
					}
				}()
				return true
			})
		}
	}
}
