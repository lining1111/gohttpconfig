function getFiles() {
    $.ajax({
        url: 'getFiles',
        dataType: 'html',
        type: 'get',
        success(res) {
            console.log(res)
            $('#container').append(res)
        },
        error(res) {
            // alert(res.responseText)
        }
    })
}
function getFile(name){
    var queryBody = {
        file: name,
    }

    console.log(queryBody)
    // $.ajax({
    //     url: 'getFile',
    //     contentType: 'application/json',
    //     type: 'post',
    //     data: JSON.stringify(queryBody),
    //     success(res) {
    //         alert('下载成功')
    //     },
    //     error(res) {
    //         // alert(res.responseText)
    //     }
    // })
    var url = 'getFile?filename='+name;
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);    // 也可以使用POST方式，根据接口
    xhr.responseType = "blob";  // 返回类型blob
    // 定义请求完成的处理函数，请求前也可以增加加载框/禁用下载按钮逻辑
    xhr.onload = function () {
        // 请求完成
        if (this.status === 200) {
            // 返回200
            var blob = this.response;
            var reader = new FileReader();
            reader.readAsDataURL(blob);  // 转换为base64，可以直接放入a表情href
            reader.onload = function (e) {
                // 转换完成，创建一个a标签用于下载
                var a = document.createElement('a');
                a.download = name;
                a.href = e.target.result;
                $("body").append(a);  // 修复firefox中无法触发click
                a.click();
                $(a).remove();
            }
        }
    };
    // 发送ajax请求
    xhr.send()
}
$(function () {
    getFiles()
})