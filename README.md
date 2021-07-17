### 第一步：搭建TCP网桥服务器

**目的：**

建立一个TCP长连接一个网桥服务器,所谓的网桥，其实质就是一个TCP服务器和客户端。用来保证传递云端与终端的实时通信。这样的服务器百度有很多例子，在此不多赘述。

[bridgo.go](./bridge.go)  为部署在云端的网桥服务。

[client.go](./client.go) 部署在本地终端设备。



**思考：**

1.如何解决TCP粘包的问题？

2.客户端（终端）连接如何验证其身份？

3.是否需要指定类似HTTP协议一样的消息协议，以便于解析？
