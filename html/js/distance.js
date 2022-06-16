function getConfigInfo() {
    $.ajax({
        url: 'getConfig_info',
        contentType: 'application/json',
        dataType: 'json',
        type: 'get',
        success(res) {
            //base
            form_base.width.value = res.base.width
            form_base.height.value = res.base.height
            //distance
            form_distance.x_distance.value = res.distance.x_distance
            form_distance.y_distance.value = res.distance.y_distance
            form_distance.altitude.value = res.distance.altitude
            // form_distance.y_value.value = res.distance.y_value
            form_distance.coefficient.value = res.distance.coefficient
            form_distance.matrix00.value = res.distance.matrix00
            form_distance.matrix01.value = res.distance.matrix01
            form_distance.matrix02.value = res.distance.matrix02
            form_distance.matrix10.value = res.distance.matrix10
            form_distance.matrix11.value = res.distance.matrix11
            form_distance.matrix12.value = res.distance.matrix12
            form_distance.matrix20.value = res.distance.matrix20
            form_distance.matrix21.value = res.distance.matrix21
            form_distance.matrix22.value = res.distance.matrix22
            form_distance.radar_x.value = res.distance.radar_x
            form_distance.radar_y.value = res.distance.radar_y
            form_distance.radar_theta.value = res.distance.radar_theta
            form_distance.camera_x.value = res.distance.camera_x
            form_distance.camera_y.value = res.distance.camera_y
            form_distance.camera_theta.value = res.distance.camera_theta
            form_distance.MPPW.value = res.distance.MPPW
            form_distance.MPPH.value = res.distance.MPPH
            //vibrate_setting
            form_vibrate_setting.x_vibrate_max.value = res.vibrate_setting.x_vibrate_max
            form_vibrate_setting.y_vibrate_max.value = res.vibrate_setting.y_vibrate_max
            form_vibrate_setting.matchBoxX.value = res.vibrate_setting.matchBoxX
            form_vibrate_setting.matchBoxY.value = res.vibrate_setting.matchBoxY
            form_vibrate_setting.matchBoxWidth.value = res.vibrate_setting.matchBoxWidth
            form_vibrate_setting.matchBoxHeight.value = res.vibrate_setting.matchBoxHeight
            //crossing_setting
            form_crossing_setting.orientations.value = res.crossing_setting.orientations
            form_crossing_setting.deltax_south.value = res.crossing_setting.deltax_south
            form_crossing_setting.deltay_south.value = res.crossing_setting.deltay_south
            form_crossing_setting.deltax_north.value = res.crossing_setting.deltax_north
            form_crossing_setting.deltay_north.value = res.crossing_setting.deltay_north
            form_crossing_setting.deltax_west.value = res.crossing_setting.deltax_west
            form_crossing_setting.deltay_west.value = res.crossing_setting.deltay_west
            form_crossing_setting.deltax_east.value = res.crossing_setting.deltax_east
            form_crossing_setting.deltay_east.value = res.crossing_setting.deltay_east
            form_crossing_setting.flag_south.value = res.crossing_setting.flag_south
            form_crossing_setting.flag_north.value = res.crossing_setting.flag_north
            form_crossing_setting.flag_west.value = res.crossing_setting.flag_west
            form_crossing_setting.flag_east.value = res.crossing_setting.flag_east
            form_crossing_setting.widthX.value = res.crossing_setting.widthX
            form_crossing_setting.widthY.value = res.crossing_setting.widthY
            // //real_loc
            // form_real_loc.real_left_point_x.value = res.real_loc.real_left_point_x
            // form_real_loc.real_left_point_y.value = res.real_loc.real_left_point_y
            // form_real_loc.real_right_point_x.value = res.real_loc.real_right_point_x
            // form_real_loc.real_right_point_y.value = res.real_loc.real_right_point_y
            // form_real_loc.real_top_point_x.value = res.real_loc.real_top_point_x
            // form_real_loc.real_top_point_y.value = res.real_loc.real_top_point_y
            // form_real_loc.real_bottom_point_x.value = res.real_loc.real_bottom_point_x
            // form_real_loc.real_bottom_point_y.value = res.real_loc.real_bottom_point_y
            // //pixel_loc
            // form_pixel_loc.pixel_left_point_x.value = res.pixel_loc.pixel_left_point_x
            // form_pixel_loc.pixel_left_point_y.value = res.pixel_loc.pixel_left_point_y
            // form_pixel_loc.pixel_right_point_x.value = res.pixel_loc.pixel_right_point_x
            // form_pixel_loc.pixel_right_point_y.value = res.pixel_loc.pixel_right_point_y
            // form_pixel_loc.pixel_top_point_x.value = res.pixel_loc.pixel_top_point_x
            // form_pixel_loc.pixel_top_point_y.value = res.pixel_loc.pixel_top_point_y
            // form_pixel_loc.pixel_bottom_point_x.value = res.pixel_loc.pixel_bottom_point_x
            // form_pixel_loc.pixel_bottom_point_y.value = res.pixel_loc.pixel_bottom_point_y
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function setConfigInfo() {
    var urlPath = 'setConfig_info'

    var queryBody = {}

    //base
    var base = {}
    var obj_base = $('#form_base').serializeArray()
    $.each(obj_base, function (i, field) {
        base[field.name] = field.value
    })
    //distance
    var distance = {}
    var obj_distance = $('#form_distance').serializeArray()
    $.each(obj_distance, function (i, field) {
        distance[field.name] = field.value
    })
    //vibrate_setting
    var vibrate_setting = {}
    var obj_vibrate_setting = $('#form_vibrate_setting').serializeArray()
    $.each(obj_vibrate_setting, function (i, field) {
        vibrate_setting[field.name] = field.value
    })
    //crossing_setting
    var crossing_setting = {}
    var obj_crossing_setting = $('#form_crossing_setting').serializeArray()
    $.each(obj_crossing_setting, function (i, field) {
        crossing_setting[field.name] = field.value
    })

    queryBody["base"] = base
    queryBody["distance"] = distance
    queryBody["vibrate_setting"] = vibrate_setting
    queryBody["crossing_setting"] = crossing_setting

    $.ajax({
        url: urlPath,
        contentType: 'application/json',
        type: 'post',
        data: JSON.stringify(queryBody),
        success(res) {
            alert(res)
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function getConfigBase() {
    $.ajax({
        url: 'getConfig_base',
        contentType: 'application/json',
        dataType: 'json',
        type: 'get',
        success(res) {
            form_base.width.value = res.width
            form_base.height.value = res.height
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function getConfigDistance() {
    $.ajax({
        url: 'getConfig_distance',
        contentType: 'application/json',
        dataType: 'json',
        type: 'get',
        success(res) {
            form_distance.x_distance.value = res.x_distance
            form_distance.y_distance.value = res.y_distance
            form_distance.altitude.value = res.altitude
            // form_distance.y_value.value = res.y_value
            form_distance.coefficient.value = res.coefficient
            form_distance.matrix00.value = res.matrix00
            form_distance.matrix01.value = res.matrix01
            form_distance.matrix02.value = res.matrix02
            form_distance.matrix10.value = res.matrix10
            form_distance.matrix11.value = res.matrix11
            form_distance.matrix12.value = res.matrix12
            form_distance.matrix20.value = res.matrix20
            form_distance.matrix21.value = res.matrix21
            form_distance.matrix22.value = res.matrix22
            form_distance.radar_x.value = res.radar_x
            form_distance.radar_y.value = res.radar_y
            form_distance.radar_theta.value = res.radar_theta
            form_distance.camera_x.value = res.camera_x
            form_distance.camera_y.value = res.camera_y
            form_distance.camera_theta.value = res.camera_theta
            form_distance.MPPW.value = res.MPPW
            form_distance.MPPH.value = res.MPPH
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function getConfigVibrate_setting() {
    $.ajax({
        url: 'getConfig_vibrate_setting',
        contentType: 'application/json',
        dataType: 'json',
        type: 'get',
        success(res) {
            form_vibrate_setting.x_vibrate_max.value = res.x_vibrate_max
            form_vibrate_setting.y_vibrate_max.value = res.y_vibrate_max
            form_vibrate_setting.matchBoxX.value = res.matchBoxX
            form_vibrate_setting.matchBoxY.value = res.matchBoxY
            form_vibrate_setting.matchBoxWidth.value = res.matchBoxWidth
            form_vibrate_setting.matchBoxHeight.value = res.matchBoxHeight
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function getConfigCrossing_setting() {
    $.ajax({
        url: 'getConfig_crossing_setting',
        contentType: 'application/json',
        dataType: 'json',
        type: 'get',
        success(res) {
            form_crossing_setting.orientations.value = res.orientations
            form_crossing_setting.deltax_south.value = res.deltax_south
            form_crossing_setting.deltay_south.value = res.deltay_south
            form_crossing_setting.deltax_north.value = res.deltax_north
            form_crossing_setting.deltay_north.value = res.deltay_north
            form_crossing_setting.deltax_west.value = res.deltax_west
            form_crossing_setting.deltay_west.value = res.deltay_west
            form_crossing_setting.deltax_east.value = res.deltax_east
            form_crossing_setting.deltay_east.value = res.deltay_east
            form_crossing_setting.flag_south.value = res.flag_south
            form_crossing_setting.flag_north.value = res.flag_north
            form_crossing_setting.flag_west.value = res.flag_west
            form_crossing_setting.flag_east.value = res.flag_east
            form_crossing_setting.widthX.value = res.widthX
            form_crossing_setting.widthY.value = res.widthY
        },
        error(res) {
            alert(res.responseText)
        }
    })
}


