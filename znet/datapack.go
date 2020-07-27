package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/iface"
	"zinx/util"
)

var _ iface.IdataPack = new(DataPack)

type DataPack struct {

}

//初始化方法
func NewDataPack() *DataPack {
	return &DataPack{}
}

func (d *DataPack) GetHeadLen() uint32 {
	//长度（4）+ 类型（4）
	return 8
}

func (d *DataPack) Pack(msg iface.IMessage) ([]byte, error) {
	dataBuff := bytes.NewBuffer([]byte{})

	//写入len
	err := binary.Write(dataBuff,binary.LittleEndian,msg.GetMsgLen())
	if err != nil {
		return nil, err
	}
	//写入类型 msgId
	err = binary.Write(dataBuff,binary.LittleEndian,msg.GetMsgId())
	if err != nil {
		return nil, err
	}
	//写入数据
	err = binary.Write(dataBuff,binary.LittleEndian,msg.GetData())
	if err != nil {
		return nil, err
	}

	return dataBuff.Bytes(),nil
}

func (d *DataPack) Unpack(byteData []byte) (iface.IMessage, error) {
	//先读取head信息，根据head信息读取内容
	dataBuff := bytes.NewReader(byteData)

	msg := &Message{}

	err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen)
	if err != nil {
		return nil, err
	}
	//长度超出了允许的最大长度
	if msg.DataLen > util.GlobalObject.MaxPackageSize {
		return nil,errors.New(" The data is too large")
	}
	
	err = binary.Read(dataBuff, binary.LittleEndian, &msg.Id)
	if err != nil {
		return nil, err
	}

	return msg,nil
}


