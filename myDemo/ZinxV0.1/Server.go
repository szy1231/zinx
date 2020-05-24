package main

import "zinx/net"

/*
	基于zinx框架来开发的，服务器端应用程序
*/

func main() {
	//1.创建一个server句柄
	s := net.NewServer("zinx0.1")
	//启动server
	s.Serve()
}
