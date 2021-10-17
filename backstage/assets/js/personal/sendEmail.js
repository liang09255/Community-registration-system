//发送邮件
function sendEmail(){
    let selectPerson = $("#selectPerson").val()
    let emailContent = $("#emailContent").val()
    let email_list = "selectPerson=" + selectPerson + "&emailContent=" + emailContent
    let request_email = new XMLHttpRequest()
    request_email.open("POST", "/email?")
    request_email.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    request_email.send(email_list)
    request_email.onreadystatechange = function (){
        if (request_email.readyState === 4 && request_email.status === 200) {
            if (request_email.responseText === "邮件发送成功"){
                toastr.success(request_email.responseText)
            }else{
                toastr.error(request_email.responseText)
            }
        }
    }
}