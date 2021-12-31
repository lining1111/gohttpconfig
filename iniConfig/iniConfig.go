package iniConfig

//base
type Base struct {
	Width  int `ini:"width"`
	Height int `ini:"height"`
}

//distance
type Distance struct {
	X_distance  float64 `ini:"x_distance"`
	Y_distance  float64 `ini:"y_distance"`
	Altitude    float64 `ini:"altitude"`
	Y_value     float64 `ini:"y_value"`
	Coefficient float64 `ini:"coefficient"`

	Matrix00    float64 `ini:"matrix00"`
	Matrix01    float64 `ini:"matrix01"`
	Matrix02    float64 `ini:"matrix02"`
	Matrix10    float64 `ini:"matrix10"`
	Matrix11    float64 `ini:"matrix11"`
	Matrix12    float64 `ini:"matrix12"`
	Matrix20    float64 `ini:"matrix20"`
	Matrix21    float64 `ini:"matrix21"`
	Matrix22    float64 `ini:"matrix22"`
	Radar_x     float64 `ini:"radar_x"`
	Radar_y     float64 `ini:"radar_y"`
	Radar_theta float64 `ini:"radar_theta"`
	MPPW        float64 `ini:"MPPW"`
	MPPH        float64 `ini:"MPPH"`
}

//vibrate_setting
type Vibrate_setting struct {
	X_vibrate_max  float64 `ini:"x_vibrate_max"`
	Y_vibrate_max  float64 `ini:"y_vibrate_max"`
	MatchBoxX      float64 `ini:"matchBoxX"`
	MatchBoxY      float64 `ini:"matchBoxY"`
	MatchBoxWidth  float64 `ini:"matchBoxWidth"`
	MatchBoxHeight float64 `ini:"matchBoxHeight"`
}

//crossing_setting

type Crossing_setting struct {
	Orientations int     `ini:"orientations"`
	Deltax_south float64 `ini:"deltax_south"`
	Deltay_south float64 `ini:"deltay_south"`
	Deltax_north float64 `ini:"deltax_north"`
	Deltay_north float64 `ini:"deltay_north"`
	Deltax_west  float64 `ini:"deltax_west"`
	Deltay_west  float64 `ini:"deltay_west"`
	Deltax_east  float64 `ini:"deltax_east"`
	Deltay_east  float64 `ini:"deltay_east"`
	Flag_south   int     `ini:"flag_south"`
	Flag_north   int     `ini:"flag_north"`
	Flag_west    int     `ini:"flag_west"`
	Flag_east    int     `ini:"flag_east"`
	WidthX       float64 `ini:"widthX"`
	WidthY       float64 `ini:"widthY"`
}

//real_loc
type Real_loc struct {
	Real_left_point_x   float64 `ini:"real_left_point_x"`
	Real_left_point_y   float64 `ini:"real_left_point_y"`
	Real_right_point_x  float64 `ini:"real_right_point_x"`
	Real_right_point_y  float64 `ini:"real_right_point_y"`
	Real_top_point_x    float64 `ini:"real_top_point_x"`
	Real_top_point_y    float64 `ini:"real_top_point_y"`
	Real_bottom_point_x float64 `ini:"real_bottom_point_x"`
	Real_bottom_point_y float64 `ini:"real_bottom_point_y"`
}

//pixel_loc
type Pixel_loc struct {
	Pixel_left_point_x   float64 `ini:"pixel_left_point_x"`
	Pixel_left_point_y   float64 `ini:"pixel_left_point_y"`
	Pixel_right_point_x  float64 `ini:"pixel_right_point_x"`
	Pixel_right_point_y  float64 `ini:"pixel_right_point_y"`
	Pixel_top_point_x    float64 `ini:"pixel_top_point_x"`
	Pixel_top_point_y    float64 `ini:"pixel_top_point_y"`
	Pixel_bottom_point_x float64 `ini:"pixel_bottom_point_x"`
	Pixel_bottom_point_y float64 `ini:"pixel_bottom_point_y"`
}

//info
type Info struct {
	Base             Base             `ini:"base"`
	Distance         Distance         `ini:"distance"`
	Vibrate_setting  Vibrate_setting  `ini:"vibrate_setting"`
	Crossing_setting Crossing_setting `ini:"crossing_setting"`
	Real_loc         Real_loc         `ini:"real_loc"`
	Pixel_loc        Pixel_loc        `ini:"pixel_loc"`
}
