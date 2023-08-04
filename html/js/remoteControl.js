// 相机远程控制
function setInfoCameraRemote() {
    if (!checkIp(cameraForm.ip.value)) {
        alert("cameraForm ip 格式不对")
        return;
    }
    var queryBody = {}
    var obj_camera = $('#cameraForm').serializeArray()
    $.each(obj_camera, function (i, field) {
        queryBody[field.name] = field.value
    })

    $.ajax({
        url: 'setInfoCameraRemote',
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
