//查询
function searchPhone(){

    let sMyPhone = document.getElementById("sMyPhone").value

    //加了个后台入口
    if (sMyPhone === "root"){
        let password = prompt("请输入管理员密码")
        let request_root = new XMLHttpRequest()
        let list_root = "password=" + password
        request_root.open("POST", "/loginBack?")
        request_root.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
        request_root.send(list_root)
        request_root.onreadystatechange = function (){
            if (request_root.readyState === 4 && request_root.status === 200) {
                if (request_root.responseText === "/backstage/"){
                    window.location.href = request_root.responseText
                }else{
                    toastr.error("密码错误")
                }
            }
        }
        return
    }

    let list_phone = "sMyPhone=" + sMyPhone
    let request_search = new XMLHttpRequest()
    request_search.open("POST", "/search?")
    request_search.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    request_search.send(list_phone)
    request_search.onreadystatechange = function (){
        if (request_search.readyState === 4 && request_search.status === 200){
            if (request_search.responseText === "查询失败"){
                toastr.error("请先填写报名表", "查询失败")
                document.getElementById("MyName").innerHTML = "未查询到信息"
                document.getElementById("MyEmail").innerHTML = "未查询到信息"
                document.getElementById("MyPart").innerHTML = "未查询到信息"
                document.getElementById("MyAbout").innerHTML = "未查询到信息"
                document.getElementById("MyStatus").innerHTML = "未查询到信息"
                return
            }
            let res = JSON.parse(request_search.responseText)
            let data = res.Result
            document.getElementById("MyName").innerHTML = data.MyName
            document.getElementById("MyEmail").innerHTML = data.MyEmail
            document.getElementById("MyPart").innerHTML = data.MyPart
            document.getElementById("MyAbout").innerHTML = data.MyAbout
            document.getElementById("MyStatus").innerHTML = data.MyStatus
            toastr.success("成功查询到报名信息")
        }
    }
}
