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

            //AbnormalStop
            AbnormalStopForm.AbnormalStopControl.value = res.abnormalStop.control
            AbnormalStopForm.AbnormalStopDurationFrom.value = res.abnormalStop.timeFrom
            AbnormalStopForm.AbnormalStopDurationTo.value = res.abnormalStop.timeTo

            //LogControl
            LogControl.value = res.logControl.control

            //OSD
            OSD.value = res.OSD.content
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
            Document.form_camera.flag.value = res.flag
            Document.form_camera.ip.value = res.ip
            Document.form_camera.url.value = res.url
            Document.form_camera.path = res.path
            Document.form_camera.delay_time = res.delay_time
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

//异常停车事件
function setControlAbnormalStop() {
    var queryBody = {}
    var obj_abnormalStop = $('#AbnormalStopForm').serializeArray()
    $.each(obj_abnormalStop, function (i, field) {
        queryBody[field.name] = field.value
    })

    $.ajax({
        url: 'setControlAbnormalStop',
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
//日志打印
function setControlLogControl() {
    var queryBody = {}
    var obj_log = $('#LogControl').serializeArray()
    $.each(obj_log, function (i, field) {
        queryBody[field.name] = field.value
    })

    $.ajax({
        url: 'setControlLogControl',
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
//OSD叠加字符串
function setInfoOSD() {
    var queryBody = {}
    var obj_OSD = $('#OSD').serializeArray()
    $.each(obj_OSD, function (i, field) {
        queryBody[field.name] = field.value
    })

    $.ajax({
        url: 'setInfoOSD',
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