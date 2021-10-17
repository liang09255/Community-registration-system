package newFunction

import (
	"connectDB"
	"fmt"
	"net/http"
)

type Information struct {
	MyName string
	MyPhone string
	MyEmail string
	MyPart string
	MyAbout string
	MyStatus string
}

func NewInformation(myName string, myPhone string, myEmail string, myPart string, myAbout string, myStatus string) *Information{
	return &Information{
		MyName: myName,
		MyPhone: myPhone,
		MyEmail: myEmail,
		MyPart: myPart,
		MyAbout: myAbout,
		MyStatus: myStatus,
	}
}

func SubmitValues(w http.ResponseWriter, r *http.Request)  {
	//获取用户提交的报名信息
	var user Information
	user.MyName = r.FormValue("myName")
	user.MyPhone = r.FormValue("myPhone")
	user.MyEmail = r.FormValue("myEmail")
	user.MyPart = r.FormValue("myPart")
	user.MyAbout = r.FormValue("myAbout")
	user.MyStatus = r.FormValue("myStatus")

	sqlStr0 := "select my_phone from sign_up where id > ?"
	stmt0, err := connectDB.DB.Prepare(sqlStr0)
	if err != nil {
		fmt.Println("prepare failed, err: ", err)
		return
	}
	defer stmt0.Close()
	rows, err := stmt0.Query(0)
	if err != nil {
		fmt.Println("scan failed, err: ", err)
		return
	}
	judgeExistence := 0
	for rows.Next() {
		 sPhone := ""
		 err := rows.Scan(&sPhone)
		 if err != nil {
		 	fmt.Println("scan failed, err: ", err)
			 return
		 }
		 if sPhone == user.MyPhone{
		 	judgeExistence = 1
		 	break
		 }
	}
	if judgeExistence == 1 {
		sqlStr1 := "update sign_up set my_name = ?, my_phone = ?, my_email = ?, my_part = ?, my_about = ? , my_status = ? where my_phone = ?"
		stmt1, err := connectDB.DB.Prepare(sqlStr1)
		if err != nil {
			fmt.Println("prepare failed, err: ", err)
			return
		}
		defer stmt1.Close()
		_, err = stmt1.Exec(user.MyName, user.MyPhone, user.MyEmail, user.MyPart, user.MyAbout, user.MyStatus, user.MyPhone)
		if err != nil {
			fmt.Println("update failed, err: ", err)
			return
		}
		w.Write([]byte("update"))
		fmt.Println("成功更新一条报名信息")
	}else {
		//写入数据库
		sqlStr2 := "insert into sign_up(my_name, my_phone, my_email, my_part, my_about, my_status) values (?, ?, ?, ?, ?, ?)"
		stmt2, err := connectDB.DB.Prepare(sqlStr2)
		if err != nil {
			fmt.Println("prepare sqlStr failed, err: ", err)
			return
		}
		defer stmt2.Close()
		_, err = stmt2.Exec(user.MyName, user.MyPhone, user.MyEmail, user.MyPart, user.MyAbout, user.MyStatus)
		if err != nil{
			fmt.Println("insert failed, err: ", err)
			return
		}else {
			w.Write([]byte("success"))
			fmt.Println("成功写入一条数据")
		}
	}
}
