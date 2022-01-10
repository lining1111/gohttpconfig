package main

import (
	"flag"
	"fmt"
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

	//s := "-41.5"
	//value, err := strconv.ParseFloat(s, 64)
	//fmt.Printf("value:%v,err:%v\n", value, err)

	//默认的端口和配置文件地址
	var port int
	var path string
	var dbPath string
	var configWay int
	var version string = "1.0.0"
	//读取传入的参数
	flag.IntVar(&port, "port", 8080, "http 端口号 默认 8080")
	flag.StringVar(&path, "path", "./config/distanceN1.ini", "配置文件路径，默认 ./config/distanceN1.ini")
	flag.StringVar(&dbPath, "dbPath", "./config/config.db", "配置文件路径，默认 ./config/config.db")
	flag.IntVar(&configWay, "configWay", 1, "配置存放方式(默认为1)： 0：ini文件，1：sqlite")
	if len(os.Args) == 2 {
		if os.Args[1] == "-v" {
			fmt.Println("version:", version)
			os.Exit(1)
		}

	}

	switch configWay {
	case 0:
		serverHttp.ConfigType = serverHttp.ConfigIni
	case 1:
		serverHttp.ConfigType = serverHttp.ConfigSqlite

	}

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
