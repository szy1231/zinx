package net

import "zinx/iface"

var _ iface.IRequest = new(Request)

type Request struct {
	//已经和客户端建立好的链接
	conn iface.IConnection

	//客服端请求的数据
	data []byte
}

//得到当前链接
func (r *Request) GetConnection() iface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}

