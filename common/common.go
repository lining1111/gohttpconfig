package common

import (
	"gohttpconfig/configStruct"
	"strconv"
)

//base
type Base struct {
	Width  string `json:"width"`
	Height string `json:"height"`
}

//distance
type Distance struct {
	X_distance  string `json:"x_distance"`
	Y_distance  string `json:"y_distance"`
	Altitude    string `json:"altitude"`
	Y_value     string `json:"y_value"`
	Coefficient string `json:"coefficient"`

	Matrix00    string `json:"matrix00"`
	Matrix01    string `json:"matrix01"`
	Matrix02    string `json:"matrix02"`
	Matrix10    string `json:"matrix10"`
	Matrix11    string `json:"matrix11"`
	Matrix12    string `json:"matrix12"`
	Matrix20    string `json:"matrix20"`
	Matrix21    string `json:"matrix21"`
	Matrix22    string `json:"matrix22"`
	Radar_x     string `json:"radar_x"`
	Radar_y     string `json:"radar_y"`
	Radar_theta string `json:"radar_theta"`
	MPPW        string `json:"MPPW"`
	MPPH        string `json:"MPPH"`
}

//vibrate_setting
type Vibrate_setting struct {
	X_vibrate_max  string `json:"x_vibrate_max"`
	Y_vibrate_max  string `json:"y_vibrate_max"`
	MatchBoxX      string `json:"matchBoxX"`
	MatchBoxY      string `json:"matchBoxY"`
	MatchBoxWidth  string `json:"matchBoxWidth"`
	MatchBoxHeight string `json:"matchBoxHeight"`
}

//crossing_setting
type Crossing_setting struct {
	Orientations string `json:"orientations"`
	Deltax_south string `json:"deltax_south"`
	Deltay_south string `json:"deltay_south"`
	Deltax_north string `json:"deltax_north"`
	Deltay_north string `json:"deltay_north"`
	Deltax_west  string `json:"deltax_west"`
	Deltay_west  string `json:"deltay_west"`
	Deltax_east  string `json:"deltax_east"`
	Deltay_east  string `json:"deltay_east"`
	Flag_south   string `json:"flag_south"`
	Flag_north   string `json:"flag_north"`
	Flag_west    string `json:"flag_west"`
	Flag_east    string `json:"flag_east"`
	WidthX       string `json:"widthX"`
	WidthY       string `json:"widthY"`
}

//real_loc
type Real_loc struct {
	Real_left_point_x   string `json:"real_left_point_x"`
	Real_left_point_y   string `json:"real_left_point_y"`
	Real_right_point_x  string `json:"real_right_point_x"`
	Real_right_point_y  string `json:"real_right_point_y"`
	Real_top_point_x    string `json:"real_top_point_x"`
	Real_top_point_y    string `json:"real_top_point_y"`
	Real_bottom_point_x string `json:"real_bottom_point_x"`
	Real_bottom_point_y string `json:"real_bottom_point_y"`
}

//pixel_loc
type Pixel_loc struct {
	Pixel_left_point_x   string `json:"pixel_left_point_x"`
	Pixel_left_point_y   string `json:"pixel_left_point_y"`
	Pixel_right_point_x  string `json:"pixel_right_point_x"`
	Pixel_right_point_y  string `json:"pixel_right_point_y"`
	Pixel_top_point_x    string `json:"pixel_top_point_x"`
	Pixel_top_point_y    string `json:"pixel_top_point_y"`
	Pixel_bottom_point_x string `json:"pixel_bottom_point_x"`
	Pixel_bottom_point_y string `json:"pixel_bottom_point_y"`
}

//info
type Info struct {
	Base             Base             `json:"base"`
	Distance         Distance         `json:"distance"`
	Vibrate_setting  Vibrate_setting  `json:"vibrate_setting"`
	Crossing_setting Crossing_setting `json:"crossing_setting"`
	Real_loc         Real_loc         `json:"real_loc"`
	Pixel_loc        Pixel_loc        `json:"pixel_loc"`
}

func Common2ConfigStruct_base(dst *configStruct.Base, src *Base) {
	dst.Width, _ = strconv.Atoi(src.Width)
	dst.Height, _ = strconv.Atoi(src.Height)
}
func Common2ConfigStruct_distance(dst *configStruct.Distance, src *Distance) {
	dst.X_distance, _ = strconv.ParseFloat(src.X_distance, 64)
	dst.Y_distance, _ = strconv.ParseFloat(src.Y_distance, 64)
	dst.Altitude, _ = strconv.ParseFloat(src.Altitude, 64)
	dst.Y_value, _ = strconv.ParseFloat(src.Y_value, 64)
	dst.Coefficient, _ = strconv.ParseFloat(src.Coefficient, 64)
	dst.Matrix00, _ = strconv.ParseFloat(src.Matrix00, 64)
	dst.Matrix01, _ = strconv.ParseFloat(src.Matrix01, 64)
	dst.Matrix02, _ = strconv.ParseFloat(src.Matrix02, 64)
	dst.Matrix10, _ = strconv.ParseFloat(src.Matrix10, 64)
	dst.Matrix11, _ = strconv.ParseFloat(src.Matrix11, 64)
	dst.Matrix12, _ = strconv.ParseFloat(src.Matrix12, 64)
	dst.Matrix20, _ = strconv.ParseFloat(src.Matrix20, 64)
	dst.Matrix21, _ = strconv.ParseFloat(src.Matrix21, 64)
	dst.Matrix22, _ = strconv.ParseFloat(src.Matrix22, 64)
	dst.Radar_x, _ = strconv.ParseFloat(src.Radar_x, 64)
	dst.Radar_y, _ = strconv.ParseFloat(src.Radar_y, 64)
	dst.Radar_theta, _ = strconv.ParseFloat(src.Radar_theta, 64)
	dst.MPPW, _ = strconv.ParseFloat(src.MPPW, 64)
	dst.MPPH, _ = strconv.ParseFloat(src.MPPH, 64)
}
func Common2ConfigStruct_vibrate_setting(dst *configStruct.Vibrate_setting, src *Vibrate_setting) {
	dst.X_vibrate_max, _ = strconv.ParseFloat(src.X_vibrate_max, 64)
	dst.Y_vibrate_max, _ = strconv.ParseFloat(src.Y_vibrate_max, 64)
	dst.MatchBoxX, _ = strconv.ParseFloat(src.MatchBoxX, 64)
	dst.MatchBoxY, _ = strconv.ParseFloat(src.MatchBoxY, 64)
	dst.MatchBoxWidth, _ = strconv.ParseFloat(src.MatchBoxWidth, 64)
	dst.MatchBoxHeight, _ = strconv.ParseFloat(src.MatchBoxHeight, 64)
}
func Common2ConfigStruct_crossing_setting(dst *configStruct.Crossing_setting, src *Crossing_setting) {
	dst.Orientations, _ = strconv.Atoi(src.Orientations)
	dst.Deltax_south, _ = strconv.ParseFloat(src.Deltax_south, 64)
	dst.Deltay_south, _ = strconv.ParseFloat(src.Deltay_south, 64)
	dst.Deltax_north, _ = strconv.ParseFloat(src.Deltax_north, 64)
	dst.Deltay_north, _ = strconv.ParseFloat(src.Deltay_north, 64)
	dst.Deltax_west, _ = strconv.ParseFloat(src.Deltax_west, 64)
	dst.Deltay_west, _ = strconv.ParseFloat(src.Deltay_west, 64)
	dst.Deltax_east, _ = strconv.ParseFloat(src.Deltax_east, 64)
	dst.Deltay_east, _ = strconv.ParseFloat(src.Deltay_east, 64)
	dst.Flag_south, _ = strconv.Atoi(src.Flag_south)
	dst.Flag_north, _ = strconv.Atoi(src.Flag_north)
	dst.Flag_west, _ = strconv.Atoi(src.Flag_west)
	dst.Flag_east, _ = strconv.Atoi(src.Flag_east)
	dst.WidthX, _ = strconv.ParseFloat(src.WidthX, 64)
	dst.WidthY, _ = strconv.ParseFloat(src.WidthY, 64)
}
func Common2ConfigStruct_real_loc(dst *configStruct.Real_loc, src *Real_loc) {
	dst.Real_left_point_x, _ = strconv.ParseFloat(src.Real_left_point_x, 64)
	dst.Real_left_point_y, _ = strconv.ParseFloat(src.Real_left_point_y, 64)
	dst.Real_right_point_x, _ = strconv.ParseFloat(src.Real_right_point_x, 64)
	dst.Real_right_point_y, _ = strconv.ParseFloat(src.Real_right_point_y, 64)
	dst.Real_top_point_x, _ = strconv.ParseFloat(src.Real_top_point_x, 64)
	dst.Real_top_point_y, _ = strconv.ParseFloat(src.Real_top_point_y, 64)
	dst.Real_bottom_point_x, _ = strconv.ParseFloat(src.Real_bottom_point_x, 64)
	dst.Real_bottom_point_y, _ = strconv.ParseFloat(src.Real_bottom_point_y, 64)
}
func Common2ConfigStruct_pixel_loc(dst *configStruct.Pixel_loc, src *Pixel_loc) {
	dst.Pixel_left_point_x, _ = strconv.ParseFloat(src.Pixel_left_point_x, 64)
	dst.Pixel_left_point_y, _ = strconv.ParseFloat(src.Pixel_left_point_y, 64)
	dst.Pixel_right_point_x, _ = strconv.ParseFloat(src.Pixel_right_point_x, 64)
	dst.Pixel_right_point_y, _ = strconv.ParseFloat(src.Pixel_right_point_y, 64)
	dst.Pixel_top_point_x, _ = strconv.ParseFloat(src.Pixel_top_point_x, 64)
	dst.Pixel_top_point_y, _ = strconv.ParseFloat(src.Pixel_top_point_y, 64)
	dst.Pixel_bottom_point_x, _ = strconv.ParseFloat(src.Pixel_bottom_point_x, 64)
	dst.Pixel_bottom_point_y, _ = strconv.ParseFloat(src.Pixel_bottom_point_y, 64)
}

func Common2ConfigStruct_info(dst *configStruct.Info, src *Info) {
	Common2ConfigStruct_base(&(dst.Base), &(src.Base))
	Common2ConfigStruct_distance(&(dst.Distance), &(src.Distance))
	Common2ConfigStruct_vibrate_setting(&(dst.Vibrate_setting), &(src.Vibrate_setting))
	Common2ConfigStruct_crossing_setting(&(dst.Crossing_setting), &(src.Crossing_setting))
	Common2ConfigStruct_real_loc(&(dst.Real_loc), &(src.Real_loc))
	Common2ConfigStruct_pixel_loc(&(dst.Pixel_loc), &(src.Pixel_loc))
}

func ConfigStruct2Common_base(dst *Base, src *configStruct.Base) {
	dst.Width = strconv.Itoa(src.Width)
	dst.Height = strconv.Itoa(src.Height)
}

func ConfigStruct2Common_distance(dst *Distance, src *configStruct.Distance) {
	dst.X_distance = strconv.FormatFloat(src.X_distance, 'f', -1, 64)
	dst.Y_distance = strconv.FormatFloat(src.Y_distance, 'f', -1, 64)
	dst.Altitude = strconv.FormatFloat(src.Altitude, 'f', -1, 64)
	dst.Y_value = strconv.FormatFloat(src.Y_value, 'f', -1, 64)
	dst.Coefficient = strconv.FormatFloat(src.Coefficient, 'f', -1, 64)
	dst.Matrix00 = strconv.FormatFloat(src.Matrix00, 'f', -1, 64)
	dst.Matrix01 = strconv.FormatFloat(src.Matrix01, 'f', -1, 64)
	dst.Matrix02 = strconv.FormatFloat(src.Matrix02, 'f', -1, 64)
	dst.Matrix10 = strconv.FormatFloat(src.Matrix10, 'f', -1, 64)
	dst.Matrix11 = strconv.FormatFloat(src.Matrix11, 'f', -1, 64)
	dst.Matrix12 = strconv.FormatFloat(src.Matrix12, 'f', -1, 64)
	dst.Matrix20 = strconv.FormatFloat(src.Matrix20, 'f', -1, 64)
	dst.Matrix21 = strconv.FormatFloat(src.Matrix21, 'f', -1, 64)
	dst.Matrix22 = strconv.FormatFloat(src.Matrix22, 'f', -1, 64)
	dst.Radar_x = strconv.FormatFloat(src.Radar_x, 'f', -1, 64)
	dst.Radar_y = strconv.FormatFloat(src.Radar_y, 'f', -1, 64)
	dst.Radar_theta = strconv.FormatFloat(src.Radar_theta, 'f', -1, 64)
	dst.MPPW = strconv.FormatFloat(src.MPPW, 'f', -1, 64)
	dst.MPPH = strconv.FormatFloat(src.MPPH, 'f', -1, 64)
}
func ConfigStruct2Common_vibrate_setting(dst *Vibrate_setting, src *configStruct.Vibrate_setting) {
	dst.X_vibrate_max = strconv.FormatFloat(src.X_vibrate_max, 'f', -1, 64)
	dst.Y_vibrate_max = strconv.FormatFloat(src.Y_vibrate_max, 'f', -1, 64)
	dst.MatchBoxX = strconv.FormatFloat(src.MatchBoxX, 'f', -1, 64)
	dst.MatchBoxY = strconv.FormatFloat(src.MatchBoxY, 'f', -1, 64)
	dst.MatchBoxWidth = strconv.FormatFloat(src.MatchBoxWidth, 'f', -1, 64)
	dst.MatchBoxHeight = strconv.FormatFloat(src.MatchBoxHeight, 'f', -1, 64)
}
func ConfigStruct2Common_crossing_setting(dst *Crossing_setting, src *configStruct.Crossing_setting) {
	dst.Orientations = strconv.Itoa(src.Orientations)
	dst.Deltax_south = strconv.FormatFloat(src.Deltax_south, 'f', -1, 64)
	dst.Deltay_south = strconv.FormatFloat(src.Deltay_south, 'f', -1, 64)
	dst.Deltax_north = strconv.FormatFloat(src.Deltax_north, 'f', -1, 64)
	dst.Deltay_north = strconv.FormatFloat(src.Deltay_north, 'f', -1, 64)
	dst.Deltax_west = strconv.FormatFloat(src.Deltax_west, 'f', -1, 64)
	dst.Deltay_west = strconv.FormatFloat(src.Deltay_west, 'f', -1, 64)
	dst.Deltax_east = strconv.FormatFloat(src.Deltax_east, 'f', -1, 64)
	dst.Deltay_east = strconv.FormatFloat(src.Deltay_east, 'f', -1, 64)
	dst.Flag_south = strconv.Itoa(src.Flag_south)
	dst.Flag_north = strconv.Itoa(src.Flag_north)
	dst.Flag_west = strconv.Itoa(src.Flag_west)
	dst.Flag_east = strconv.Itoa(src.Flag_east)
	dst.WidthX = strconv.FormatFloat(src.WidthX, 'f', -1, 64)
	dst.WidthY = strconv.FormatFloat(src.WidthY, 'f', -1, 64)
}
func ConfigStruct2Common_real_loc(dst *Real_loc, src *configStruct.Real_loc) {
	dst.Real_left_point_x = strconv.FormatFloat(src.Real_left_point_x, 'f', -1, 64)
	dst.Real_left_point_y = strconv.FormatFloat(src.Real_left_point_y, 'f', -1, 64)
	dst.Real_right_point_x = strconv.FormatFloat(src.Real_right_point_x, 'f', -1, 64)
	dst.Real_right_point_y = strconv.FormatFloat(src.Real_right_point_y, 'f', -1, 64)
	dst.Real_top_point_x = strconv.FormatFloat(src.Real_top_point_x, 'f', -1, 64)
	dst.Real_top_point_y = strconv.FormatFloat(src.Real_top_point_y, 'f', -1, 64)
	dst.Real_bottom_point_x = strconv.FormatFloat(src.Real_bottom_point_x, 'f', -1, 64)
	dst.Real_bottom_point_y = strconv.FormatFloat(src.Real_bottom_point_y, 'f', -1, 64)
}
func ConfigStruct2Common_pixel_loc(dst *Pixel_loc, src *configStruct.Pixel_loc) {
	dst.Pixel_left_point_x = strconv.FormatFloat(src.Pixel_left_point_x, 'f', -1, 64)
	dst.Pixel_left_point_y = strconv.FormatFloat(src.Pixel_left_point_y, 'f', -1, 64)
	dst.Pixel_right_point_x = strconv.FormatFloat(src.Pixel_right_point_x, 'f', -1, 64)
	dst.Pixel_right_point_y = strconv.FormatFloat(src.Pixel_right_point_y, 'f', -1, 64)
	dst.Pixel_top_point_x = strconv.FormatFloat(src.Pixel_top_point_x, 'f', -1, 64)
	dst.Pixel_top_point_y = strconv.FormatFloat(src.Pixel_top_point_y, 'f', -1, 64)
	dst.Pixel_bottom_point_x = strconv.FormatFloat(src.Pixel_bottom_point_x, 'f', -1, 64)
	dst.Pixel_bottom_point_y = strconv.FormatFloat(src.Pixel_bottom_point_y, 'f', -1, 64)
}

func ConfigStruct2Common_info(dst *Info, src *configStruct.Info) {
	ConfigStruct2Common_base(&(dst.Base), &(src.Base))
	ConfigStruct2Common_distance(&(dst.Distance), &(src.Distance))
	ConfigStruct2Common_vibrate_setting(&(dst.Vibrate_setting), &(src.Vibrate_setting))
	ConfigStruct2Common_crossing_setting(&(dst.Crossing_setting), &(src.Crossing_setting))
	ConfigStruct2Common_real_loc(&(dst.Real_loc), &(src.Real_loc))
	ConfigStruct2Common_pixel_loc(&(dst.Pixel_loc), &(src.Pixel_loc))
}

//camera
type Camera struct {
	Flag       string `json:"flag"`
	Ip         string `json:"ip"`
	Url        string `json:"url"`
	Path       string `json:"path"`
	Delay_time string `json:"delay_time"`
}

//cloud
type Cloud struct {
	Ip        string `json:"ip"`
	Port      string `json:"port"`
	Http_ip   string `json:"http_ip"`
	Http_port string `json:"http_port"`
}

//radar
type Radar struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

//annuciator
type Annuciator struct {
	Flag string `json:"flag"`
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

//hardinfo
type HardInfo struct {
	MatrixNo  string `json:"MatrixNo"`
	Hard_code string `json:"hard_code"`
}

//communicate
type Communicate struct {
	Camera     Camera     `json:"camera"`
	Cloud      Cloud      `json:"cloud"`
	Radar      Radar      `json:"radar"`
	Annuciator Annuciator `json:"annuciator"`
	HardInfo   HardInfo   `json:"hardinfo"`
}

func Common2ConfigStruct_camera(dst *configStruct.Camera, src *Camera) {
	dst.Flag, _ = strconv.Atoi(src.Flag)
	dst.Ip = src.Ip
	dst.Url = src.Url
	dst.Path = src.Path
	dst.Delay_time, _ = strconv.Atoi(src.Delay_time)
}

func Common2ConfigStruct_cloud(dst *configStruct.Cloud, src *Cloud) {
	dst.Ip = src.Ip
	dst.Port, _ = strconv.Atoi(src.Port)
	dst.Http_ip = src.Http_ip
	dst.Http_port, _ = strconv.Atoi(src.Http_port)
}

func Common2ConfigStruct_radar(dst *configStruct.Radar, src *Radar) {
	dst.Ip = src.Ip
	dst.Port, _ = strconv.Atoi(src.Port)
}

func Common2ConfigStruct_annuciator(dst *configStruct.Annuciator, src *Annuciator) {
	dst.Flag, _ = strconv.Atoi(src.Flag)
	dst.Ip = src.Ip
	dst.Port, _ = strconv.Atoi(src.Port)
}

func Common2ConfigStruct_hardinfo(dst *configStruct.HardInfo, src *HardInfo) {
	dst.MatrixNo = src.MatrixNo
	dst.Hard_code = src.Hard_code
}

func Common2ConfigStruct_communicate(dst *configStruct.Communicate, src *Communicate) {
	Common2ConfigStruct_camera(&(dst.Camera), &(src.Camera))
	Common2ConfigStruct_cloud(&(dst.Cloud), &(src.Cloud))
	Common2ConfigStruct_radar(&(dst.Radar), &(src.Radar))
	Common2ConfigStruct_annuciator(&(dst.Annuciator), &(src.Annuciator))
	Common2ConfigStruct_hardinfo(&(dst.HardInfo), &(src.HardInfo))
}

func ConfigStruct2Common_camera(dst *Camera, src *configStruct.Camera) {
	dst.Flag = strconv.Itoa(src.Flag)
	dst.Ip = src.Ip
	dst.Url = src.Url
	dst.Path = src.Path
	dst.Delay_time = strconv.Itoa(src.Delay_time)
}

func ConfigStruct2Common_cloud(dst *Cloud, src *configStruct.Cloud) {
	dst.Ip = src.Ip
	dst.Port = strconv.Itoa(src.Port)
	dst.Http_ip = src.Http_ip
	dst.Http_port = strconv.Itoa(src.Http_port)
}

func ConfigStruct2Common_radar(dst *Radar, src *configStruct.Radar) {
	dst.Ip = src.Ip
	dst.Port = strconv.Itoa(src.Port)
}

func ConfigStruct2Common_annuciator(dst *Annuciator, src *configStruct.Annuciator) {
	dst.Flag = strconv.Itoa(src.Flag)
	dst.Ip = src.Ip
	dst.Port = strconv.Itoa(src.Port)
}

func ConfigStruct2Common_hardinfo(dst *HardInfo, src *configStruct.HardInfo) {
	dst.MatrixNo = src.MatrixNo
	dst.Hard_code = src.Hard_code
}

func ConfigStruct2Common_communicate(dst *Communicate, src *configStruct.Communicate) {
	ConfigStruct2Common_camera(&(dst.Camera), &(src.Camera))
	ConfigStruct2Common_cloud(&(dst.Cloud), &(src.Cloud))
	ConfigStruct2Common_radar(&(dst.Radar), &(src.Radar))
	ConfigStruct2Common_annuciator(&(dst.Annuciator), &(src.Annuciator))
	ConfigStruct2Common_hardinfo(&(dst.HardInfo), &(src.HardInfo))
}
