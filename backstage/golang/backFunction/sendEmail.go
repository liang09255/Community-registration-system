package backFunction

import (
	"connectDB"
	"database/sql"
	"fmt"
	"net/http"
	"newFunction"
)

func SendEmail(w http.ResponseWriter, r *http.Request)  {
	session, err := store.Get(r, "loginUser")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	username := session.Values["user"]
	str := username.(string)
	if str != "root" {
		return
	}



	selectPerson := r.FormValue("selectPerson")
	emailContent := r.FormValue("emailContent")

	mailTo := []string{}
	subject := "广药融媒体中心"
	body := emailContent
	var sqlStr string
	if selectPerson == "所有人" {
		sqlStr = "select my_email from sign_up where id > ?"
	}else{
		sqlStr = "select my_email from sign_up where my_status = ?"
	}
	stmt, err := connectDB.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed, err: ", err)
		return
	}
	defer stmt.Close()
	var rows *sql.Rows
	if selectPerson == "所有人" {
		rows, err = stmt.Query(0)
	}else{
		rows, err = stmt.Query(selectPerson)
	}
	if err != nil {
		fmt.Println("scan failed, err: ", err)
		return
	}
	var email string
	for rows.Next(){
		err := rows.Scan(&email)
		if err != nil {
			fmt.Println("scan failed, err: ", err)
			return
		}
		mailTo = append(mailTo, email)
	}

	//邮件测试
	err = newFunction.SendEmain(mailTo, subject, body)
	if err != nil {
		fmt.Println("send email failed, err: ", err)
		w.Write([]byte("邮件发送失败"))
		return
	}else {
		fmt.Println("成功发送邮件")
		w.Write([]byte("邮件发送成功"))
	}
}
