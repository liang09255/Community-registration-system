/*  梳理一下逻辑
    实时监控的东西：
    1.手机号，正则匹配正确就开放验证码按钮，错误就重新给他关闭
    2.邮箱正则匹配
    按提交按钮后要判断的东西：
    1.手机号是否已验证成功，这里通过获取手机号框框是否已锁定来判断
    2.邮箱正则是否匹配正确，再check一遍邮箱，错误就弹窗
    3.是否有内容为空
*/
let myName = document.getElementById("myName")
let myPhone = document.getElementById("myPhone")
let myEmail = document.getElementById("myEmail")
let myPart = document.getElementById("myPart")
let myAbout = document.getElementById("myAbout")
let myCode = document.getElementById("myCode")

function submitSignUp(){
    let smyName = myName.value
    let smyPhone = myPhone.value
    let smyEmail = myEmail.value
    let smyPart = myPart.value
    let smyAbout = myAbout.value

    if (smyName === "" || smyPhone === "" || smyEmail === "" || smyPart === "" || smyAbout === ""){
        toastr.warning("请填写完整所有信息", "提交失败")
        return
    }
    if ($("#myPhone").prop("disabled") !== true){
        toastr.warning("验证码有误，请重新获取验证码", "提交失败")
        return;
    }

    //留个位置补一下判断邮箱
    if (getEmail() === false){
        return;
    }

    //7.16改Bug记录：用户输入内容包含分号（;）时会被截断，转换为url编码解决
    smyAbout = smyAbout.replace(/;/g, "%3b")      //全局匹配模式

    let list_submit = "myName=" + smyName + "&myPhone=" + smyPhone + "&myEmail=" + smyEmail + "&myPart=" + smyPart + "&myAbout=" + smyAbout + "&myStatus=" + "报名成功"

    let request_submit = new XMLHttpRequest()
    request_submit.open("POST", "/submit?")
    request_submit.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
    request_submit.send(list_submit)
    request_submit.onreadystatechange = function (){
        if (request_submit.readyState === 4 && request_submit.status === 200){
            if (request_submit.responseText === "success"){
                toastr.success("报名信息提交成功")
            }else if (request_submit.responseText === "update"){
                toastr.success("报名信息更新成功")
            }
        }
    }

}

//获取验证码
$(document).ready(function(){
    let orderTime=60   //设置再次发送验证码等待时间
    let timeLeft=orderTime
    let btn=$("#codeButton")
    let phone=$("#myPhone")
    //电话号码匹配正则
    let reg = /^1(?:3\d|4[4-9]|5[0-35-9]|6[67]|7[013-8]|8\d|9\d)\d{8}$/

    phone.keyup(function(){
        if (reg.test(phone.val())){
            btn.removeAttr("disabled")  //当号码符合规则后发送验证码按钮可点击
            //移除disabledBtn类
            let classVal = document.getElementById("codeButton").getAttribute("class")
            classVal = classVal.replace("disabledBtn", "")
            document.getElementById("codeButton").setAttribute("class", classVal)
        }
        else{
            btn.attr("disabled",true)
            let reg2 = /disabledBtn/
            let classVal = document.getElementById("codeButton").getAttribute("class")
            if (!reg2.test(classVal)){
                classVal = classVal.concat("disabledBtn")
                document.getElementById("codeButton").setAttribute("class", classVal)
            }
            if (phone.val().length === 11){
                toastr.warning("请输入正确的手机号")
            }
        }
    })

    //计时函数
    function timeCount(){
        if (phone.prop("disabled") === true){
            btn.val("验证成功")
            btn.css("background", "#198754")
            return
        }
        timeLeft-=1
        if (timeLeft>0){
            btn.val(timeLeft+" 秒后重发");
            setTimeout(timeCount,1000)
        }
        else {
            btn.val("重新发送");
            timeLeft=orderTime   //重置等待时间
            btn.removeAttr("disabled");
        }
    }

    //事件处理函数
    btn.on("click",function(){
        $(this).attr("disabled",true); //防止多次点击
        //此处可添加 ajax请求 向后台发送 获取验证码请求
        timeCount(this);

        //获取验证码
        let myPhoneNumber = document.getElementById("myPhone").value
        let list = "myPhone=" + myPhoneNumber
        let request_getCode = new XMLHttpRequest()
        request_getCode.open("POST", "/pCode?")
        request_getCode.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
        request_getCode.send(list)
        request_getCode.onreadystatechange = function (){
            if (request_getCode.readyState === 4 && request_getCode.status === 200){
                //这里是阿里云那个验证码限制，随便返回了个字符串来判断
                if (request_getCode.responseText === "so fast"){
                    toastr.error("请过段时间再试", "获取过于频繁")
                }
            }
        }
    })
})

//这里写实时监听的那两个东西
$(function(){
    //邮箱验证，只做了下正则
    myEmail.addEventListener("change", getEmail, false)
    //手机验证
    myCode.addEventListener("change", getWord, false)
    function getWord(){
        let myCodeNumber = myCode.value
        let myPhoneNumber = myPhone.value
        //六位数发送请求
        if (myCodeNumber.length === 6){
            let list_check = "myCodeNumber=" + myCodeNumber + "&myPhoneNumber=" + myPhoneNumber
            let request_check = new XMLHttpRequest()
            request_check.open("POST", "/check?")
            request_check.setRequestHeader("Content-Type", "application/x-www-form-urlencoded")
            request_check.send(list_check)
            request_check.onreadystatechange = function (){
                if (request_check.readyState === 4 && request_check.status === 200){
                    let rText = request_check.responseText
                    let phone = $("#myPhone")
                    let code = $("#myCode")
                    let codeButton = $("#codeButton")
                    //把三个框 disabled了，然后改绿
                    if (rText === "验证成功"){
                        phone.attr("disabled", true)
                        code.attr("disabled", true)
                        codeButton.attr("disabled", true)
                        codeButton.val("验证成功")
                        codeButton.css("background", "#198754")
                        toastr.success("输入框已锁定", "手机号验证成功")
                    }else{
                        toastr.error("请检查验证码是否填写错误", "手机号验证失败")
                    }
                }
            }
        }
    }
})
//邮箱验证，正确返回true，错误false
function getEmail() {
    let email = myEmail.value
    let regular = /^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$/
    if (regular.test(email)){
        return true
    }else{
        toastr.warning("请输入正确的邮箱")
        return false
    }
}
