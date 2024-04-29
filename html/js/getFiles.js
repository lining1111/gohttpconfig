function getFiles() {
    $.ajax({
        url: 'getFiles',
        dataType: 'html',
        type: 'get',
        success(res) {
            console.log(res)
            $('#files').append(res)
        },
        error(res) {
           alert(res.responseText)
        }
    })
}

function getFile(name) {
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
    var url = 'getFile?filename=' + name;
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
            reader.readAsDataURL(blob);  // 转换为base64，可以直接放入a标签href
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

/*************上传*****************/
function getFilesInfo() {
    location.href = '../getFiles.html'
}

/****************上传更新**************/

//监听选择文件信息
function fileSelect() {
    console.log("fileSelect")
    var file = document.getElementById('userfile').files[0]
    if (file) {
        var fileSize = 0
        if (file.size > (1024 * 1024)) {
            fileSize = (Math.round(file.size * 100 / (1024 * 1024)) / 100).toString() + 'MB'
        } else {
            fileSize = (Math.round(file.size * 100 / 1024) / 100).toString() + 'KB'
        }
        //set html
        document.getElementById('fileName').innerHTML = 'Name: ' + file.name
        document.getElementById('fileSize').innerHTML = 'Size: ' + fileSize
        document.getElementById('fileType').innerHTML = 'Type: ' + file.type
    }
}


//上传文件
function uploadFile() {
    var formData = new FormData()
    formData.append('updateFile', $("#userfile")[0].files[0])
    // var  formData = new FormData(document.forms.namedItem("uploadForm"))
    $.ajax({
        url: 'update',
        type: 'POST',
        data: formData,
        contentType: false,
        processData: false,
        xhr: function () {
            var xhr = new XMLHttpRequest();
            //使用XMLHttpRequest.upload监听上传过程，注册progress事件，打印回调函数中的event事件
            // xhr.upload.addEventListener('progress', uploadProgress, false)
            xhr.upload.addEventListener('progress', function (e) {
                console.log(e);
                //loaded代表上传了多少
                //total代表总数为多少
                var progressRate = (e.loaded / e.total) * 100 + '%';

                //通过设置进度条的宽度达到效果
                $('.progress > div').css('width', progressRate);
            })
            return xhr;
        },
        success: function (res) {
            alert('上传成功')
        },
        error: function (res) {
            alert('上传失败')
        }
    })
}

$(function () {
    getFiles()
})