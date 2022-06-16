/*************杀死进程*******************/
function killProc(proc) {
    var queryBody = {
        reset: '1',
        proc: proc,
    }

    console.log(queryBody)
    $.ajax({
        url: 'killProc',
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

//重启服务器
function resetServer() {
    $.ajax({
        url: 'resetServer',
        type: 'POST',
        success: function (res) {
            // alert('重启成功')
        },
        error: function (res) {
            // alert('重启失败')
        }
    })
}
