package main

import (
	"fmt"
	"strconv"
)

func main() {
	conf := ReadConfig()
	fmt.Println("学校名称：" + conf.SchoolName)
	switch conf.MangerType {
	case "supwisdom":
		fmt.Println("教务系统：树维教务系统")
		break
	}
	fmt.Println("绑定端口：" + strconv.Itoa(conf.SocketPort))
	StartWebSocket()
}
