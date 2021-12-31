package httpServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-ini/ini"
	"gohttpconfig/iniConfig"
	"gohttpconfig/msgJson"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

var Config *ini.File
var ConfigPath string
var IsConfigExist = false

const (
	ConfigIni = iota
	ConfigSqlite
)

var ConfigType = ConfigIni

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

func StructAssignTest() {
	type S struct {
		A string
		B int
	}

	type S0 struct {
		D int
	}
	type S1 struct {
		S  S
		C  string
		S0 S0
	}

	sa := S1{
		S: S{
			A: "nihao",
			B: 5,
		},
		C: "aaa",
		S0: S0{
			5,
		},
	}
	sb := S1{}

	structAssign(&sb, &sa)

}

func Run(port int, configPath string) {
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
	//set
	http.HandleFunc("/setConfig_base", setConfig_base)
	http.HandleFunc("/setConfig_distance", setConfig_distance)
	http.HandleFunc("/setConfig_vibrate_setting", setConfig_vibrate_setting)
	http.HandleFunc("/setConfig_crossing_setting", setConfig_crossing_setting)
	http.HandleFunc("/setConfig_real_loc", setConfig_real_loc)
	http.HandleFunc("/setConfig_pixel_loc", setConfig_pixel_loc)
	http.HandleFunc("/setConfig_all", setConfig_all)
	//get
	http.HandleFunc("/getConfig_base", getConfig_base)
	http.HandleFunc("/getConfig_distance", getConfig_distance)
	http.HandleFunc("/getConfig_vibrate_setting", getConfig_vibrate_setting)
	http.HandleFunc("/getConfig_crossing_setting", getConfig_crossing_setting)
	http.HandleFunc("/getConfig_real_loc", getConfig_real_loc)
	http.HandleFunc("/getConfig_pixel_loc", getConfig_pixel_loc)
	http.HandleFunc("/getConfig_all", getConfig_all)

	addr := "localhost:" + strconv.Itoa(port)
	defer http.ListenAndServe(addr, nil)
}

// 基础函数
func getConfigIni(w http.ResponseWriter, r *http.Request, sectionName string) error {
	var jsonP interface{}
	var iniP interface{}

	switch sectionName {
	case ini.DefaultSection:
		jsonP = &msgJson.Info{}
		iniP = &iniConfig.Info{}
	case "base":
		jsonP = &msgJson.Base{}
		iniP = &iniConfig.Base{}
	case "distance":
		jsonP = &msgJson.Distance{}
		iniP = &iniConfig.Distance{}
	case "vibrate_setting":
		jsonP = &msgJson.Vibrate_setting{}
		iniP = &iniConfig.Vibrate_setting{}
	case "crossing_setting":
		jsonP = &msgJson.Crossing_setting{}
		iniP = &iniConfig.Crossing_setting{}
	case "real_loc":
		jsonP = &msgJson.Real_loc{}
		iniP = &iniConfig.Real_loc{}
	case "pixel_loc":
		jsonP = &msgJson.Pixel_loc{}
		iniP = &iniConfig.Pixel_loc{}
	}

	//1.解析http请求
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return err
	}
	fmt.Printf("body:%s\n", rBody)

	//2. 配置文件不存在相应失败 退出
	if !IsConfigExist {
		fmt.Printf("config file not exist\n")
		w.Write([]byte("失败：配置文件不存在"))
		return errors.New("配置文件不存在")
	}

	//3.读取ini文件指定分区信息，转换为json信息
	section, errGetsection := Config.GetSection(sectionName)
	if errGetsection != nil {
		fmt.Printf("获取分区失败:%v\n", errGetsection.Error())
		w.Write([]byte("获取分区失败"))
		return errGetsection
	}
	errSection := section.MapTo(iniP)
	if errSection != nil {
		fmt.Printf("分区信息转换失败：%v\n", errSection.Error())
		w.Write([]byte("分区信息转换失败"))
		return errSection
	}
	if sectionName == ini.DefaultSection {
		structAssign(&(jsonP.(*msgJson.Info).Base), &(iniP.(*iniConfig.Info).Base))
		structAssign(&(jsonP.(*msgJson.Info).Distance), &(iniP.(*iniConfig.Info).Distance))
		structAssign(&(jsonP.(*msgJson.Info).Vibrate_setting), &(iniP.(*iniConfig.Info).Vibrate_setting))
		structAssign(&(jsonP.(*msgJson.Info).Crossing_setting), &(iniP.(*iniConfig.Info).Crossing_setting))
		structAssign(&(jsonP.(*msgJson.Info).Real_loc), &(iniP.(*iniConfig.Info).Real_loc))
		structAssign(&(jsonP.(*msgJson.Info).Pixel_loc), &(iniP.(*iniConfig.Info).Pixel_loc))
	} else {
		structAssign(jsonP, iniP)
	}
	//4 json信息组织回复
	wBody, errBody := json.Marshal(jsonP)
	if errBody != nil {
		fmt.Printf("json unmarshal err:%v\n", err.Error())
		w.Write([]byte("失败：json解析失败"))
		return errBody
	}
	w.WriteHeader(http.StatusOK)
	w.Write(wBody)
	return nil
}

func setConfigIni(w http.ResponseWriter, r *http.Request, sectionName string) error {

	var jsonP interface{}
	var iniP interface{}

	switch sectionName {
	case ini.DefaultSection:
		jsonP = &msgJson.Info{}
		iniP = &iniConfig.Info{}
	case "base":
		jsonP = &msgJson.Base{}
		iniP = &iniConfig.Base{}
	case "distance":
		jsonP = &msgJson.Distance{}
		iniP = &iniConfig.Distance{}
	case "vibrate_setting":
		jsonP = &msgJson.Vibrate_setting{}
		iniP = &iniConfig.Vibrate_setting{}
	case "crossing_setting":
		jsonP = &msgJson.Crossing_setting{}
		iniP = &iniConfig.Crossing_setting{}
	case "real_loc":
		jsonP = &msgJson.Real_loc{}
		iniP = &iniConfig.Real_loc{}
	case "pixel_loc":
		jsonP = &msgJson.Pixel_loc{}
		iniP = &iniConfig.Pixel_loc{}
	}

	//1.解析http请求
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return err
	}
	fmt.Printf("body:%s\n", rBody)

	//2. 配置文件不存在相应失败 退出
	if !IsConfigExist {
		fmt.Printf("config file not exist\n")
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：配置文件不存在"))
		return errors.New("配置文件不存在")
	}

	//3.将请求主体转化为json结构体，然后将json结构体转化为ini结构体，注意，两个结构体变量名称保持一致
	err = json.Unmarshal(rBody, jsonP)
	if err != nil {
		fmt.Printf("json unmarshal err:%v\n", err.Error())
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：json解析失败"))
		return err
	}
	if sectionName == ini.DefaultSection {
		structAssign(&(iniP.(*iniConfig.Info).Base), &(jsonP.(*msgJson.Info).Base))
		structAssign(&(iniP.(*iniConfig.Info).Distance), &(jsonP.(*msgJson.Info).Distance))
		structAssign(&(iniP.(*iniConfig.Info).Vibrate_setting), &(jsonP.(*msgJson.Info).Vibrate_setting))
		structAssign(&(iniP.(*iniConfig.Info).Crossing_setting), &(jsonP.(*msgJson.Info).Crossing_setting))
		structAssign(&(iniP.(*iniConfig.Info).Real_loc), &(jsonP.(*msgJson.Info).Real_loc))
		structAssign(&(iniP.(*iniConfig.Info).Pixel_loc), &(jsonP.(*msgJson.Info).Pixel_loc))
	} else {
		structAssign(iniP, jsonP)
	}
	//4 结构体写入ini分区
	//判断分区名称是否为空
	if len(sectionName) == 0 {
		errSection := Config.ReflectFrom(iniP)
		if errSection != nil {
			fmt.Printf("ini map jsonBase fail:%v\n", errSection.Error())
			w.WriteHeader(http.StatusGone)
			w.Write([]byte("失败:ini写入分区失败"))
			return errSection
		}
	} else {
		section, _ := Config.NewSection(sectionName)
		errSection := section.ReflectFrom(iniP)
		if errSection != nil {
			fmt.Printf("ini map jsonBase fail:%v\n", errSection.Error())
			w.WriteHeader(http.StatusGone)
			w.Write([]byte("失败:ini写入分区失败"))
			return errSection
		}
	}

	errSave := Config.SaveTo(ConfigPath)
	if errSave != nil {
		fmt.Printf("ini config save fail:%v\n", errSave.Error())
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：配置文件分区信息写入失败"))
		return errSave
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("成功：写入分区信息成功"))
	return nil
}

//web api
func getConfig_all(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, ini.DefaultSection)
	case ConfigSqlite:

	}
}

func getConfig_pixel_loc(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "pixel_loc")
	case ConfigSqlite:

	}
}

func getConfig_real_loc(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "real_loc")
	case ConfigSqlite:

	}
}

func getConfig_crossing_setting(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "crossing_setting")
	case ConfigSqlite:

	}
}

func getConfig_vibrate_setting(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "vibrate_setting")
	case ConfigSqlite:

	}
}

func getConfig_distance(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "distance")
	case ConfigSqlite:

	}
}

func getConfig_base(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "base")
	case ConfigSqlite:

	}
}

func setConfig_all(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, ini.DefaultSection)
	case ConfigSqlite:

	}
}

func setConfig_pixel_loc(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "pixel_loc")
	case ConfigSqlite:

	}
}

func setConfig_real_loc(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "real_loc")
	case ConfigSqlite:

	}
}

func setConfig_crossing_setting(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "crossing_setting")
	case ConfigSqlite:

	}
}

func setConfig_vibrate_setting(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "vibrate_setting")
	case ConfigSqlite:

	}
}

func setConfig_distance(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "distance")
	case ConfigSqlite:

	}
}

func setConfig_base(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "base")
	case ConfigSqlite:

	}
}
