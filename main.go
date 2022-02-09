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
	var htmlPath string
	var pathCommunicate string
	var staticFilesPath string
	var updateFilePath string
	var version string = "1.0.2"
	//读取传入的参数
	flag.IntVar(&port, "port", 8080, "http 端口号 默认 8080")
	flag.StringVar(&path, "path", "./config/distanceN1.ini", "配置文件路径，默认 ./config/distanceN1.ini")
	flag.StringVar(&dbPath, "dbPath", "./config/config.db", "配置文件路径，默认 ./config/config.db")
	flag.IntVar(&configWay, "configWay", 1, "配置存放方式(默认为1)： 0：ini文件，1：sqlite")
	flag.StringVar(&htmlPath, "htmlPath", "./html", "html文件夹路径，默认 ./html")
	flag.StringVar(&pathCommunicate, "pathCommunicate", "./config/communicate.ini",
		"摄像头配置文件文件夹路径，默认 ./config/communicate.ini")
	flag.StringVar(&staticFilesPath, "staticFilesPath", "./image", "可下载文件路径，默认 ./image")
	flag.StringVar(&updateFilePath, "updateFilePath", "./update", "可下载文件路径，默认 ./update")
	if len(os.Args) == 2 {
		if os.Args[1] == "-v" {
			fmt.Println("version:", version)
			os.Exit(1)
		}

	}

	flag.Parse()

	switch configWay {
	case 0:
		serverHttp.ConfigType = serverHttp.ConfigIni
	case 1:
		serverHttp.ConfigType = serverHttp.ConfigSqlite

	}

	//distanceN1
	serverHttp.ConfigPath = path
	//communicate
	serverHttp.ConfigPathCommuniate = pathCommunicate
	//可下载文件路径
	serverHttp.StaticFilePath = staticFilesPath
	//上传更新文件的存储路径
	serverHttp.UpdateFilePath = updateFilePath

	//如果文件不存在则创建
	file, err := os.Open(path)
	if err != nil && os.IsNotExist(err) {
		file, _ = os.Create(path)
	}
	file.Close()
	//摄像头配置文件
	fileCommunicate, errCommunicate := os.Open(pathCommunicate)
	if errCommunicate != nil && os.IsNotExist(err) {
		file, _ = os.Create(path)
	}
	fileCommunicate.Close()
	//打开数据库
	db.Open(dbPath)
	serverHttp.Run(port, htmlPath)
}
