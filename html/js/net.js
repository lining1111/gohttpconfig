function setInfoNet() {
    if (!checkIp(eth0Form.ip.value)) {
        alert("eth0 ip 格式不对")
        return;
    }
    if (!checkIp(eth0Form.mask.value)) {
        alert("eth0 mask 格式不对")
        return;
    }

    if (!checkIp(eth1Form.ip.value)) {
        alert("eth1 ip 格式不对")
        return;
    }
    if (!checkIp(eth1Form.mask.value)) {
        alert("eth1 mask 格式不对")
        return;
    }

    if (!checkIp(mainDNS.value)) {
        alert("主DNS ip 格式不对")
        return;
    }

    if (!checkIp(slaveDNS.value)) {
        alert("从DNS ip 格式不对")
        return;
    }

    if (!checkIp(eocCloudIp.value)) {
        alert("eoc cloud ip 格式不对")
        return;
    }

    var eth0 = {}
    var obj_eth0 = $('#eth0Form').serializeArray()
    $.each(obj_eth0, function (i, field) {
        eth0[field.name] = field.value
    })

    var eth1 = {}
    var obj_eth1 = $('#eth1Form').serializeArray()
    $.each(obj_eth1, function (i, field) {
        eth1[field.name] = field.value
    })

    var eoc = {}
    var obj_eoc = $('#eocForm').serializeArray()
    $.each(obj_eoc, function (i, field) {
        eoc[field.name] = field.value
    })

    var queryBody = {}
    queryBody["eth0"] = eth0
    queryBody["eth1"] = eth1
    queryBody["mainDNS"] = mainDNS.value
    queryBody["slaveDNS"] = slaveDNS.value
    queryBody["eoc"] = eoc
    queryBody["city"] = city.value

    $.ajax({
        url: 'setInfoNet',
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

function getInfoNet() {
    $.ajax({
        url: 'getInfoNet',
        contentType: 'application/json',
        dataType: 'json',
        type: 'get',
        success(res) {
            eth0Form.eth0_type.value = res.eth0.type
            eth0Form.eth0_ip.value = res.eth0.ip
            eth0Form.eth0_mask.value = res.eth0.mask
            eth0Form.eth0_gateWay.value = res.eth0.gateWay

            eth1Form.eth1_type.value = res.eth1.type
            eth1Form.eth1_ip.value = res.eth1.ip
            eth1Form.eth1_mask.value = res.eth1.mask
            eth1Form.eth1_gateWay.value = res.eth1.gateWay

            mainDNS.vale = res.mainDNS
            slaveDNS.value = res.slaveDNS

            eocForm.eocCloudIp.value = res.eoc.ip
            eocForm.eocCloudPort.value = res.eoc.port

            city.value = res.city
        },
        error(res) {
            alert(res.responseText)
        }
    })
}


$(function () {
    //网页加载后执行获取
    // getConfigInfo()
    // getConfigCommunicate()
    getInfoNet()
})
