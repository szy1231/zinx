package net

import (
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
}
