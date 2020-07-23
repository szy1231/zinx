package net

import (
	"fmt"
	"net"
	"zinx/iface"
)

/*
	链接模块
*/

type Connection struct {
	//当前链接的socket TCP套接字
	Conn *net.TCPConn

	//链接的ID
	ConnId uint32

	//当前链接的状态
	isClosed bool

	//当前链接所绑定的处理业务方法的API
	handleAPI iface.HandleFunc

	//告知当前链接已经退出的/停止 channel
	ExitChan chan bool
}

//初始化链接模块的方法
func NewConnection(conn *net.TCPConn,ConnId uint32, callbackApi iface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnId:    ConnId,
		isClosed:  false,
		handleAPI: callbackApi,
		ExitChan:  make(chan bool,1),
	}
	return c
}

var _ iface.IConnection = new(Connection)

func (c *Connection) StartReader()  {
	fmt.Println("reader go is running ")
	defer c.Stop()
	defer fmt.Printf("connID=%d,Reader is exit,remote addr is %s\n",c.ConnId,c.GetRemoteAddr().String())

	for {
		//读取客户端的数据，最大512
		buf := make([]byte,512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err:",err)
			continue
		}

		//调用当前链接所绑定的HandleAPI
		if err := c.handleAPI(c.Conn,buf,cnt);err != nil {
			fmt.Printf("ConnID=%d,handle error:%s\n",c.ConnId,err.Error())
			break
		}
	}
}
func (c *Connection) Start() {
	fmt.Println("conn start()... connId = ",c.ConnId)

	//启动从当前链接的读数据业务
	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()... ConnID = ",c.ConnId)

	//如果当前链接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//关闭链接
	c.Conn.Close()
	//关闭管道，回收资源
	close(c.ExitChan)
}

//获取当前链接的绑定socket conn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnId
}

func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	panic("implement me")
}