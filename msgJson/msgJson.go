package msgJson

//base
type Base struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

//distance
type Distance struct {
	X_distance  float64 `json:"x_distance"`
	Y_distance  float64 `json:"y_distance"`
	Altitude    float64 `json:"altitude"`
	Y_value     float64 `json:"y_value"`
	Coefficient float64 `json:"coefficient"`

	Matrix00    float64 `json:"matrix00"`
	Matrix01    float64 `json:"matrix01"`
	Matrix02    float64 `json:"matrix02"`
	Matrix10    float64 `json:"matrix10"`
	Matrix11    float64 `json:"matrix11"`
	Matrix12    float64 `json:"matrix12"`
	Matrix20    float64 `json:"matrix20"`
	Matrix21    float64 `json:"matrix21"`
	Matrix22    float64 `json:"matrix22"`
	Radar_x     float64 `json:"radar_x"`
	Radar_y     float64 `json:"radar_y"`
	Radar_theta float64 `json:"radar_theta"`
	MPPW        float64 `json:"MPPW"`
	MPPH        float64 `json:"MPPH"`
}

//vibrate_setting
type Vibrate_setting struct {
	X_vibrate_max  float64 `json:"x_vibrate_max"`
	Y_vibrate_max  float64 `json:"y_vibrate_max"`
	MatchBoxX      float64 `json:"matchBoxX"`
	MatchBoxY      float64 `json:"matchBoxY"`
	MatchBoxWidth  float64 `json:"matchBoxWidth"`
	MatchBoxHeight float64 `json:"matchBoxHeight"`
}

//crossing_setting

type Crossing_setting struct {
	Orientations int     `json:"orientations"`
	Deltax_south float64 `json:"deltax_south"`
	Deltay_south float64 `json:"deltay_south"`
	Deltax_north float64 `json:"deltax_north"`
	Deltay_north float64 `json:"deltay_north"`
	Deltax_west  float64 `json:"deltax_west"`
	Deltay_west  float64 `json:"deltay_west"`
	Deltax_east  float64 `json:"deltax_east"`
	Deltay_east  float64 `json:"deltay_east"`
	Flag_south   int     `json:"flag_south"`
	Flag_north   int     `json:"flag_north"`
	Flag_west    int     `json:"flag_west"`
	Flag_east    int     `json:"flag_east"`
	WidthX       float64 `json:"widthX"`
	WidthY       float64 `json:"widthY"`
}

//real_loc
type Real_loc struct {
	Real_left_point_x   float64 `json:"real_left_point_x"`
	Real_left_point_y   float64 `json:"real_left_point_y"`
	Real_right_point_x  float64 `json:"real_right_point_x"`
	Real_right_point_y  float64 `json:"real_right_point_y"`
	Real_top_point_x    float64 `json:"real_top_point_x"`
	Real_top_point_y    float64 `json:"real_top_point_y"`
	Real_bottom_point_x float64 `json:"real_bottom_point_x"`
	Real_bottom_point_y float64 `json:"real_bottom_point_y"`
}

//pixel_loc
type Pixel_loc struct {
	Pixel_left_point_x   float64 `json:"pixel_left_point_x"`
	Pixel_left_point_y   float64 `json:"pixel_left_point_y"`
	Pixel_right_point_x  float64 `json:"pixel_right_point_x"`
	Pixel_right_point_y  float64 `json:"pixel_right_point_y"`
	Pixel_top_point_x    float64 `json:"pixel_top_point_x"`
	Pixel_top_point_y    float64 `json:"pixel_top_point_y"`
	Pixel_bottom_point_x float64 `json:"pixel_bottom_point_x"`
	Pixel_bottom_point_y float64 `json:"pixel_bottom_point_y"`
}

//result
type Result struct {
	Msg string `json:"msg"`
}
