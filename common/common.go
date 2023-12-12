package common

/********************distanceN1****************************/
//base
type Base struct {
	Width  int `json:"width,string" ini:"width" db:"width"`
	Height int `json:"height,string" ini:"height" db:"height"`
}

//distance
type Distance struct {
	X_distance  float64 `json:"x_distance,string" ini:"x_distance" db:"x_distance"`
	Y_distance  float64 `json:"y_distance,string" ini:"y_distance" db:"y_distance"`
	Altitude    float64 `json:"altitude,string" ini:"altitude" db:"altitude"`
	Y_value     float64 `json:"y_value,string,omitempty" ini:"y_value" db:"y_value"`
	Coefficient float64 `json:"coefficient,string" ini:"coefficient" db:"coefficient"`

	Matrix00     float64 `json:"matrix00,string" ini:"matrix00" db:"matrix00"`
	Matrix01     float64 `json:"matrix01,string" ini:"matrix01" db:"matrix01"`
	Matrix02     float64 `json:"matrix02,string" ini:"matrix02" db:"matrix02"`
	Matrix10     float64 `json:"matrix10,string" ini:"matrix10" db:"matrix10"`
	Matrix11     float64 `json:"matrix11,string" ini:"matrix11" db:"matrix11"`
	Matrix12     float64 `json:"matrix12,string" ini:"matrix12" db:"matrix12"`
	Matrix20     float64 `json:"matrix20,string" ini:"matrix20" db:"matrix20"`
	Matrix21     float64 `json:"matrix21,string" ini:"matrix21" db:"matrix21"`
	Matrix22     float64 `json:"matrix22,string" ini:"matrix22" db:"matrix22"`
	Radar_x      float64 `json:"radar_x,string" ini:"radar_x" db:"radar_x"`
	Radar_y      float64 `json:"radar_y,string" ini:"radar_y" db:"radar_y"`
	Radar_theta  float64 `json:"radar_theta,string" ini:"radar_theta" db:"radar_theta"`
	Camera_x     float64 `json:"camera_x,string" ini:"camera_x" db:"camera_x"`
	Camera_y     float64 `json:"camera_y,string" ini:"camera_y" db:"camera_y"`
	Camera_theta float64 `json:"camera_theta,string" ini:"camera_theta" db:"camera_theta"`
	MPPW         float64 `json:"MPPW,string" ini:"MPPW" db:"MPPW"`
	MPPH         float64 `json:"MPPH,string" ini:"MPPH" db:"MPPH"`
}

//vibrate_setting
type Vibrate_setting struct {
	X_vibrate_max  float64 `json:"x_vibrate_max,string" ini:"x_vibrate_max" db:"x_vibrate_max"`
	Y_vibrate_max  float64 `json:"y_vibrate_max,string" ini:"y_vibrate_max" db:"y_vibrate_max"`
	MatchBoxX      float64 `json:"matchBoxX,string" ini:"matchBoxX" db:"matchBoxX"`
	MatchBoxY      float64 `json:"matchBoxY,string" ini:"matchBoxY" db:"matchBoxY"`
	MatchBoxWidth  float64 `json:"matchBoxWidth,string" ini:"matchBoxWidth" db:"matchBoxWidth"`
	MatchBoxHeight float64 `json:"matchBoxHeight,string" ini:"matchBoxHeight" db:"matchBoxHeight"`
}

//crossing_setting
type Crossing_setting struct {
	Orientations int     `json:"orientations,string" ini:"orientations" db:"orientations"`
	Deltax_south float64 `json:"deltax_south,string" ini:"deltax_south" db:"deltax_south"`
	Deltay_south float64 `json:"deltay_south,string" ini:"deltay_south" db:"deltay_south"`
	Deltax_north float64 `json:"deltax_north,string" ini:"deltax_north" db:"deltax_north"`
	Deltay_north float64 `json:"deltay_north,string" ini:"deltay_north" db:"deltay_north"`
	Deltax_west  float64 `json:"deltax_west,string" ini:"deltax_west" db:"deltax_west"`
	Deltay_west  float64 `json:"deltay_west,string" ini:"deltay_west" db:"deltay_west"`
	Deltax_east  float64 `json:"deltax_east,string" ini:"deltax_east" db:"deltax_east"`
	Deltay_east  float64 `json:"deltay_east,string" ini:"deltay_east" db:"deltay_east"`
	Flag_south   int     `json:"flag_south,string" ini:"flag_south" db:"flag_south"`
	Flag_north   int     `json:"flag_north,string" ini:"flag_north" db:"flag_north"`
	Flag_west    int     `json:"flag_west,string" ini:"flag_west" db:"flag_west"`
	Flag_east    int     `json:"flag_east,string" ini:"flag_east" db:"flag_east"`
	WidthX       float64 `json:"widthX,string" ini:"widthX" db:"widthX"`
	WidthY       float64 `json:"widthY,string" ini:"widthY" db:"widthY"`
}

//real_loc
type Real_loc struct {
	Real_left_point_x   float64 `json:"real_left_point_x,string" ini:"real_left_point_x" db:"real_left_point_x"`
	Real_left_point_y   float64 `json:"real_left_point_y,string" ini:"real_left_point_y" db:"real_left_point_y"`
	Real_right_point_x  float64 `json:"real_right_point_x,string" ini:"real_right_point_x" db:"real_right_point_x"`
	Real_right_point_y  float64 `json:"real_right_point_y,string" ini:"real_right_point_y" db:"real_right_point_y"`
	Real_top_point_x    float64 `json:"real_top_point_x,string" ini:"real_top_point_x" db:"real_top_point_x"`
	Real_top_point_y    float64 `json:"real_top_point_y,string" ini:"real_top_point_y" db:"real_top_point_y"`
	Real_bottom_point_x float64 `json:"real_bottom_point_x,string" ini:"real_bottom_point_x" db:"real_bottom_point_x"`
	Real_bottom_point_y float64 `json:"real_bottom_point_y,string" ini:"real_bottom_point_y" db:"real_bottom_point_y"`
}

//pixel_loc
type Pixel_loc struct {
	Pixel_left_point_x   float64 `json:"pixel_left_point_x,string" ini:"pixel_left_point_x" db:"pixel_left_point_x"`
	Pixel_left_point_y   float64 `json:"pixel_left_point_y,string" ini:"pixel_left_point_y" db:"pixel_left_point_y"`
	Pixel_right_point_x  float64 `json:"pixel_right_point_x,string" ini:"pixel_right_point_x" db:"pixel_right_point_x"`
	Pixel_right_point_y  float64 `json:"pixel_right_point_y,string" ini:"pixel_right_point_y" db:"pixel_right_point_y"`
	Pixel_top_point_x    float64 `json:"pixel_top_point_x,string" ini:"pixel_top_point_x" db:"pixel_top_point_x"`
	Pixel_top_point_y    float64 `json:"pixel_top_point_y,string" ini:"pixel_top_point_y" db:"pixel_top_point_y"`
	Pixel_bottom_point_x float64 `json:"pixel_bottom_point_x,string" ini:"pixel_bottom_point_x" db:"pixel_bottom_point_x"`
	Pixel_bottom_point_y float64 `json:"pixel_bottom_point_y,string" ini:"pixel_bottom_point_y" db:"pixel_bottom_point_y"`
}

//info
type Info struct {
	Base             Base             `json:"base" ini:"base"`
	Distance         Distance         `json:"distance" ini:"distance"`
	Vibrate_setting  Vibrate_setting  `json:"vibrate_setting" ini:"vibrate_setting"`
	Crossing_setting Crossing_setting `json:"crossing_setting" ini:"crossing_setting"`
	LaneAssociation  LaneAssociation  `json:"laneAssociation" ini:"laneAssociation"`
	//Real_loc         Real_loc         `json:"real_loc"`
	//Pixel_loc        Pixel_loc        `json:"pixel_loc"`
}

/*******************communicate**************************/
//camera
type Camera struct {
	Flag       int    `json:"flag,string" ini:"flag" db:"flag"`
	Ip         string `json:"ip" ini:"ip" db:"ip"`
	Url        string `json:"url" ini:"url" db:"url"`
	Path       string `json:"path" ini:"path" db:"path"`
	Delay_time string `json:"delay_time" ini:"delay_time" db:"delay_time"`
}

//cloud
type Cloud struct {
	Ip        string `json:"ip" ini:"ip" db:"ip"`
	Port      int    `json:"port,string" ini:"port" db:"port"`
	Http_ip   string `json:"http_ip" ini:"http_ip" db:"http_ip"`
	Http_port int    `json:"http_port,string" ini:"http_port" db:"http_port"`
}

//radar
type Radar struct {
	Ip   string `json:"ip" ini:"ip" db:"ip"`
	Port int    `json:"port,string" ini:"port" db:"port"`
}

//annuciator
type Annuciator struct {
	Flag int    `json:"flag,string" ini:"flag" db:"flag"`
	Ip   string `json:"ip" ini:"ip" db:"ip"`
	Port int    `json:"port,string" ini:"port" db:"port"`
}

//hardinfo
type HardInfo struct {
	MatrixNo  string `json:"MatrixNo" ini:"MatrixNo" db:"MatrixNo"`
	Hard_code string `json:"hard_code" ini:"hard_code" db:"hard_code"`
}

//communicate
type Communicate struct {
	Camera        Camera        `json:"camera" ini:"camera"`
	Cloud         Cloud         `json:"cloud" ini:"cloud"`
	Radar         Radar         `json:"radar" ini:"radar"`
	Annuciator    Annuciator    `json:"annuciator" ini:"annuciator"`
	HardInfo      HardInfo      `json:"hardinfo" ini:"hardinfo"`
	AbnormalStop  AbnormalStop  `json:"abnormalStop" ini:"abnormalStop"`
	LogControl    LogControl    `json:"logControl" ini:"logControl"`
	OSD           OSD           `json:"OSD" ini:"OSD"`
	JamParam      JamParam      `json:"jamParam" ini:"jamParam"`
	OverFlowParam OverFlowParam `json:"overFlowParam" ini:"overFlowParam"`
	Debug         Debug         `json:"debug" ini:"debug"`
	Wfzp          Wfzp          `json:"wfzp" ini:"wfzp"`
}

/***************reset proc*********************/
type ResetProc struct {
	Reset string `json:"reset"`
	Proc  string `json:"proc"`
}

/***************setInfoNTP*********************/
type NTP struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

/***************setInfoNet*********************/
type Eth struct {
	Type    string `json:"type"`
	Ip      string `json:"ip"`
	Mask    string `json:"mask"`
	GateWay string `json:"gateWay"`
}

type Eoc struct {
	Ip   string `json:"ip"`
	Port string `json:"port"`
}

type Net struct {
	Eth0     Eth    `json:"eth0"`
	Eth1     Eth    `json:"eth1"`
	MainDNS  string `json:"mainDNS"`
	SlaveDNS string `json:"slaveDNS"`
	Eoc      Eoc    `json:"eoc"`
	City     string `json:"city"`
}

/********************setInfoCameraRemote******************************/
type CameraRemote struct {
	Ip   string `json:"ip"`
	Time int    `json:"time,string"`
}

/*******************setControlAbnormalStop***************************/
type AbnormalStop struct {
	Control  int    `json:"control,string" ini:"control"`
	TimeFrom string `json:"timeFrom" ini:"timeFrom"`
	TimeTo   string `json:"timeTo" ini:"timeTo"`
}

/*******************setControlLog***************************/
type LogControl struct {
	Control int `json:"control,string" ini:"control"`
}

/*******************setInfoOSD***************************/
type OSD struct {
	Content string `json:"content" ini:"content"`
}

/*******************JamParam**************************/
type JamParam struct {
	AreaVehicle    int     `json:"areaVehicle,string" ini:"areaVehicle"`
	VehicleSpeed   float64 `json:"vehicleSpeed,string" ini:"vehicleSpeed"`
	DetectDuration int     `json:"detectDuration,string" ini:"detectDuration"`
}

/*******************OverFlowParam**************************/
type OverFlowParam struct {
	StopVehicle    int `json:"stopVehicle,string" ini:"stopVehicle"`
	DetectDuration int `json:"detectDuration,string" ini:"detectDuration"`
}

type Debug struct {
	Print_json             int `json:"print_json,string" ini:"print_json"`
	DrawCoilArea           int `json:"drawCoilArea,string" ini:"drawCoilArea"`
	DrawPassStopLine       int `json:"drawPassStopLine,string" ini:"drawPassStopLine"`
	DrawPass42mLine        int `json:"drawPass42mLine,string" ini:"drawPass42mLine"`
	DrawPass60mLine        int `json:"drawPass60mLine,string" ini:"drawPass60mLine"`
	DrawAbnormalStop       int `json:"drawAbnormalStop,string" ini:"drawAbnormalStop"`
	DrawCrossingCongestion int `json:"drawCrossingCongestion,string" ini:"drawCrossingCongestion"`
	DrawOverFlow           int `json:"drawOverFlow,string" ini:"drawOverFlow"`
	DrawLaneLength         int `json:"drawLaneLength,string" ini:"drawLaneLength"`
	DrawTestInfo           int `json:"drawTestInfo,string" ini:"drawTestInfo"`
}

type Wfzp struct {
	UseDjData             int `json:"useDjData,string" ini:"useDjData"`
	JudgeRunRedLight      int `json:"judgeRunRedLight,string" ini:"judgeRunRedLight"`
	JudgeNotFollowLane    int `json:"judgeNotFollowLane,string" ini:"judgeNotFollowLane"`
	JudgeOverLane         int `json:"judgeOverLane,string" ini:"judgeOverLane"`
	JudgeStayNonmotorArea int `json:"judgeStayNonmotorArea,string" ini:"judgeStayNonmotorArea"`
	JudgeWrongDirection   int `json:"judgeWrongDirection,string" ini:"judgeWrongDirection"`
}

type LaneAssociation struct {
	LaneCode1  string `json:"laneCode1" ini:"laneCode1" db:"laneCode1"`
	LaneCode2  string `json:"laneCode2" ini:"laneCode2" db:"laneCode2"`
	LaneCode3  string `json:"laneCode3" ini:"laneCode3" db:"laneCode3"`
	LaneCode4  string `json:"laneCode4" ini:"laneCode4" db:"laneCode4"`
	LaneCode5  string `json:"laneCode5" ini:"laneCode5" db:"laneCode5"`
	LaneCode6  string `json:"laneCode6" ini:"laneCode6" db:"laneCode6"`
	LaneCode7  string `json:"laneCode7" ini:"laneCode7" db:"laneCode7"`
	LaneCode8  string `json:"laneCode8" ini:"laneCode8" db:"laneCode8"`
	LaneCode9  string `json:"laneCode9" ini:"laneCode9" db:"laneCode9"`
	LaneCode10 string `json:"laneCode10" ini:"laneCode10" db:"laneCode10"`
	LaneCode11 string `json:"laneCode11" ini:"laneCode11" db:"laneCode11"`
	LaneCode12 string `json:"laneCode12" ini:"laneCode12" db:"laneCode12"`
}
