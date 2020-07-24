package main

import (
	"fmt"
	"zinx/iface"
	"zinx/znet"
)

/*
	基于zinx框架来开发的，服务器端应用程序
*/

//自定义路由
type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) PreHandle(request iface.IRequest) {
	fmt.Println("call PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping\n"))
	if err != nil {
		fmt.Println("call back before ping error:",err)
	}
}

func (p *PingRouter) Handle(request iface.IRequest) {
	fmt.Println("call Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping\n"))
	if err != nil {
		fmt.Println("call back Handle ping error:",err)
	}
}

func (p *PingRouter) PostHandle(request iface.IRequest) {
	fmt.Println("call PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping\n"))
	if err != nil {
		fmt.Println("call back PostHandle after ping error:",err)
	}
}

func main() {
	//1.创建一个server句柄
	s := znet.NewServer("zinx4")
	//添加自定义router
	s.AddRouter(&PingRouter{})
	//启动server
	s.Serve()
}
