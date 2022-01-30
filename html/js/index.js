function checkIp(str) {
    const re = /^(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[0-9])\.((1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)\.){2}(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)$/

    return re.test(str)
}

function getConfigBase() {
    $.ajax({
        url: 'getConfig_base',
        contentType: 'application/json',
        dataType: 'json',
        type: 'get',
        success(res) {
            form_base.width = res.width
            form_base.height = res.height
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

function setConfig(btnID, formID) {
    var urlPath = document.getElementById(btnID).id

    var queryBody = {}
    var obj = $(formID).serializeArray()

    $.each(obj, function (i, field) {
        queryBody[field.name] = field.value
    })

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

//communicate
function setConfigCommunicate() {

    if (!checkIp(form_camera.ip.value)) {
        alert("camera ip 格式不对")
        return;
    }

    if (!checkIp(form_cloud.ip.value)) {
        alert("cloud ip 格式不对")
        return;
    }

    if (!checkIp(form_cloud.http_ip.value)) {
        alert("cloud http_ip 格式不对")
        return;
    }

    if (!checkIp(form_radar.ip.value)) {
        alert("radar ip 格式不对")
        return;
    }

    if (!checkIp(form_annuciator.ip.value)) {
        alert("annuciator ip 格式不对")
        return;
    }

    var queryBody = {}

    //camera
    var camera = {}
    var obj_camera = $('#form_camera').serializeArray()
    $.each(obj_camera, function (i, field) {
        camera[field.name] = field.value
    })
    //cloud
    var cloud = {}
    var obj_cloud = $('#form_cloud').serializeArray()
    $.each(obj_cloud, function (i, field) {
        cloud[field.name] = field.value
    })
    //radar
    var radar = {}
    var obj_radar = $('#form_radar').serializeArray()
    $.each(obj_radar, function (i, field) {
        radar[field.name] = field.value
    })
    //annuciator
    var annuciator = {}
    var obj_annuciator = $('#form_annuciator').serializeArray()
    $.each(obj_annuciator, function (i, field) {
        annuciator[field.name] = field.value
    })
    //hardinfo
    var hardinfo = {}
    var obj_hardinfo = $('#form_hardinfo').serializeArray()
    $.each(obj_hardinfo, function (i, field) {
        hardinfo[field.name] = field.value
    })

    queryBody["camera"] = camera
    queryBody["cloud"] = cloud
    queryBody["radar"] = radar
    queryBody["annuciator"] = annuciator
    queryBody["hardinfo"] = hardinfo

    console.log(queryBody)
    $.ajax({
        url: 'setConfig_communicate',
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

function getConfigCommunicate() {
    $.ajax({
        url: 'getConfig_communicate',
        contentType: 'application/json',
        dataType: 'json',
        type: 'get',
        success(res) {
            //camera
            form_camera.flag.value = res.camera.flag
            form_camera.ip.value = res.camera.ip
            form_camera.url.value = res.camera.url
            form_camera.path.value = res.camera.path
            form_camera.delay_time.value = res.camera.delay_time
            //cloud
            form_cloud.ip.value = res.cloud.ip
            form_cloud.port.value = res.cloud.port
            form_cloud.http_ip.value = res.cloud.http_ip
            form_cloud.http_port.value = res.cloud.http_port

            //radar
            form_radar.ip.value = res.radar.ip
            form_radar.port.value = res.radar.port
            //annuciator
            form_annuciator.flag.value = res.annuciator.flag
            form_annuciator.ip.value = res.annuciator.ip
            form_annuciator.port.value = res.annuciator.port
            //hardinfo
            form_hardinfo.MatrixNo.value = res.hardinfo.MatrixNo
            form_hardinfo.hard_code.value = res.hardinfo.hard_code
        },
        error(res) {
            alert(res.responseText)
        }
    })
}


function setConfigCommunicateCamera() {

    if (!checkIp(form_camera.ip.value)) {
        alert("camera ip 格式不对")
        return;
    }

    var camera = {}
    var obj_camera = $('#form_camera').serializeArray()
    $.each(obj_camera, function (i, field) {
        camera[field.name] = field.value
    })

    $.ajax({
        url: 'setConfig_camera',
        contentType: 'application/json',
        type: 'post',
        data: JSON.stringify(camera),
        success(res) {
            alert(res)
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function getConfigCommunicateCamera() {
    $.ajax({
        url: 'getConfig_camera',
        contentType: 'application/json',
        dataType: 'json',
        type: 'get',
        success(res) {
            form_camera.flag.value = res.flag
            form_camera.ip.value = res.ip
            form_camera.url.value = res.url
            form_camera.path = res.path
            form_camera.delay_time = res.delay_time
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function setConfigCommunicateCloud() {

    if (!checkIp(form_cloud.ip.value)) {
        alert("cloud ip 格式不对")
        return;
    }
    if (!checkIp(form_cloud.http_ip.value)) {
        alert("cloud http_ip 格式不对")
        return;
    }

    var cloud = {}
    var obj_cloud = $('#form_cloud').serializeArray()
    $.each(obj_cloud, function (i, field) {
        cloud[field.name] = field.value
    })

    $.ajax({
        url: 'setConfig_cloud',
        contentType: 'application/json',
        type: 'post',
        data: JSON.stringify(cloud),
        success(res) {
            alert(res)
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function getConfigCommunicateCloud() {
    $.ajax({
        url: 'getConfig_cloud',
        contentType: 'application/json',
        dataType: 'json',
        type: 'get',
        success(res) {
            form_cloud.ip.value = res.ip
            form_cloud.port.value = res.port
            form_cloud.http_ip.value = res.http_ip
            form_cloud.http_port.value = res.http_port
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function setConfigCommunicateRadar() {

    if (!checkIp(form_radar.ip.value)) {
        alert("radar ip 格式不对")
        return;
    }

    var radar = {}
    var obj_radar = $('#form_radar').serializeArray()
    $.each(obj_radar, function (i, field) {
        radar[field.name] = field.value
    })

    $.ajax({
        url: 'setConfig_radar',
        contentType: 'application/json',
        type: 'post',
        data: JSON.stringify(radar),
        success(res) {
            alert(res)
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function getConfigCommunicateRadar() {
    $.ajax({
        url: 'getConfig_radar',
        contentType: 'application/json',
        dataType: 'json',
        type: 'get',
        success(res) {
            form_radar.ip.value = res.ip
            form_radar.port.value = res.port
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function setConfigCommunicateAnnuciator() {

    if (!checkIp(form_annuciator.ip.value)) {
        alert("annuciator ip 格式不对")
        return;
    }

    var annuciator = {}
    var obj_annuciator = $('#form_annuciator').serializeArray()
    $.each(obj_annuciator, function (i, field) {
        annuciator[field.name] = field.value
    })

    $.ajax({
        url: 'setConfig_annuciator',
        contentType: 'application/json',
        type: 'post',
        data: JSON.stringify(annuciator),
        success(res) {
            alert(res)
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function getConfigCommunicateAnnuciator() {
    $.ajax({
        url: 'getConfig_annuciator',
        contentType: 'application/json',
        dataType: 'json',
        type: 'get',
        success(res) {
            form_annuciator.flag.value = res.flag
            form_annuciator.ip.value = res.ip
            form_annuciator.port.value = res.port
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function setConfigCommunicateHardinfo() {

    var hardinfo = {}
    var obj_hardinfo = $('#form_hardinfo').serializeArray()
    $.each(obj_hardinfo, function (i, field) {
        hardinfo[field.name] = field.value
    })

    $.ajax({
        url: 'setConfig_hardinfo',
        contentType: 'application/json',
        type: 'post',
        data: JSON.stringify(hardinfo),
        success(res) {
            alert(res)
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function getConfigCommunicateHardinfo() {
    $.ajax({
        url: 'getConfig_hardinfo',
        contentType: 'application/json',
        dataType: 'json',
        type: 'get',
        success(res) {
            form_hardinfo.MatrixNo.value = res.MatrixNo
            form_hardinfo.hard_code.value = res.hard_code
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function resetProc(proc) {
    var queryBody = {
        reset: '1',
        proc: proc,
    }

    console.log(queryBody)
    $.ajax({
        url: 'resetProc',
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

function getFiles() {
    location.href = '../getFiles.html'
}

$(function () {
    //网页加载后执行获取
    getConfigInfo()
    getConfigCommunicate()
})