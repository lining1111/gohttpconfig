package configStruct

/********distanceN1*********/
//base
type Base struct {
	Width  int `ini:"width" db:"width"`
	Height int `ini:"height" db:"height"`
}

//distance
type Distance struct {
	X_distance  float64 `ini:"x_distance" db:"x_distance"`
	Y_distance  float64 `ini:"y_distance" db:"y_distance"`
	Altitude    float64 `ini:"altitude" db:"altitude"`
	Y_value     float64 `ini:"y_value" db:"y_value"`
	Coefficient float64 `ini:"coefficient" db:"coefficient"`

	Matrix00    float64 `ini:"matrix00" db:"matrix00"`
	Matrix01    float64 `ini:"matrix01" db:"matrix01"`
	Matrix02    float64 `ini:"matrix02" db:"matrix02"`
	Matrix10    float64 `ini:"matrix10" db:"matrix10"`
	Matrix11    float64 `ini:"matrix11" db:"matrix11"`
	Matrix12    float64 `ini:"matrix12" db:"matrix12"`
	Matrix20    float64 `ini:"matrix20" db:"matrix20"`
	Matrix21    float64 `ini:"matrix21" db:"matrix21"`
	Matrix22    float64 `ini:"matrix22" db:"matrix22"`
	Radar_x     float64 `ini:"radar_x" db:"radar_x"`
	Radar_y     float64 `ini:"radar_y" db:"radar_y"`
	Radar_theta float64 `ini:"radar_theta" db:"radar_theta"`
	MPPW        float64 `ini:"MPPW" db:"MPPW"`
	MPPH        float64 `ini:"MPPH" db:"MPPH"`
}

//vibrate_setting
type Vibrate_setting struct {
	X_vibrate_max  float64 `ini:"x_vibrate_max" db:"x_vibrate_max"`
	Y_vibrate_max  float64 `ini:"y_vibrate_max" db:"y_vibrate_max"`
	MatchBoxX      float64 `ini:"matchBoxX" db:"matchBoxX"`
	MatchBoxY      float64 `ini:"matchBoxY" db:"matchBoxY"`
	MatchBoxWidth  float64 `ini:"matchBoxWidth" db:"matchBoxWidth"`
	MatchBoxHeight float64 `ini:"matchBoxHeight" db:"matchBoxHeight"`
}

//crossing_setting
type Crossing_setting struct {
	Orientations int     `ini:"orientations" db:"orientations"`
	Deltax_south float64 `ini:"deltax_south" db:"deltax_south"`
	Deltay_south float64 `ini:"deltay_south" db:"deltay_south"`
	Deltax_north float64 `ini:"deltax_north" db:"deltax_north"`
	Deltay_north float64 `ini:"deltay_north" db:"deltay_north"`
	Deltax_west  float64 `ini:"deltax_west" db:"deltax_west"`
	Deltay_west  float64 `ini:"deltay_west" db:"deltay_west"`
	Deltax_east  float64 `ini:"deltax_east" db:"deltax_east"`
	Deltay_east  float64 `ini:"deltay_east" db:"deltay_east"`
	Flag_south   int     `ini:"flag_south" db:"flag_south"`
	Flag_north   int     `ini:"flag_north" db:"flag_north"`
	Flag_west    int     `ini:"flag_west" db:"flag_west"`
	Flag_east    int     `ini:"flag_east" db:"flag_east"`
	WidthX       float64 `ini:"widthX" db:"widthX"`
	WidthY       float64 `ini:"widthY" db:"widthY"`
}

//real_loc
type Real_loc struct {
	Real_left_point_x   float64 `ini:"real_left_point_x" db:"real_left_point_x"`
	Real_left_point_y   float64 `ini:"real_left_point_y" db:"real_left_point_y"`
	Real_right_point_x  float64 `ini:"real_right_point_x" db:"real_right_point_x"`
	Real_right_point_y  float64 `ini:"real_right_point_y" db:"real_right_point_y"`
	Real_top_point_x    float64 `ini:"real_top_point_x" db:"real_top_point_x"`
	Real_top_point_y    float64 `json:"real_top_point_y" ini:"real_top_point_y" db:"real_top_point_y"`
	Real_bottom_point_x float64 `json:"real_bottom_point_x" ini:"real_bottom_point_x" db:"real_bottom_point_x"`
	Real_bottom_point_y float64 `json:"real_bottom_point_y" ini:"real_bottom_point_y" db:"real_bottom_point_y"`
}

//pixel_loc
type Pixel_loc struct {
	Pixel_left_point_x   float64 `ini:"pixel_left_point_x" db:"pixel_left_point_x"`
	Pixel_left_point_y   float64 `ini:"pixel_left_point_y" db:"pixel_left_point_y"`
	Pixel_right_point_x  float64 `ini:"pixel_right_point_x" db:"pixel_right_point_x"`
	Pixel_right_point_y  float64 `ini:"pixel_right_point_y" db:"pixel_right_point_y"`
	Pixel_top_point_x    float64 `ini:"pixel_top_point_x" db:"pixel_top_point_x"`
	Pixel_top_point_y    float64 `ini:"pixel_top_point_y" db:"pixel_top_point_y"`
	Pixel_bottom_point_x float64 `ini:"pixel_bottom_point_x" db:"pixel_bottom_point_x"`
	Pixel_bottom_point_y float64 `ini:"pixel_bottom_point_y" db:"pixel_bottom_point_y"`
}

//info
type Info struct {
	Base             Base             `ini:"base"`
	Distance         Distance         `ini:"distance"`
	Vibrate_setting  Vibrate_setting  `ini:"vibrate_setting"`
	Crossing_setting Crossing_setting `ini:"crossing_setting"`
	//Real_loc         Real_loc         `ini:"real_loc"`
	//Pixel_loc        Pixel_loc        `ini:"pixel_loc"`
}

/************communicate****************/

//camera
type Camera struct {
	Flag       int    `ini:"flag" db:"flag"`
	Ip         string `ini:"ip" db:"ip"`
	Url        string `ini:"url" db:"url"`
	Path       string `ini:"path" db:"path"`
	Delay_time int    `ini:"delay_time" db:"delay_time"`
}

//cloud
type Cloud struct {
	Ip        string `ini:"ip" db:"ip"`
	Port      int    `ini:"port" db:"port"`
	Http_ip   string `ini:"http_ip" db:"http_ip"`
	Http_port int    `ini:"http_port" db:"http_port"`
}

//radar
type Radar struct {
	Ip   string `ini:"ip" db:"ip"`
	Port int    `ini:"port" db:"port"`
}

//annuciator
type Annuciator struct {
	Flag int    `ini:"flag" db:"flag"`
	Ip   string `ini:"ip" db:"ip"`
	Port int    `ini:"port" db:"port"`
}

//hardinfo
type HardInfo struct {
	MatrixNo  string `ini:"MatrixNo" db:"MatrixNo"`
	Hard_code string `ini:"hard_code" db:"hard_code"`
}

//communicate
type Communicate struct {
	Camera     Camera     `ini:"camera"`
	Cloud      Cloud      `ini:"cloud"`
	Radar      Radar      `ini:"radar"`
	Annuciator Annuciator `ini:"annuciator"`
	HardInfo   HardInfo   `ini:"hardinfo"`
}

//NTP
type NTP struct {
	Ip   string `ini:"ip"`
	Port int    `ini:"port"`
}
