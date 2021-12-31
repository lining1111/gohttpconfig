package main

import (
	"flag"
	"gohttpconfig/httpServer"
	"os"
)

/**
利用go http server 来配置ini文件

*/

func main() {
	//test
	httpServer.StructAssignTest()

	//默认的端口和配置文件地址
	var port int
	var path string
	//读取传入的参数
	flag.IntVar(&port, "port", 8080, "http 端口号 默认 8080")
	flag.StringVar(&path, "path", "./config/distanceN1.ini", "配置文件路径，默认 ./config/distanceN1.ini")

	flag.Parse()

	//如果文件不存在则创建
	file, err := os.Open(path)
	if err != nil && os.IsNotExist(err) {
		file, _ = os.Create(path)
	}
	file.Close()
	httpServer.Run(port, path)

}
