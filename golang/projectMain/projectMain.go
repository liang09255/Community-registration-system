package main

import (
	"backFunction"
	"connectDB"
	"fmt"
	"net/http"
	"newFunction"
)

func main(){
	fmt.Println("程序开始运行")

	connectDB.InitDB()

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("../../"))))

	//提交报名信息
	http.HandleFunc("/submit", newFunction.SubmitValues)

	//获取手机验证码
	http.HandleFunc("/pCode", newFunction.GetPhoneCode)

	//验证手机验证码
	http.HandleFunc("/check", newFunction.CheckCode)

	//查询报名信息
	http.HandleFunc("/search", newFunction.SearchPhone)

	//后台预览报名信息
	http.HandleFunc("/output", backFunction.SeeSignUp)

	//后台导出报名表
	http.HandleFunc("/document", backFunction.GetDocument)
	http.Handle("/get/", http.StripPrefix("/get/", http.FileServer(http.Dir("../../document"))))

	//邮件通知
	http.HandleFunc("/email", backFunction.SendEmail)

	//更改报名状态
	http.HandleFunc("/changeStatus", backFunction.ChangeStatus)

	//登录后台
	http.HandleFunc("/loginBack", backFunction.LoginBack)

	//检查Session
	http.HandleFunc("/checkSession", backFunction.GetSession)

	//手机验证码测试
	//err := newFunction.SendPhoneCode(tea.StringSlice(os.Args[1:]))
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}

	err := http.ListenAndServe(":8888", nil)
	if err != nil{
		fmt.Println(err.Error())
	}
}