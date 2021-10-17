$(function (){

    let outPutTemplate = "<tr>\n" +
        "                                        <td>{{id}}</td>\n" +
        "                                        <td>{{name}}</td>\n" +
        "                                        <td>{{phone}}</td>\n" +
        "                                        <td>{{email}}</td>\n" +
        "                                        <td>{{part}}</td>\n" +
        "                                        <td><select  id=\"{{id}}\" onchange=\"change({{id}})\">\n" +
        "                                                <optgroup label=\"报名成功\">\n" +
        "                                                    <option>报名成功</option>\n" +
        "                                                </optgroup>\n" +
        "                                                <optgroup label=\"一面\">\n" +
        "                                                    <option>一面成功</option>\n" +
        "                                                    <option>一面失败</option>\n" +
        "                                                </optgroup>\n" +
        "                                                <optgroup label=\"二面\">\n" +
        "                                                    <option>二面成功</option>\n" +
        "                                                    <option>二面失败</option>\n" +
        "                                                </optgroup>\n" +
        "                                                <optgroup label=\"三面\">\n" +
        "                                                    <option>三面成功</option>\n" +
        "                                                    <option>三面失败</option>\n" +
        "                                                </optgroup>\n" +
        "                                            </select></td>\n" +
        "                                    </tr>"

    let request_output = new XMLHttpRequest()
    request_output.open("POST", "/output?")
    request_output.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    request_output.send()
    request_output.onreadystatechange = function (){
        if (request_output.readyState === 4 && request_output.status === 200){
            let data = JSON.parse(request_output.responseText).Result
            let temp = ""
            for (let p = 0; p < data.length; p++) {
                let d = data[p]
                temp = outPutTemplate.replace(/{{id}}/g, d.Id).replace(/{{name}}/g, d.Name).replace(/{{phone}}/g, d.Phone).replace(/{{email}}/g, d.Email).replace(/{{part}}/g, d.Part)
                document.getElementById("outputTbody").insertAdjacentHTML("beforeend",temp)
                let selectButton = $("#" + d.Id)
                selectButton.val(d.Status)
                selectButton.addClass("selectpicker")
                selectButton.selectpicker('refresh')
                selectButton.selectpicker('render')
            }
        }
    }
})

//导出报名表
function getDocument(){
    let request_document = new XMLHttpRequest()
    request_document.open("POST", "/document?")
    request_document.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    request_document.send()
    request_document.onreadystatechange = function (){
        if (request_document.readyState === 4 && request_document.status === 200) {
            window.location.href = request_document.responseText
        }
    }
}

//更改报名状态
function change(id){
    let changeInfo = document.getElementById(id).value
    let request_list = "id=" + id + "&myStatus=" + changeInfo
    let request_change = new XMLHttpRequest()
    request_change.open("POST", "/changeStatus?")
    request_change.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    request_change.send(request_list)
    request_change.onreadystatechange = function (){
        if (request_change.readyState === 4 && request_change.status === 200){
            if (request_change.responseText === "success"){
                toastr.success("成功更改报名状态")
            }
        }
    }
}
