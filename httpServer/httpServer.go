package httpServer

import (
	"encoding/json"
	"fmt"
	"github.com/go-ini/ini"
	"gohttpconfig/iniConfig"
	"gohttpconfig/msgJson"
	"io/ioutil"
	"net/http"
	"reflect"
)

var Config *ini.File
var ConfigPath string
var IsConfigExist = false

//dst 要修改的结构体 src 有数据的结构体
func structAssign(dst interface{}, src interface{}) {
	srcElem := reflect.ValueOf(src).Elem() //获取reflect.Type类型
	dstElem := reflect.ValueOf(dst).Elem() //获取reflect.Type类型

	srcType := srcElem.Type()

	//在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
	for i := 0; i < srcElem.NumField(); i++ {
		name := srcType.Field(i).Name
		if ok := dstElem.FieldByName(name).IsValid(); ok {
			dstElem.FieldByName(name).Set(reflect.ValueOf(srcElem.Field(i).Interface()))
		}
	}

}

func Run(port int16, configPath string) {
	ConfigPath = configPath
	var err error
	Config, err = ini.Load(configPath)
	if err != nil {
		fmt.Printf("cant not load ini file:%s\n", configPath)
		IsConfigExist = false
	} else {
		IsConfigExist = true
	}

	http.Handle("/", http.FileServer(http.Dir("html")))
	http.HandleFunc("/SendConfig_base", sendConfig_base)
	http.HandleFunc("/SendConfig_distance", sendConfig_distance)
	http.HandleFunc("/SendConfig_vibrate_setting", sendConfig_vibrate_setting)
	http.HandleFunc("/SendConfig_crossing_setting", sendConfig_crossing_setting)
	http.HandleFunc("/SendConfig_real_loc", sendConfig_real_loc)
	http.HandleFunc("/SendConfig_pixel_loc", sendConfig_pixel_loc)

	defer http.ListenAndServe("localhost:8080", nil)
}

func sendConfig_pixel_loc(writer http.ResponseWriter, request *http.Request) {
	//1.解析http请求
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return
	}
	fmt.Printf("body:%s\n", body)

	//2. 配置文件不存在相应失败 退出
	if !IsConfigExist {
		fmt.Printf("config file not exist\n")
		writer.Write([]byte("失败：配置文件不存在"))
		return
	}

	//3.将请求主体转化为json结构体，然后将json结构体转化为ini结构体，注意，两个结构体变量名称保持一致
	jsonPixel_loc := msgJson.Pixel_loc{}
	err = json.Unmarshal(body, &jsonPixel_loc)
	if err != nil {
		fmt.Printf("json unmarshal err:%v\n", err.Error())
		writer.Write([]byte("失败：json解析失败"))
		return
	}
	iniPixel_loc := iniConfig.Pixel_loc{}
	structAssign(&iniPixel_loc, &jsonPixel_loc)
	//4 结构体写入ini分区
	section, _ := Config.NewSection("pixel_loc")
	err_section := section.ReflectFrom(&iniPixel_loc)
	if err_section != nil {
		fmt.Printf("ini map jsonBase fail:%v\n", err_section.Error())
		writer.Write([]byte("失败:ini写入分区失败"))
		return
	}
	err_save := Config.SaveTo(ConfigPath)
	if err_save != nil {
		fmt.Printf("ini config save fail:%v\n", err_save.Error())
		writer.Write([]byte("失败：配置文件分区信息写入失败"))
		return
	}
	writer.Write([]byte("成功：写入分区信息成功"))
}

func sendConfig_real_loc(writer http.ResponseWriter, request *http.Request) {
	//1.解析http请求
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return
	}
	fmt.Printf("body:%s\n", body)

	//2. 配置文件不存在相应失败 退出
	if !IsConfigExist {
		fmt.Printf("config file not exist\n")
		writer.Write([]byte("失败：配置文件不存在"))
		return
	}

	//3.将请求主体转化为json结构体，然后将json结构体转化为ini结构体，注意，两个结构体变量名称保持一致
	jsonReal_loc := msgJson.Real_loc{}
	err = json.Unmarshal(body, &jsonReal_loc)
	if err != nil {
		fmt.Printf("json unmarshal err:%v\n", err.Error())
		writer.Write([]byte("失败：json解析失败"))
		return
	}
	iniReal_loc := iniConfig.Real_loc{}
	structAssign(&iniReal_loc, &jsonReal_loc)
	//4 结构体写入ini分区
	section, _ := Config.NewSection("real_loc")
	err_section := section.ReflectFrom(&iniReal_loc)
	if err_section != nil {
		fmt.Printf("ini map jsonBase fail:%v\n", err_section.Error())
		writer.Write([]byte("失败:ini写入分区失败"))
		return
	}
	err_save := Config.SaveTo(ConfigPath)
	if err_save != nil {
		fmt.Printf("ini config save fail:%v\n", err_save.Error())
		writer.Write([]byte("失败：配置文件分区信息写入失败"))
		return
	}
	writer.Write([]byte("成功：写入分区信息成功"))
}

func sendConfig_crossing_setting(writer http.ResponseWriter, request *http.Request) {
	//1.解析http请求
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return
	}
	fmt.Printf("body:%s\n", body)

	//2. 配置文件不存在相应失败 退出
	if !IsConfigExist {
		fmt.Printf("config file not exist\n")
		writer.Write([]byte("失败：配置文件不存在"))
		return
	}

	//3.将请求主体转化为json结构体，然后将json结构体转化为ini结构体，注意，两个结构体变量名称保持一致
	jsonCrossing_setting := msgJson.Crossing_setting{}
	err = json.Unmarshal(body, &jsonCrossing_setting)
	if err != nil {
		fmt.Printf("json unmarshal err:%v\n", err.Error())
		writer.Write([]byte("失败：json解析失败"))
		return
	}
	iniCrossing_setting := iniConfig.Crossing_setting{}
	structAssign(&iniCrossing_setting, &jsonCrossing_setting)
	//4 结构体写入ini分区
	section, _ := Config.NewSection("crossing_setting")
	err_section := section.ReflectFrom(&iniCrossing_setting)
	if err_section != nil {
		fmt.Printf("ini map jsonBase fail:%v\n", err_section.Error())
		writer.Write([]byte("失败:ini写入分区失败"))
		return
	}
	err_save := Config.SaveTo(ConfigPath)
	if err_save != nil {
		fmt.Printf("ini config save fail:%v\n", err_save.Error())
		writer.Write([]byte("失败：配置文件分区信息写入失败"))
		return
	}
	writer.Write([]byte("成功：写入分区信息成功"))
}

func sendConfig_vibrate_setting(writer http.ResponseWriter, request *http.Request) {
	//1.解析http请求
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return
	}
	fmt.Printf("body:%s\n", body)

	//2. 配置文件不存在相应失败 退出
	if !IsConfigExist {
		fmt.Printf("config file not exist\n")
		writer.Write([]byte("失败：配置文件不存在"))
		return
	}

	//3.将请求主体转化为json结构体，然后将json结构体转化为ini结构体，注意，两个结构体变量名称保持一致
	jsonVibrate_setting := msgJson.Vibrate_setting{}
	err = json.Unmarshal(body, &jsonVibrate_setting)
	if err != nil {
		fmt.Printf("json unmarshal err:%v\n", err.Error())
		writer.Write([]byte("失败：json解析失败"))
		return
	}
	iniVibrate_setting := iniConfig.Vibrate_setting{}
	structAssign(&iniVibrate_setting, &jsonVibrate_setting)
	//4 结构体写入ini分区
	section, _ := Config.NewSection("vibrate_setting")
	err_section := section.ReflectFrom(&iniVibrate_setting)
	if err_section != nil {
		fmt.Printf("ini map jsonBase fail:%v\n", err_section.Error())
		writer.Write([]byte("失败:ini写入分区失败"))
		return
	}
	err_save := Config.SaveTo(ConfigPath)
	if err_save != nil {
		fmt.Printf("ini config save fail:%v\n", err_save.Error())
		writer.Write([]byte("失败：配置文件分区信息写入失败"))
		return
	}
	writer.Write([]byte("成功：写入分区信息成功"))
}

func sendConfig_distance(writer http.ResponseWriter, request *http.Request) {
	//1.解析http请求
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return
	}
	fmt.Printf("body:%s\n", body)

	//2. 配置文件不存在相应失败 退出
	if !IsConfigExist {
		fmt.Printf("config file not exist\n")
		writer.Write([]byte("失败：配置文件不存在"))
		return
	}

	//3.将请求主体转化为json结构体，然后将json结构体转化为ini结构体，注意，两个结构体变量名称保持一致
	jsonDistance := msgJson.Distance{}
	err = json.Unmarshal(body, &jsonDistance)
	if err != nil {
		fmt.Printf("json unmarshal err:%v\n", err.Error())
		writer.Write([]byte("失败：json解析失败"))
		return
	}
	iniDistance := iniConfig.Distance{}
	structAssign(&iniDistance, &jsonDistance)
	//4 结构体写入ini分区
	section, _ := Config.NewSection("distance")
	err_section := section.ReflectFrom(&iniDistance)
	if err_section != nil {
		fmt.Printf("ini map jsonBase fail:%v\n", err_section.Error())
		writer.Write([]byte("失败:ini写入分区失败"))
		return
	}
	err_save := Config.SaveTo(ConfigPath)
	if err_save != nil {
		fmt.Printf("ini config save fail:%v\n", err_save.Error())
		writer.Write([]byte("失败：配置文件分区信息写入失败"))
		return
	}
	writer.Write([]byte("成功：写入分区信息成功"))
}

func sendConfig_base(writer http.ResponseWriter, request *http.Request) {

	//1.解析http请求
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return
	}
	fmt.Printf("body:%s\n", body)

	//2. 配置文件不存在相应失败 退出
	if !IsConfigExist {
		fmt.Printf("config file not exist\n")
		writer.Write([]byte("失败：配置文件不存在"))
		return
	}

	//3.将请求主体转化为json结构体，然后将json结构体转化为ini结构体，注意，两个结构体变量名称保持一致
	jsonBase := msgJson.Base{}
	err = json.Unmarshal(body, &jsonBase)
	if err != nil {
		fmt.Printf("json unmarshal err:%v\n", err.Error())
		writer.Write([]byte("失败：json解析失败"))
		return
	}
	iniBase := iniConfig.Base{}
	structAssign(&iniBase, &jsonBase)
	//4 结构体写入ini分区
	section, _ := Config.NewSection("base")
	err_section := section.ReflectFrom(&iniBase)
	if err_section != nil {
		fmt.Printf("ini map jsonBase fail:%v\n", err_section.Error())
		writer.Write([]byte("失败:ini写入分区失败"))
		return
	}
	err_save := Config.SaveTo(ConfigPath)
	if err_save != nil {
		fmt.Printf("ini config save fail:%v\n", err_save.Error())
		writer.Write([]byte("失败：配置文件分区信息写入失败"))
		return
	}
	writer.Write([]byte("成功：写入分区信息成功"))
}
