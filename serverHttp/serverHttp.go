package serverHttp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DataDog/gopsutil/process"
	"github.com/go-ini/ini"
	"github.com/wxnacy/wgo/arrays"
	"gohttpconfig/common"
	"gohttpconfig/configStruct"
	"gohttpconfig/db"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"strconv"
)

var IndexPath string = "./html"
var Config *ini.File
var IsConfigExist = false
var ConfigPath string
var ConfigCommunicate *ini.File
var isConfigExistCommunicate = false
var ConfigPathCommuniate string

var StaticFilePath string = "./image"

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

func Run(port int, htmlPath string) {
	IndexPath = htmlPath

	//开启静态web根目录
	http.Handle("/", http.FileServer(http.Dir(htmlPath)))
	/**config distanceN1**/
	//set
	http.HandleFunc("/setConfig_base", setConfig_base)
	http.HandleFunc("/setConfig_distance", setConfig_distance)
	http.HandleFunc("/setConfig_vibrate_setting", setConfig_vibrate_setting)
	http.HandleFunc("/setConfig_crossing_setting", setConfig_crossing_setting)
	http.HandleFunc("/setConfig_real_loc", setConfig_real_loc)
	http.HandleFunc("/setConfig_pixel_loc", setConfig_pixel_loc)
	http.HandleFunc("/setConfig_info", setConfig_info)
	//get
	http.HandleFunc("/getConfig_base", getConfig_base)
	http.HandleFunc("/getConfig_distance", getConfig_distance)
	http.HandleFunc("/getConfig_vibrate_setting", getConfig_vibrate_setting)
	http.HandleFunc("/getConfig_crossing_setting", getConfig_crossing_setting)
	http.HandleFunc("/getConfig_real_loc", getConfig_real_loc)
	http.HandleFunc("/getConfig_pixel_loc", getConfig_pixel_loc)
	http.HandleFunc("/getConfig_info", getConfig_info)

	/**config communicate**/
	//set
	http.HandleFunc("/setConfig_camera", setConfig_camera)
	http.HandleFunc("/setConfig_cloud", setConfig_cloud)
	http.HandleFunc("/setConfig_radar", setConfig_radar)
	http.HandleFunc("/setConfig_annuciator", setConfig_annuciator)
	http.HandleFunc("/setConfig_hardinfo", setConfig_hardinfo)
	http.HandleFunc("/setConfig_communicate", setConfig_communicate)
	//get
	http.HandleFunc("/getConfig_camera", getConfig_camera)
	http.HandleFunc("/getConfig_cloud", getConfig_cloud)
	http.HandleFunc("/getConfig_radar", getConfig_radar)
	http.HandleFunc("/getConfig_annuciator", getConfig_annuciator)
	http.HandleFunc("/getConfig_hardinfo", getConfig_hardinfo)
	http.HandleFunc("/getConfig_communicate", getConfig_communicate)
	/**reset proc**/
	http.HandleFunc("/resetProc", resetProc)
	/***file download**/
	http.HandleFunc("/getFiles", getFiles)
	http.HandleFunc("/getFile", getFile)
	//http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("./image"))))

	addr := ":" + strconv.Itoa(port)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func getFile(w http.ResponseWriter, r *http.Request) {
	//1.解析http请求
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return
	}
	fmt.Printf("body:%s\n", rBody)
	//2.获取 json中 file 属性，也就是文件的名称
	//type GetFile struct {
	//	File string `json:"file"`
	//}
	//fileInfo := GetFile{}
	//
	//err_json := json.Unmarshal(rBody, &fileInfo)
	//if err_json != nil {
	//	fmt.Println("json 解析错误:", err_json)
	//	return
	//}
	fileName := r.FormValue("filename")
	//3.下载文件
	file, _ := os.Open(StaticFilePath + "/" + fileName)
	defer file.Close()

	fileHeader := make([]byte, 512)
	file.Read(fileHeader)

	fileStat, _ := file.Stat()

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", http.DetectContentType(fileHeader))
	w.Header().Set("Content-Length", strconv.FormatInt(fileStat.Size(), 10))

	file.Seek(0, 0)
	io.Copy(w, file)

}

type FileInfo struct {
	FileName string
	DateTime string
	FileSize int64
}

func getDirList(path string) ([]FileInfo, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	fileList := make([]FileInfo, 0)
	for _, file := range files {
		//只取文件，不取文件夹
		if !file.IsDir() {
			fileList = append(fileList, FileInfo{file.Name(), file.ModTime().String(), file.Size()})
		}
	}
	return fileList, err
}

func getFiles(w http.ResponseWriter, r *http.Request) {
	//1.获取指定目录下的所有文件名称
	fileList, err := getDirList(StaticFilePath)
	if err != nil {
		fmt.Println("get files fail")
		w.WriteHeader(http.StatusGone)
		w.Write([]byte(err.Error()))
		return
	}
	if fileList != nil {
		//HTML模板路径
		htmlTmplPath := IndexPath + "/template/files.html"
		tmpl, err_tmpl := template.ParseFiles(htmlTmplPath)
		if err_tmpl != nil {
			fmt.Println("读取模板文件失败：", err_tmpl)
			return
		}
		fmt.Println("get files success")
		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, fileList)

	}
}

/*****************config distanceN1****************/
// 基础函数

func configStruct2common(dst interface{}, src interface{}, name string) error {
	switch name {
	case ini.DefaultSection, "":
		common.ConfigStruct2Common_info(dst.(*common.Info), src.(*configStruct.Info))
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
	case ini.DefaultSection, "":
		common.Common2ConfigStruct_info(dst.(*configStruct.Info), src.(*common.Info))
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
	Config, err = ini.Load(ConfigPath)
	if err != nil {
		fmt.Printf("cant not load ini file:%s\n", ConfigPath)
		IsConfigExist = false
	} else {
		IsConfigExist = true
	}
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
	Config, err = ini.Load(ConfigPath)
	if err != nil {
		fmt.Printf("cant not load ini file:%s\n", ConfigPath)
		IsConfigExist = false
	} else {
		IsConfigExist = true
	}
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
	case "base":
		dst = &common.Base{}
		src = &configStruct.Base{}
	case "distance":
		dst = &common.Distance{}
		src = &configStruct.Distance{}
	case "vibrate_setting":
		dst = &common.Vibrate_setting{}
		src = &configStruct.Vibrate_setting{}
	case "crossing_setting":
		dst = &common.Crossing_setting{}
		src = &configStruct.Crossing_setting{}
	case "real_loc":
		dst = &common.Real_loc{}
		src = &configStruct.Real_loc{}
	case "pixel_loc":
		dst = &common.Pixel_loc{}
		src = &configStruct.Pixel_loc{}
	default:
		fmt.Printf("unknown name:%s\n", tableName)
		return errors.New("unknown name")
	}
	//3.1获取
	errDb = db.GetConfig(tableName, src)
	if errDb != nil {
		fmt.Printf("db get fail:%v\n", errDb.Error())
		w.Write([]byte("失败：数据库读取失败"))
		return errDb
	}

	//3.2转换
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
	errDb := db.SetConfig(tableName, dst)
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
func getConfig_info(w http.ResponseWriter, r *http.Request) {
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

func setConfig_info(w http.ResponseWriter, r *http.Request) {
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

/**********config communicate*************/
func configStruct2commonCommunicate(dst interface{}, src interface{}, name string) error {
	switch name {
	case ini.DefaultSection, "":
		common.ConfigStruct2Common_communicate(dst.(*common.Communicate), src.(*configStruct.Communicate))
	case "camera":
		common.ConfigStruct2Common_camera(dst.(*common.Camera), src.(*configStruct.Camera))
	case "cloud":
		common.ConfigStruct2Common_cloud(dst.(*common.Cloud), src.(*configStruct.Cloud))
	case "radar":
		common.ConfigStruct2Common_radar(dst.(*common.Radar), src.(*configStruct.Radar))
	case "annuciator":
		common.ConfigStruct2Common_annuciator(dst.(*common.Annuciator), src.(*configStruct.Annuciator))
	case "hardinfo":
		common.ConfigStruct2Common_hardinfo(dst.(*common.HardInfo), src.(*configStruct.HardInfo))
	default:
		fmt.Printf("unknown name:%s\n", name)
		return errors.New("unknown name")
	}
	return nil
}

func common2configStructCommunicate(dst interface{}, src interface{}, name string) error {
	switch name {
	case ini.DefaultSection, "":
		common.Common2ConfigStruct_communicate(dst.(*configStruct.Communicate), src.(*common.Communicate))
	case "camera":
		common.Common2ConfigStruct_camera(dst.(*configStruct.Camera), src.(*common.Camera))
	case "cloud":
		common.Common2ConfigStruct_cloud(dst.(*configStruct.Cloud), src.(*common.Cloud))
	case "radar":
		common.Common2ConfigStruct_radar(dst.(*configStruct.Radar), src.(*common.Radar))
	case "annuciator":
		common.Common2ConfigStruct_annuciator(dst.(*configStruct.Annuciator), src.(*common.Annuciator))
	case "hardinfo":
		common.Common2ConfigStruct_hardinfo(dst.(*configStruct.HardInfo), src.(*common.HardInfo))
	default:
		fmt.Printf("unknown name:%s\n", name)
		return errors.New("unknown name")
	}
	return nil
}

func getConfigIniCommunicate(w http.ResponseWriter, r *http.Request, sectionName string) error {
	//1.解析http请求
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return err
	}
	fmt.Printf("body:%s\n", rBody)

	//2.配置文件不存在相应失败 退出
	ConfigCommunicate, err = ini.Load(ConfigPathCommuniate)
	if err != nil {
		fmt.Printf("cant not load ini file:%s\n", ConfigPathCommuniate)
		isConfigExistCommunicate = false
	} else {
		isConfigExistCommunicate = true
	}

	if !isConfigExistCommunicate {
		fmt.Printf("config file not exist\n")
		w.Write([]byte("失败：配置文件不存在"))
		return errors.New("配置文件不存在")
	}

	//3.读取ini文件指定分区信息，转换为json信息
	var src interface{}
	var dst interface{}

	switch sectionName {
	case ini.DefaultSection:
		src = &configStruct.Communicate{}
		dst = &common.Communicate{}
	case "camera":
		src = &configStruct.Camera{}
		dst = &common.Camera{}
	case "cloud":
		src = &configStruct.Cloud{}
		dst = &common.Cloud{}
	case "radar":
		src = &configStruct.Radar{}
		dst = &common.Radar{}
	case "annuciator":
		src = &configStruct.Annuciator{}
		dst = &common.Annuciator{}
	case "hardinfo":
		src = &configStruct.HardInfo{}
		dst = &common.HardInfo{}
	default:
		fmt.Printf("unknown name:%s\n", sectionName)
		return errors.New("unknown name")
	}
	//3.1读取
	section, errGetsection := ConfigCommunicate.GetSection(sectionName)
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
	errChange := configStruct2commonCommunicate(dst, src, sectionName)
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

func setConfigIniCommunicate(w http.ResponseWriter, r *http.Request, sectionName string) error {
	//1.解析http请求
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return err
	}
	fmt.Printf("body:%s\n", rBody)

	//2.配置文件不存在相应失败 退出
	ConfigCommunicate, err = ini.Load(ConfigPathCommuniate)
	if err != nil {
		fmt.Printf("cant not load ini file:%s\n", ConfigPathCommuniate)
		isConfigExistCommunicate = false
	} else {
		isConfigExistCommunicate = true
	}

	if !isConfigExistCommunicate {
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
		src = &common.Communicate{}
		dst = &configStruct.Communicate{}
	case "camera":
		src = &common.Camera{}
		dst = &configStruct.Camera{}
	case "cloud":
		src = &common.Cloud{}
		dst = &configStruct.Cloud{}
	case "radar":
		src = &common.Radar{}
		dst = &configStruct.Radar{}
	case "annuciator":
		src = &common.Annuciator{}
		dst = &configStruct.Annuciator{}
	case "hardinfo":
		src = &common.HardInfo{}
		dst = &configStruct.HardInfo{}
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
	errChange := common2configStructCommunicate(dst, src, sectionName)
	if errChange != nil {
		return errChange
	}

	//4.结构体写入ini分区
	section, _ := ConfigCommunicate.NewSection(sectionName)
	errSection := section.ReflectFrom(dst)
	if errSection != nil {
		fmt.Printf("ini map jsonBase fail:%v\n", errSection.Error())
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败:ini写入分区失败"))
		return errSection
	}

	errSave := ConfigCommunicate.SaveTo(ConfigPathCommuniate)
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
func getConfig_communicate(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	//switch ConfigType {
	//case ConfigIni:
	//	getConfigIni(w, r, ini.DefaultSection)
	//case ConfigSqlite:
	//	getConfigDb(w, r, "")
	//}

	//暂时强制配置方式为ini
	getConfigIniCommunicate(w, r, ini.DefaultSection)
}

func getConfig_hardinfo(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	//switch ConfigType {
	//case ConfigIni:
	//	getConfigIni(w, r, "hardinfo")
	//case ConfigSqlite:
	//	getConfigDb(w, r, "hardinfo")
	//}

	//暂时强制配置方式为ini
	getConfigIniCommunicate(w, r, "hardinfo")
}

func getConfig_annuciator(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	//switch ConfigType {
	//case ConfigIni:
	//	getConfigIni(w, r, "annuciator")
	//case ConfigSqlite:
	//	getConfigDb(w, r, "annuciator")
	//}

	//暂时强制配置方式为ini
	getConfigIniCommunicate(w, r, "annuciator")
}

func getConfig_radar(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	//switch ConfigType {
	//case ConfigIni:
	//	getConfigIni(w, r, "radar")
	//case ConfigSqlite:
	//	getConfigDb(w, r, "radar")
	//}

	//暂时强制配置方式为ini
	getConfigIniCommunicate(w, r, "radar")
}

func getConfig_cloud(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	//switch ConfigType {
	//case ConfigIni:
	//	getConfigIni(w, r, "cloud")
	//case ConfigSqlite:
	//	getConfigDb(w, r, "cloud")
	//}

	//暂时强制配置方式为ini
	getConfigIniCommunicate(w, r, "cloud")
}

func getConfig_camera(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	//switch ConfigType {
	//case ConfigIni:
	//	getConfigIni(w, r, "camera")
	//case ConfigSqlite:
	//	getConfigDb(w, r, "camera")
	//}

	//暂时强制配置方式为ini
	getConfigIniCommunicate(w, r, "camera")
}

func setConfig_communicate(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	//switch ConfigType {
	//case ConfigIni:
	//	setConfigIni(w, r, ini.DefaultSection)
	//case ConfigSqlite:
	//	setConfigDb(w, r, "")
	//}

	//暂时强制配置方式为ini
	setConfigIniCommunicate(w, r, ini.DefaultSection)
}

func setConfig_hardinfo(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	//switch ConfigType {
	//case ConfigIni:
	//	setConfigIni(w, r, "hardinfo")
	//case ConfigSqlite:
	//	setConfigDb(w, r, "hardinfo")
	//}

	//暂时强制配置方式为ini
	setConfigIniCommunicate(w, r, "hardinfo")

}

func setConfig_annuciator(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	//switch ConfigType {
	//case ConfigIni:
	//	setConfigIni(w, r, "annuciator")
	//case ConfigSqlite:
	//	setConfigDb(w, r, "annuciator")
	//}

	//暂时强制配置方式为ini
	setConfigIniCommunicate(w, r, "annuciator")
}

func setConfig_radar(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	//switch ConfigType {
	//case ConfigIni:
	//	setConfigIni(w, r, "radar")
	//case ConfigSqlite:
	//	setConfigDb(w, r, "radar")
	//}

	//暂时强制配置方式为ini
	setConfigIniCommunicate(w, r, "radar")
}

func setConfig_cloud(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	//switch ConfigType {
	//case ConfigIni:
	//	setConfigIni(w, r, "cloud")
	//case ConfigSqlite:
	//	setConfigDb(w, r, "cloud")
	//}

	//暂时强制配置方式为ini
	setConfigIniCommunicate(w, r, "cloud")
}

func setConfig_camera(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()

	//switch ConfigType {
	//case ConfigIni:
	//	setConfigIni(w, r, "camera")
	//case ConfigSqlite:
	//	setConfigDb(w, r, "camera")
	//}

	//暂时强制配置方式为ini
	setConfigIniCommunicate(w, r, "camera")
}

/****************reset proc*********************/

func processId() (pid []int32) {
	pids, _ := process.Pids()
	for _, p := range pids {
		pid = append(pid, p)
	}
	return pid
}

func processName() (pname []string) {
	pids, _ := process.Pids()
	for _, p := range pids {
		pn, _ := process.NewProcess(p)
		pName, _ := pn.Name()
		pname = append(pname, pName)
	}
	return pname
}

func resetProc(w http.ResponseWriter, r *http.Request) {
	defer func() {
		err := recover()
		switch err.(type) {
		case runtime.Error: //运行时错误
			fmt.Println("run time err:", err)
		}
	}()
	//1.解析http请求
	rBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("req body read err:%v\n", err.Error())
		return
	}
	fmt.Printf("body:%s\n", rBody)

	//2.将请求主体转换为json结构体
	var req common.ResetProc
	//2.1 读取
	err = json.Unmarshal(rBody, &req)
	if err != nil {
		fmt.Printf("json unmarshal err:%v\n", err.Error())
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：json解析失败"))
		return
	}
	//2.2获取程序的名称
	procName := req.Proc
	//2.3查找系统进程的名称，看是否存在此进程
	procNames := processName()
	index := arrays.ContainsString(procNames, procName)
	if index == -1 {
		fmt.Printf("proc:%s not run\n", procName)
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：程序未运行"))
		return
	}

	//2.3杀死程序进程
	shell := "killall -9 " + procName
	cmd := exec.Command("/bin/bash", "-c", shell)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("cmd %s exec fail:%v\n", cmd.String(), err.Error())
		w.WriteHeader(http.StatusGone)
		w.Write([]byte("失败：重启失败"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("成功：重启程序"))

}
