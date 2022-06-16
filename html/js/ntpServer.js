function setInfoNTP() {
    if (!checkIp(ntpForm.ip.value)) {
        alert("ntp ip 格式不对")
        return;
    }

    var ntp = {}
    var obj_ntp = $('#ntpForm').serializeArray()
    $.each(obj_ntp, function (i, field) {
        ntp[field.name] = field.value
    })

    $.ajax({
        url: 'setInfoNTP',
        contentType: 'application/json',
        type: 'post',
        data: JSON.stringify(ntp),
        success(res) {
            alert(res)
        },
        error(res) {
            alert(res.responseText)
        }
    })
}

function getInfoNTP() {
    $.ajax({
        url: 'getInfoNTP',
        contentType: 'application/json',
        dataType: 'json',
        type: 'get',
        success(res) {
            ntpForm.ntp_ip.value = res.ip
            ntpForm.ntp_port.value = res.port

        },
        error(res) {
            alert(res.responseText)
        }
    })
}

$(function () {
    getInfoNTP()
})