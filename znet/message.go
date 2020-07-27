package znet

import "zinx/iface"

var _ iface.IMessage = new(Message)

type Message struct {
	Id uint32 //消息id
	DataLen uint32 //消息长度
	Data []byte	//内容
}

func NewMsg(id uint32,data []byte) *Message {
	return &Message{
		Id: id,
		DataLen: uint32(len(data)),
		Data: data,
	}
}

func (m *Message) GetMsgId() uint32 {
	return m.Id
}

func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetMsgId(id uint32) {
	m.Id = id
}

func (m *Message) SetMsgLen(len uint32) {
	m.DataLen = len
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}