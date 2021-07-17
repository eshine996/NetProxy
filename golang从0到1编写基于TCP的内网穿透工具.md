# golang从0到1编写基于TCP的内网穿透工具

1.概述。

2.搭建TCP网桥服务器。

网桥服务器的作用是保证云端与终端长连接通讯。

bridge.go

```go
var tcpAddr *net.TCPAddr
	var err error

	if tcpAddr, err = net.ResolveTCPAddr("tcp", addr); err != nil {
		return nil, myerr.New(err.Error())
	}

	var tcpListener *net.TCPListener
	if tcpListener, err = net.ListenTCP("tcp", tcpAddr); err != nil {
		return nil, myerr.New(err.Error())
	}
```

