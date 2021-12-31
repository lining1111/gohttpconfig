package common

//base
type Base struct {
	Width  int `json:"width" ini:"width"`
	Height int `json:"height" ini:"height"`
}

//distance
type Distance struct {
	X_distance  float64 `json:"x_distance" ini:"x_distance"`
	Y_distance  float64 `json:"y_distance" ini:"y_distance"`
	Altitude    float64 `json:"altitude" ini:"altitude"`
	Y_value     float64 `json:"y_value" ini:"y_value"`
	Coefficient float64 `json:"coefficient" ini:"coefficient"`

	Matrix00    float64 `json:"matrix00" ini:"matrix_00"`
	Matrix01    float64 `json:"matrix01" ini:"matrix_01"`
	Matrix02    float64 `json:"matrix02" ini:"matrix_02"`
	Matrix10    float64 `json:"matrix10" ini:"matrix_10"`
	Matrix11    float64 `json:"matrix11" ini:"matrix_11"`
	Matrix12    float64 `json:"matrix12" ini:"matrix_12"`
	Matrix20    float64 `json:"matrix20" ini:"matrix_20"`
	Matrix21    float64 `json:"matrix21" ini:"matrix_21"`
	Matrix22    float64 `json:"matrix22" ini:"matrix_22"`
	Radar_x     float64 `json:"radar_x" ini:"radar_x"`
	Radar_y     float64 `json:"radar_y" ini:"radar_y"`
	Radar_theta float64 `json:"radar_theta" ini:"radar_theta"`
	MPPW        float64 `json:"MPPW" ini:"mppw"`
	MPPH        float64 `json:"MPPH" ini:"mpph"`
}

//vibrate_setting
type Vibrate_setting struct {
	X_vibrate_max  float64 `json:"x_vibrate_max" ini:"x_vibrate_max"`
	Y_vibrate_max  float64 `json:"y_vibrate_max" ini:"y_vibrate_max"`
	MatchBoxX      float64 `json:"matchBoxX" ini:"match_box_x"`
	MatchBoxY      float64 `json:"matchBoxY" ini:"match_box_y"`
	MatchBoxWidth  float64 `json:"matchBoxWidth" ini:"match_box_width"`
	MatchBoxHeight float64 `json:"matchBoxHeight" ini:"match_box_height"`
}

//crossing_setting

type Crossing_setting struct {
	Orientations int     `json:"orientations" ini:"orientations"`
	Deltax_south float64 `json:"deltax_south" ini:"deltax_south"`
	Deltay_south float64 `json:"deltay_south" ini:"deltay_south"`
	Deltax_north float64 `json:"deltax_north" ini:"deltax_north"`
	Deltay_north float64 `json:"deltay_north" ini:"deltay_north"`
	Deltax_west  float64 `json:"deltax_west" ini:"deltax_west"`
	Deltay_west  float64 `json:"deltay_west" ini:"deltay_west"`
	Deltax_east  float64 `json:"deltax_east" ini:"deltax_east"`
	Deltay_east  float64 `json:"deltay_east" ini:"deltay_east"`
	Flag_south   int     `json:"flag_south" ini:"flag_south"`
	Flag_north   int     `json:"flag_north" ini:"flag_north"`
	Flag_west    int     `json:"flag_west" ini:"flag_west"`
	Flag_east    int     `json:"flag_east" ini:"flag_east"`
	WidthX       float64 `json:"widthX" ini:"width_x"`
	WidthY       float64 `json:"widthY" ini:"width_y"`
}

//real_loc
type Real_loc struct {
	Real_left_point_x   float64 `json:"real_left_point_x" ini:"real_left_point_x"`
	Real_left_point_y   float64 `json:"real_left_point_y" ini:"real_left_point_y"`
	Real_right_point_x  float64 `json:"real_right_point_x" ini:"real_right_point_x"`
	Real_right_point_y  float64 `json:"real_right_point_y" ini:"real_right_point_y"`
	Real_top_point_x    float64 `json:"real_top_point_x" ini:"real_top_point_x"`
	Real_top_point_y    float64 `json:"real_top_point_y" ini:"real_top_point_y"`
	Real_bottom_point_x float64 `json:"real_bottom_point_x" ini:"real_bottom_point_x"`
	Real_bottom_point_y float64 `json:"real_bottom_point_y" ini:"real_bottom_point_y"`
}

//pixel_loc
type Pixel_loc struct {
	Pixel_left_point_x   float64 `json:"pixel_left_point_x" ini:"pixel_left_point_x"`
	Pixel_left_point_y   float64 `json:"pixel_left_point_y" ini:"pixel_left_point_y"`
	Pixel_right_point_x  float64 `json:"pixel_right_point_x" ini:"pixel_right_point_x"`
	Pixel_right_point_y  float64 `json:"pixel_right_point_y" ini:"pixel_right_point_y"`
	Pixel_top_point_x    float64 `json:"pixel_top_point_x" ini:"pixel_top_point_x"`
	Pixel_top_point_y    float64 `json:"pixel_top_point_y" ini:"pixel_top_point_y"`
	Pixel_bottom_point_x float64 `json:"pixel_bottom_point_x" ini:"pixel_bottom_point_x"`
	Pixel_bottom_point_y float64 `json:"pixel_bottom_point_y" ini:"pixel_bottom_point_y"`
}

//info
type Info struct {
	Base             Base             `json:"base" ini:"base"`
	Distance         Distance         `json:"distance" ini:"distance"`
	Vibrate_setting  Vibrate_setting  `json:"vibrate_setting" ini:"vibrate_setting"`
	Crossing_setting Crossing_setting `json:"crossing_setting" ini:"crossing_setting"`
	Real_loc         Real_loc         `json:"real_loc" ini:"real_loc"`
	Pixel_loc        Pixel_loc        `json:"pixel_loc" ini:"pixel_loc"`
}

//result
type Result struct {
	Msg string `json:"msg"`
}
