package main

import (
	"flag"
	"gohttpconfig/db"
	"gohttpconfig/serverHttp"
	"os"
)

/**
利用go http server 来配置ini文件

*/

func main() {
	//test
	//serverHttp.StructAssignTest()
	//db.DbTest()

	//默认的端口和配置文件地址
	var port int
	var path string
	var dbPath string
	//读取传入的参数
	flag.IntVar(&port, "port", 8080, "http 端口号 默认 8080")
	flag.StringVar(&path, "path", "./config/distanceN1.ini", "配置文件路径，默认 ./config/distanceN1.ini")
	flag.StringVar(&dbPath, "dbPath", "./config/config.db", "配置文件路径，默认 ./config/config.db")

	flag.Parse()

	//如果文件不存在则创建
	file, err := os.Open(path)
	if err != nil && os.IsNotExist(err) {
		file, _ = os.Create(path)
	}
	file.Close()
	//打开数据库
	db.Open(dbPath)
	serverHttp.Run(port, path)

}
