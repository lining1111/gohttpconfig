package serverHttp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-ini/ini"
	"gohttpconfig/common"
	"gohttpconfig/configStruct"
	"gohttpconfig/db"
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

var ConfigType = ConfigSqlite

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
	//1.解析http请求
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return err
	}
	fmt.Printf("body:%s\n", rBody)

	//2.配置文件不存在相应失败 退出
	if !IsConfigExist {
		fmt.Printf("config file not exist\n")
		w.Write([]byte("失败：配置文件不存在"))
		return errors.New("配置文件不存在")
	}

	//3.读取ini文件指定分区信息，转换为json信息
	var msgSrc interface{}
	var msgDst interface{}

	switch sectionName {
	case ini.DefaultSection:
		msgSrc = &configStruct.Info{}
		msgDst = &common.Info{}
	case "base":
		msg = &common.Base{}
	case "distance":
		msg = &common.Distance{}
	case "vibrate_setting":
		msg = &common.Vibrate_setting{}
	case "crossing_setting":
		msg = &common.Crossing_setting{}
	case "real_loc":
		msg = &common.Real_loc{}
	case "pixel_loc":
		msg = &common.Pixel_loc{}
	default:
		fmt.Printf("unknown name:%s\n", sectionName)
		return errors.New("unknown name")
	}

	section, errGetsection := Config.GetSection(sectionName)
	if errGetsection != nil {
		fmt.Printf("获取分区失败:%v\n", errGetsection.Error())
		w.Write([]byte("获取分区失败"))
		return errGetsection
	}
	errSection := section.MapTo(msg)
	if errSection != nil {
		fmt.Printf("分区信息转换失败：%v\n", errSection.Error())
		w.Write([]byte("分区信息转换失败"))
		return errSection
	}
	//4.json信息组织回复

	wBody, errBody := json.Marshal(msg)
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
	//1.解析http请求
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return err
	}
	fmt.Printf("body:%s\n", rBody)

	//2.配置文件不存在相应失败 退出
	if !IsConfigExist {
		fmt.Printf("config file not exist\n")
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：配置文件不存在"))
		return errors.New("配置文件不存在")
	}

	//3.将请求主体转化为json结构体，然后将json结构体转化为ini结构体，注意，两个结构体变量名称保持一致
	var msg interface{}

	switch sectionName {
	case ini.DefaultSection:
		msg = &common.Info{}
	case "base":
		msg = &common.Base{}
	case "distance":
		msg = &common.Distance{}
	case "vibrate_setting":
		msg = &common.Vibrate_setting{}
	case "crossing_setting":
		msg = &common.Crossing_setting{}
	case "real_loc":
		msg = &common.Real_loc{}
	case "pixel_loc":
		msg = &common.Pixel_loc{}
	default:
		fmt.Printf("unknown name:%s\n", sectionName)
		return errors.New("unknown name")
	}
	err = json.Unmarshal(rBody, msg)
	if err != nil {
		fmt.Printf("json unmarshal err:%v\n", err.Error())
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：json解析失败"))
		return err
	}

	//4.结构体写入ini分区
	section, _ := Config.NewSection(sectionName)
	errSection := section.ReflectFrom(msg)
	if errSection != nil {
		fmt.Printf("ini map jsonBase fail:%v\n", errSection.Error())
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败:ini写入分区失败"))
		return errSection
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

func getConfigDb(w http.ResponseWriter, r *http.Request, tableName string) error {
	//1.解析http请求
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return err
	}
	fmt.Printf("body:%s\n", rBody)

	//2.配置文件不存在相应失败 退出
	if !db.IsOpen {
		fmt.Printf("db not open\n")
		w.Write([]byte("失败：数据库未打开"))
		return errors.New("数据库未打开")
	}

	//3.读取数据库指定信息，转换为json信息
	var msg interface{}
	var errDb error
	switch tableName {
	case "":
		msg = &common.Info{}
		errDb = db.GetConfig_all(msg.(*common.Info))
	case "base":
		msg = &common.Base{}
		errDb = db.GetConfig_base(msg.(*common.Base))
	case "distance":
		msg = &common.Distance{}
		errDb = db.GetConfig_distance(msg.(*common.Distance))
	case "vibrate_setting":
		msg = &common.Vibrate_setting{}
		errDb = db.GetConfig_vibrate_setting(msg.(*common.Vibrate_setting))
	case "crossing_setting":
		msg = &common.Crossing_setting{}
		errDb = db.GetConfig_crossing_setting(msg.(*common.Crossing_setting))
	case "real_loc":
		msg = &common.Real_loc{}
		errDb = db.GetConfig_real_loc(msg.(*common.Real_loc))
	case "pixel_loc":
		msg = &common.Pixel_loc{}
		errDb = db.GetConfig_pixel_loc(msg.(*common.Pixel_loc))
	default:
		fmt.Printf("unknown name:%s\n", tableName)
		return errors.New("unknown name")
	}
	if errDb != nil {
		fmt.Printf("db get fail:%v\n", err.Error())
		w.Write([]byte("失败：数据库读取失败"))
		return errDb
	}

	//4.json信息组织回复
	wBody, errBody := json.Marshal(msg)
	if errBody != nil {
		fmt.Printf("json unmarshal err:%v\n", err.Error())
		w.Write([]byte("失败：json解析失败"))
		return errBody
	}
	w.WriteHeader(http.StatusOK)
	w.Write(wBody)
	return nil
}

func setConfigDb(w http.ResponseWriter, r *http.Request, tableName string) error {
	//1.解析http请求
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return err
	}
	fmt.Printf("body:%s\n", rBody)

	//2.配置文件不存在相应失败 退出
	if !db.IsOpen {
		fmt.Printf("db not exist\n")
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：数据库"))
		return errors.New("数据库不存在")
	}

	//3.将请求主体转化为json结构体，然后将json结构体转化为ini结构体，注意，两个结构体变量名称保持一致
	var msg interface{}

	switch tableName {
	case "":
		msg = &common.Info{}
	case "base":
		msg = &common.Base{}
	case "distance":
		msg = &common.Distance{}
	case "vibrate_setting":
		msg = &common.Vibrate_setting{}
	case "crossing_setting":
		msg = &common.Crossing_setting{}
	case "real_loc":
		msg = &common.Real_loc{}
	case "pixel_loc":
		msg = &common.Pixel_loc{}
	default:
		fmt.Printf("unknown name:%s\n", tableName)
		return errors.New("unknown name")
	}
	err = json.Unmarshal(rBody, msg)
	if err != nil {
		fmt.Printf("json unmarshal err:%v\n", err.Error())
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：json解析失败"))
		return err
	}

	//4.结构体写入指定的数据库表
	var errDb error
	switch tableName {
	case "":
		errDb = db.SetConfig_all(msg.(*common.Info))
	case "base":
		errDb = db.SetConfig_base(msg.(*common.Base))
	case "distance":
		errDb = db.SetConfig_distance(msg.(*common.Distance))
	case "vibrate_setting":
		errDb = db.SetConfig_vibrate_setting(msg.(*common.Vibrate_setting))
	case "crossing_setting":
		errDb = db.SetConfig_crossing_setting(msg.(*common.Crossing_setting))
	case "real_loc":
		errDb = db.SetConfig_real_loc(msg.(*common.Real_loc))
	case "pixel_loc":
		errDb = db.SetConfig_pixel_loc(msg.(*common.Pixel_loc))
	default:
		fmt.Printf("unknown name:%s\n", tableName)
		return errors.New("unknown name")
	}
	if errDb != nil {
		fmt.Printf("db write err:%v\n", err.Error())
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：数据库写入失败"))
		return err
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
		getConfigDb(w, r, "")
	}
}

func getConfig_pixel_loc(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "pixel_loc")
	case ConfigSqlite:
		getConfigDb(w, r, "pixel_loc")
	}
}

func getConfig_real_loc(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "real_loc")
	case ConfigSqlite:
		getConfigDb(w, r, "real_loc")
	}
}

func getConfig_crossing_setting(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "crossing_setting")
	case ConfigSqlite:
		getConfigDb(w, r, "crossing_setting")
	}
}

func getConfig_vibrate_setting(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "vibrate_setting")
	case ConfigSqlite:
		getConfigDb(w, r, "vibrate_setting")
	}
}

func getConfig_distance(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "distance")
	case ConfigSqlite:
		getConfigDb(w, r, "distance")
	}
}

func getConfig_base(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "base")
	case ConfigSqlite:
		getConfigDb(w, r, "base")
	}
}

func setConfig_all(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, ini.DefaultSection)
	case ConfigSqlite:
		setConfigDb(w, r, "")
	}
}

func setConfig_pixel_loc(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "pixel_loc")
	case ConfigSqlite:
		setConfigDb(w, r, "pixel_loc")
	}
}

func setConfig_real_loc(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "real_loc")
	case ConfigSqlite:
		setConfigDb(w, r, "real_loc")
	}
}

func setConfig_crossing_setting(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "crossing_setting")
	case ConfigSqlite:
		setConfigDb(w, r, "crossing_setting")
	}
}

func setConfig_vibrate_setting(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "vibrate_setting")
	case ConfigSqlite:
		setConfigDb(w, r, "vibrate_setting")
	}
}

func setConfig_distance(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "distance")
	case ConfigSqlite:
		setConfigDb(w, r, "distance")
	}
}

func setConfig_base(w http.ResponseWriter, r *http.Request) {
	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "base")
	case ConfigSqlite:
		setConfigDb(w, r, "base")
	}
}
