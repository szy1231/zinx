package iface

/*
	用于处理tcp粘包。封包、拆包
*/

type IdataPack interface {
	//获取包的头长度
	GetHeadLen() uint32
	//封包
	Pack(msg IMessage)([]byte,error)
	//拆包
	Unpack([]byte)(IMessage,error)
}
