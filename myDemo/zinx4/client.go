package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Client start")

	time.Sleep(time.Second)

	//直接链接，得到一个conn
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("client start error:", err)
		return
	}

	for {
		//链接写入数据
		_, err = conn.Write([]byte("你好,zinx4"))
		if err != nil {
			fmt.Println("write conn error:", err)
			return
		}

		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error:", err)
			return
		}

		fmt.Printf("server call back:%s\n", string(buf[:cnt]))

		//cpu阻塞
		time.Sleep(time.Second)
	}
}
