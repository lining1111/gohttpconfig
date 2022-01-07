package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"gohttpconfig/configStruct"
)

var ConfigDb *sqlx.DB
var IsOpen = false

func Open(path string) error {
	var err error
	ConfigDb, err = sqlx.Open("sqlite3", path)
	if err != nil {
		fmt.Printf("can not open db:%s,err:%v\n", path, err)
	}
	IsOpen = true
	return nil
}

//get

func GetConfig_base(result *configStruct.Base) error {
	sqlCmd := "select * from  base"
	row := ConfigDb.QueryRowx(sqlCmd)
	if row.Err() != nil {
		return row.Err()
	}
	err := row.StructScan(result)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}
func GetConfig_distance(result *configStruct.Distance) error {
	sqlCmd := "select * from  distance"
	row := ConfigDb.QueryRowx(sqlCmd)
	if row.Err() != nil {
		return row.Err()
	}
	err := row.StructScan(result)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}
func GetConfig_vibrate_setting(result *configStruct.Vibrate_setting) error {
	sqlCmd := "select * from  vibrate_setting"
	row := ConfigDb.QueryRowx(sqlCmd)
	if row.Err() != nil {
		return row.Err()
	}
	err := row.StructScan(result)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}
func GetConfig_crossing_setting(result *configStruct.Crossing_setting) error {
	sqlCmd := "select * from  crossing_setting"
	row := ConfigDb.QueryRowx(sqlCmd)
	if row.Err() != nil {
		return row.Err()
	}
	err := row.StructScan(result)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}
func GetConfig_real_loc(result *configStruct.Real_loc) error {
	sqlCmd := "select * from  real_loc"
	row := ConfigDb.QueryRowx(sqlCmd)
	if row.Err() != nil {
		return row.Err()
	}
	err := row.StructScan(result)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}
func GetConfig_pixel_loc(result *configStruct.Pixel_loc) error {
	sqlCmd := "select * from  pixel_loc"
	row := ConfigDb.QueryRowx(sqlCmd)
	if row.Err() != nil {
		return row.Err()
	}
	err := row.StructScan(result)
	if err != nil {
		fmt.Printf(err.Error())
		return err
	}
	return nil
}

func GetConfig_all(result *configStruct.Info) error {
	err := GetConfig_base(&(result.Base))
	if err != nil {
		return err
	}
	err = GetConfig_distance(&(result.Distance))
	if err != nil {
		return err
	}
	err = GetConfig_vibrate_setting(&(result.Vibrate_setting))
	if err != nil {
		return err
	}
	err = GetConfig_crossing_setting(&(result.Crossing_setting))
	if err != nil {
		return err
	}
	err = GetConfig_real_loc(&(result.Real_loc))
	if err != nil {
		return err
	}
	err = GetConfig_pixel_loc(&(result.Pixel_loc))
	if err != nil {
		return err
	}

	return nil
}

//set

func SetConfig_base(result *configStruct.Base) error {
	ConfigDb.Exec("delete from base")

	_, err := ConfigDb.Exec("replace into base("+
		"width,"+
		"height) "+
		"values(?,?)",
		result.Width,
		result.Height)
	return err
}

func SetConfig_distance(result *configStruct.Distance) error {
	ConfigDb.Exec("delete from distance")

	_, err := ConfigDb.Exec("replace into distance("+
		"x_distance,"+
		"y_distance,"+
		"altitude,"+
		"y_value,"+
		"coefficient,"+
		"matrix00,"+
		"matrix01,"+
		"matrix02,"+
		"matrix10,"+
		"matrix11,"+
		"matrix12,"+
		"matrix20,"+
		"matrix21,"+
		"matrix22,"+
		"radar_x,"+
		"radar_y,"+
		"radar_theta,"+
		"MPPW,"+
		"MPPH) "+
		"values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		result.X_distance,
		result.Y_distance,
		result.Altitude,
		result.Y_distance,
		result.Coefficient,
		result.Matrix00,
		result.Matrix01,
		result.Matrix02,
		result.Matrix10,
		result.Matrix11,
		result.Matrix12,
		result.Matrix20,
		result.Matrix21,
		result.Matrix22,
		result.Radar_x,
		result.Radar_y,
		result.Radar_theta,
		result.MPPW,
		result.MPPH)
	return err
}

func SetConfig_vibrate_setting(result *configStruct.Vibrate_setting) error {
	ConfigDb.Exec("delete from vibrate_setting")
	_, err := ConfigDb.Exec("replace into vibrate_setting("+
		"x_vibrate_max,"+
		"y_vibrate_max,"+
		"matchBoxX,"+
		"matchBoxY,"+
		"matchBoxWidth,"+
		"matchBoxHeight) "+
		"values(?,?,?,?,?,?)",
		result.X_vibrate_max,
		result.Y_vibrate_max,
		result.MatchBoxX,
		result.MatchBoxY,
		result.MatchBoxWidth,
		result.MatchBoxHeight)
	return err
}

func SetConfig_crossing_setting(result *configStruct.Crossing_setting) error {
	ConfigDb.Exec("delete from crossing_setting")
	_, err := ConfigDb.Exec("replace into crossing_setting("+
		"orientations,"+
		"deltax_south,"+
		"deltay_south,"+
		"deltax_north,"+
		"deltay_north,"+
		"deltax_west,"+
		"deltay_west,"+
		"deltax_east,"+
		"deltay_east,"+
		"flag_south,"+
		"flag_north,"+
		"flag_west,"+
		"flag_east,"+
		"widthX,"+
		"widthY) "+
		"values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		result.Orientations,
		result.Deltax_south,
		result.Deltay_south,
		result.Deltax_north,
		result.Deltay_north,
		result.Deltax_west,
		result.Deltay_west,
		result.Deltax_east,
		result.Deltay_east,
		result.Flag_south,
		result.Flag_north,
		result.Flag_west,
		result.Flag_east,
		result.WidthX,
		result.WidthY)

	return err
}

func SetConfig_real_loc(result *configStruct.Real_loc) error {
	ConfigDb.Exec("delete from real_loc")
	_, err := ConfigDb.Exec("replace into real_loc("+
		"real_left_point_x,"+
		"real_left_point_y,"+
		"real_right_point_x,"+
		"real_right_point_y,"+
		"real_top_point_x,"+
		"real_top_point_y,"+
		"real_bottom_point_x,"+
		"real_bottom_point_y) "+
		"values(?,?,?,?,?,?,?,?)",
		result.Real_left_point_x,
		result.Real_left_point_y,
		result.Real_right_point_x,
		result.Real_right_point_y,
		result.Real_top_point_x,
		result.Real_top_point_y,
		result.Real_bottom_point_x,
		result.Real_bottom_point_y)

	return err
}

func SetConfig_pixel_loc(result *configStruct.Pixel_loc) error {
	ConfigDb.Exec("delete from pixel_loc")
	_, err := ConfigDb.Exec("replace into pixel_loc("+
		"pixel_left_point_x,"+
		"pixel_left_point_y,"+
		"pixel_right_point_x,"+
		"pixel_right_point_y,"+
		"pixel_top_point_x,"+
		"pixel_top_point_y,"+
		"pixel_bottom_point_x,"+
		"pixel_bottom_point_y) "+
		"values(?,?,?,?,?,?,?,?)",
		result.Pixel_left_point_x,
		result.Pixel_left_point_y,
		result.Pixel_right_point_x,
		result.Pixel_right_point_y,
		result.Pixel_top_point_x,
		result.Pixel_top_point_y,
		result.Pixel_bottom_point_x,
		result.Pixel_bottom_point_y)

	return err
}

func SetConfig_all(result *configStruct.Info) error {
	err := SetConfig_base(&(result.Base))
	if err != nil {
		return err
	}
	err = SetConfig_distance(&(result.Distance))
	if err != nil {
		return err
	}
	err = SetConfig_vibrate_setting(&(result.Vibrate_setting))
	if err != nil {
		return err
	}
	err = SetConfig_crossing_setting(&(result.Crossing_setting))
	if err != nil {
		return err
	}
	err = SetConfig_real_loc(&(result.Real_loc))
	if err != nil {
		return err
	}
	err = SetConfig_pixel_loc(&(result.Pixel_loc))
	if err != nil {
		return err
	}

	return nil
}

func DbTest() {
	Open("./config/config.db")
	var result configStruct.Base
	GetConfig_base(&result)
	result.Width = 4
	SetConfig_base(&result)
}
