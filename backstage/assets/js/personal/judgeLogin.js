$(function (){
    let request_session = new XMLHttpRequest()
    request_session.open("POST", "/checkSession?")
    request_session.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    request_session.send()
    request_session.onreadystatechange = function (){
        if (request_session.readyState === 4 && request_session.status === 200) {
            document.getElementById("accountInfo").innerHTML = request_session.responseText
        }else if (request_session.readyState == 4 && request_session.status !== 200){
            toastr.error("请从登录界面登录后自动跳转到后台")
        }
    }
})