package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"zinx/iface"
)

type GlobalObj struct {

	Server iface.IServer
	Host string
	Port int
	Name string

	MaxConn int  //服务器主机允许的最大链接数
	MaxPackageSize uint32	//当前框架数据包最大值
}


var GlobalObject *GlobalObj

func (g *GlobalObj) Reload()  {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
func init()  {
	//没有配置文件设置的默认值
	GlobalObject = &GlobalObj{
		Name: "ZinxServer",
		Port: 8889,
		Host: "0.0.0.0",
		MaxConn: 1024,
		MaxPackageSize: 2048,
	}

	//应该尝试从conf中加载用户自定义的conf
	GlobalObject.Reload()
	fmt.Println(GlobalObject)
}
