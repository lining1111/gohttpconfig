function checkIp(str) {
    const re = /^(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|[0-9])\.((1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)\.){2}(1\d{2}|2[0-4]\d|25[0-5]|[1-9]\d|\d)$/

    return re.test(str)
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


//设置NTP服务器

$(function () {
    //网页加载后执行获取
    // getConfigInfo()
    // getConfigCommunicate()
})