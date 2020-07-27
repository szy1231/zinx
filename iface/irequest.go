package iface

/*
	把客户端请求的链接信息 和 请求的数据封装
*/

type IRequest interface {
	//得到当前链接
	GetConnection() IConnection

	//得到请求的消息数据
	GetData() []byte

	GetMsgID() uint32
}
