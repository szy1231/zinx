package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/iface"
)

/*
	链接模块
*/
var _ iface.IConnection = new(Connection)

type Connection struct {
	//当前链接的socket TCP套接字
	Conn *net.TCPConn

	//链接的ID
	ConnId uint32

	//当前链接的状态
	isClosed bool

	//当前链接所绑定的处理业务方法的API
	//handleAPI iface.HandleFunc

	//告知当前链接已经退出的/停止 channel
	ExitChan chan bool

	//该链接处理的方法Router
	Router iface.IRouter
}

//初始化链接模块的方法
func NewConnection(conn *net.TCPConn, ConnId uint32, router iface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnId:   ConnId,
		isClosed: false,
		//handleAPI: callbackApi,
		ExitChan: make(chan bool, 1),
		Router:   router,
	}
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("reader go is running ")
	defer c.Stop()
	defer fmt.Printf("connID=%d,Reader is exit,remote addr is %s\n", c.ConnId, c.GetRemoteAddr().String())

	for {
		//读取客户端的数据
		/*buf := make([]byte, util.GlobalObject.MaxPackageSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read buf err:", err)
			continue
		}*/

		//调用当前链接所绑定的HandleAPI
		/*if err := c.handleAPI(c.Conn,buf,cnt);err != nil {
			fmt.Printf("ConnID=%d,handle error:%s\n",c.ConnId,err.Error())
			break
		}*/

		//创建一个拆包解包对象
		dp := NewDataPack()
		
		headData := make([]byte,dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData);err!= nil{
			fmt.Println("read msg head error:",err)
			break
		}

		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack error:",err)
			break
		}

		var data []byte
		if msg.GetMsgLen() >0 {
			data := make([]byte,msg.GetMsgLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data);err!= nil{
				fmt.Println("read msg data error:",err)
				break
			}
		}

		msg.SetData(data)
		//得到conn数据的request请求数据
		req := Request{
			conn: c,
			msg: msg,
		}

		//执行注册的路由方法
		go func(request iface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)

	}
}

func (c *Connection) Start() {
	fmt.Println("conn start()... connId = ", c.ConnId)

	//启动从当前链接的读数据业务
	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop()... ConnID = ", c.ConnId)

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

//sendMsg,先封包，再发送
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed {
		return errors.New(" The connection is closed")
	}

	dp := NewDataPack()
	binaryMsg, err := dp.Pack(NewMsg(msgId, data))
	if err != nil {
		fmt.Println("pack error:",err)
		return errors.New("pack error")
	}

	//将数据发送给客户端
	if _, err = c.Conn.Write(binaryMsg);err != nil{
		fmt.Println("conn write error",err)
		return errors.New("conn write error")
	}

	return nil
}