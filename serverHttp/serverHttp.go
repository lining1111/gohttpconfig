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
	"runtime"
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
func configStruct2common(dst interface{}, src interface{}, name string) error {
	switch name {
	case ini.DefaultSection:
	case "":
		common.ConfigStruct2Common_all(dst.(*common.Info), src.(*configStruct.Info))
	case "base":
		common.ConfigStruct2Common_base(dst.(*common.Base), src.(*configStruct.Base))
	case "distance":
		common.ConfigStruct2Common_distance(dst.(*common.Distance), src.(*configStruct.Distance))
	case "vibrate_setting":
		common.ConfigStruct2Common_vibrate_setting(dst.(*common.Vibrate_setting), src.(*configStruct.Vibrate_setting))
	case "crossing_setting":
		common.ConfigStruct2Common_crossing_setting(dst.(*common.Crossing_setting), src.(*configStruct.Crossing_setting))
	case "real_loc":
		common.ConfigStruct2Common_real_loc(dst.(*common.Real_loc), src.(*configStruct.Real_loc))
	case "pixel_loc":
		common.ConfigStruct2Common_pixel_loc(dst.(*common.Pixel_loc), src.(*configStruct.Pixel_loc))
	default:
		fmt.Printf("unknown name:%s\n", name)
		return errors.New("unknown name")
	}
	return nil
}

func common2configStruct(dst interface{}, src interface{}, name string) error {
	switch name {
	case ini.DefaultSection:
	case "":
		common.Common2ConfigStruct_all(dst.(*configStruct.Info), src.(*common.Info))
	case "base":
		common.Common2ConfigStruct_base(dst.(*configStruct.Base), src.(*common.Base))
	case "distance":
		common.Common2ConfigStruct_distance(dst.(*configStruct.Distance), src.(*common.Distance))
	case "vibrate_setting":
		common.Common2ConfigStruct_vibrate_setting(dst.(*configStruct.Vibrate_setting), src.(*common.Vibrate_setting))
	case "crossing_setting":
		common.Common2ConfigStruct_crossing_setting(dst.(*configStruct.Crossing_setting), src.(*common.Crossing_setting))
	case "real_loc":
		common.Common2ConfigStruct_real_loc(dst.(*configStruct.Real_loc), src.(*common.Real_loc))
	case "pixel_loc":
		common.Common2ConfigStruct_pixel_loc(dst.(*configStruct.Pixel_loc), src.(*common.Pixel_loc))
	default:
		fmt.Printf("unknown name:%s\n", name)
		return errors.New("unknown name")
	}
	return nil
}

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
	var src interface{}
	var dst interface{}

	switch sectionName {
	case ini.DefaultSection:
		src = &configStruct.Info{}
		dst = &common.Info{}
	case "base":
		src = &configStruct.Base{}
		dst = &common.Base{}
	case "distance":
		src = &configStruct.Distance{}
		dst = &common.Distance{}
	case "vibrate_setting":
		src = &configStruct.Vibrate_setting{}
		dst = &common.Vibrate_setting{}
	case "crossing_setting":
		src = &configStruct.Crossing_setting{}
		dst = &common.Crossing_setting{}
	case "real_loc":
		src = &configStruct.Real_loc{}
		dst = &common.Real_loc{}
	case "pixel_loc":
		src = &configStruct.Pixel_loc{}
		dst = &common.Pixel_loc{}
	default:
		fmt.Printf("unknown name:%s\n", sectionName)
		return errors.New("unknown name")
	}
	//3.1读取
	section, errGetsection := Config.GetSection(sectionName)
	if errGetsection != nil {
		fmt.Printf("获取分区失败:%v\n", errGetsection.Error())
		w.Write([]byte("获取分区失败"))
		return errGetsection
	}
	errSection := section.MapTo(src)
	if errSection != nil {
		fmt.Printf("分区信息转换失败：%v\n", errSection.Error())
		w.Write([]byte("分区信息转换失败"))
		return errSection
	}
	//3.2转换
	errChange := configStruct2common(dst, src, sectionName)
	if errChange != nil {
		return errChange
	}

	//4.json信息组织回复
	wBody, errBody := json.Marshal(dst)
	if errBody != nil {
		fmt.Printf("json unmarshal err:%v\n", errBody.Error())
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
	//3.读取ini文件指定分区信息，转换为json信息
	var src interface{}
	var dst interface{}

	switch sectionName {
	case ini.DefaultSection:
		src = &common.Info{}
		dst = &configStruct.Info{}
	case "base":
		src = &common.Base{}
		dst = &configStruct.Base{}
	case "distance":
		src = &common.Distance{}
		dst = &configStruct.Distance{}
	case "vibrate_setting":
		src = &common.Vibrate_setting{}
		dst = &configStruct.Vibrate_setting{}
	case "crossing_setting":
		src = &common.Crossing_setting{}
		dst = &configStruct.Crossing_setting{}
	case "real_loc":
		src = &common.Real_loc{}
		dst = &configStruct.Real_loc{}
	case "pixel_loc":
		src = &common.Pixel_loc{}
		dst = &configStruct.Pixel_loc{}
	default:
		fmt.Printf("unknown name:%s\n", sectionName)
		return errors.New("unknown name")
	}
	//3.1 读取
	err = json.Unmarshal(rBody, src)
	if err != nil {
		fmt.Printf("json unmarshal err:%v\n", err.Error())
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：json解析失败"))
		return err
	}
	//3.2转换
	errChange := common2configStruct(dst, src, sectionName)
	if errChange != nil {
		return errChange
	}

	//4.结构体写入ini分区
	section, _ := Config.NewSection(sectionName)
	errSection := section.ReflectFrom(dst)
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
	var src interface{}
	var dst interface{}
	var errDb error
	switch tableName {
	case "":
		dst = &common.Info{}
		src = &configStruct.Info{}
		errDb = db.GetConfig_all(src.(*configStruct.Info))
	case "base":
		dst = &common.Base{}
		src = &configStruct.Base{}
		errDb = db.GetConfig_base(src.(*configStruct.Base))
	case "distance":
		dst = &common.Distance{}
		src = &configStruct.Distance{}
		errDb = db.GetConfig_distance(src.(*configStruct.Distance))
	case "vibrate_setting":
		dst = &common.Vibrate_setting{}
		src = &configStruct.Vibrate_setting{}
		errDb = db.GetConfig_vibrate_setting(src.(*configStruct.Vibrate_setting))
	case "crossing_setting":
		dst = &common.Crossing_setting{}
		src = &configStruct.Crossing_setting{}
		errDb = db.GetConfig_crossing_setting(src.(*configStruct.Crossing_setting))
	case "real_loc":
		dst = &common.Real_loc{}
		src = &configStruct.Real_loc{}
		errDb = db.GetConfig_real_loc(src.(*configStruct.Real_loc))
	case "pixel_loc":
		dst = &common.Pixel_loc{}
		src = &configStruct.Pixel_loc{}
		errDb = db.GetConfig_pixel_loc(src.(*configStruct.Pixel_loc))
	default:
		fmt.Printf("unknown name:%s\n", tableName)
		return errors.New("unknown name")
	}
	if errDb != nil {
		fmt.Printf("db get fail:%v\n", errDb.Error())
		w.Write([]byte("失败：数据库读取失败"))
		return errDb
	}
	//3.1转换
	errChange := configStruct2common(dst, src, tableName)
	if errChange != nil {
		fmt.Printf("change fail:%v\n", errChange.Error())
		w.Write([]byte("失败：转换失败"))
		return errChange
	}

	//4.json信息组织回复
	wBody, errBody := json.Marshal(dst)
	if errBody != nil {
		fmt.Printf("json unmarshal err:%v\n", errBody.Error())
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
	var src interface{}
	var dst interface{}

	switch tableName {
	case "":
		src = &common.Info{}
		dst = &configStruct.Info{}
	case "base":
		src = &common.Base{}
		dst = &configStruct.Base{}
	case "distance":
		src = &common.Distance{}
		dst = &configStruct.Distance{}
	case "vibrate_setting":
		src = &common.Vibrate_setting{}
		dst = &configStruct.Vibrate_setting{}
	case "crossing_setting":
		src = &common.Crossing_setting{}
		dst = &configStruct.Crossing_setting{}
	case "real_loc":
		src = &common.Real_loc{}
		dst = &configStruct.Real_loc{}
	case "pixel_loc":
		src = &common.Pixel_loc{}
		dst = &configStruct.Pixel_loc{}
	default:
		fmt.Printf("unknown name:%s\n", tableName)
		return errors.New("unknown name")
	}
	//3.1读取json设置
	err = json.Unmarshal(rBody, src)
	if err != nil {
		fmt.Printf("json unmarshal err:%v\n", err.Error())
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：json解析失败"))
		return err
	}
	//3.2转换
	errChange := common2configStruct(dst, src, tableName)
	if errChange != nil {
		fmt.Printf("change err:%v\n", errChange.Error())
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：转换失败"))
		return errChange
	}

	//4.结构体写入指定的数据库表
	var errDb error
	switch tableName {
	case "":
		errDb = db.SetConfig_all(dst.(*configStruct.Info))
	case "base":
		errDb = db.SetConfig_base(dst.(*configStruct.Base))
	case "distance":
		errDb = db.SetConfig_distance(dst.(*configStruct.Distance))
	case "vibrate_setting":
		errDb = db.SetConfig_vibrate_setting(dst.(*configStruct.Vibrate_setting))
	case "crossing_setting":
		errDb = db.SetConfig_crossing_setting(dst.(*configStruct.Crossing_setting))
	case "real_loc":
		errDb = db.SetConfig_real_loc(dst.(*configStruct.Real_loc))
	case "pixel_loc":
		errDb = db.SetConfig_pixel_loc(dst.(*configStruct.Pixel_loc))
	default:
		fmt.Printf("unknown name:%s\n", tableName)
		return errors.New("unknown name")
	}
	if errDb != nil {
		fmt.Printf("db write err:%v\n", errDb.Error())
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：数据库写入失败"))
		return err
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("成功：数据库写入成功"))
	return nil
}

//web api
func getConfig_all(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, ini.DefaultSection)
	case ConfigSqlite:
		getConfigDb(w, r, "")
	}
}

func getConfig_pixel_loc(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "pixel_loc")
	case ConfigSqlite:
		getConfigDb(w, r, "pixel_loc")
	}
}

func getConfig_real_loc(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "real_loc")
	case ConfigSqlite:
		getConfigDb(w, r, "real_loc")
	}
}

func getConfig_crossing_setting(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "crossing_setting")
	case ConfigSqlite:
		getConfigDb(w, r, "crossing_setting")
	}
}

func getConfig_vibrate_setting(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "vibrate_setting")
	case ConfigSqlite:
		getConfigDb(w, r, "vibrate_setting")
	}
}

func getConfig_distance(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "distance")
	case ConfigSqlite:
		getConfigDb(w, r, "distance")
	}
}

func getConfig_base(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	switch ConfigType {
	case ConfigIni:
		getConfigIni(w, r, "base")
	case ConfigSqlite:
		getConfigDb(w, r, "base")
	}
}

func setConfig_all(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, ini.DefaultSection)
	case ConfigSqlite:
		setConfigDb(w, r, "")
	}
}

func setConfig_pixel_loc(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "pixel_loc")
	case ConfigSqlite:
		setConfigDb(w, r, "pixel_loc")
	}
}

func setConfig_real_loc(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "real_loc")
	case ConfigSqlite:
		setConfigDb(w, r, "real_loc")
	}
}

func setConfig_crossing_setting(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "crossing_setting")
	case ConfigSqlite:
		setConfigDb(w, r, "crossing_setting")
	}
}

func setConfig_vibrate_setting(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "vibrate_setting")
	case ConfigSqlite:
		setConfigDb(w, r, "vibrate_setting")
	}
}

func setConfig_distance(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "distance")
	case ConfigSqlite:
		setConfigDb(w, r, "distance")
	}
}

func setConfig_base(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	switch ConfigType {
	case ConfigIni:
		setConfigIni(w, r, "base")
	case ConfigSqlite:
		setConfigDb(w, r, "base")
	}
}
