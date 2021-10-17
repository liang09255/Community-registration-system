package newFunction

import (
	"connectDB"
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct{
	Code 	int
	Msg 	string
	Result 	interface{}
}

func NewJsonResponse(code int, msg string, result interface{}) *JsonResponse {
	return &JsonResponse{
		Code: code,
		Msg: msg,
		Result: result,
	}
}

func SearchPhone(w http.ResponseWriter, r *http.Request)  {
	sMyPhone := r.FormValue("sMyPhone")

	sqlStr := "select my_name, my_phone, my_email, my_part, my_about, my_status from sign_up where id > ?"
	stmt, err := connectDB.DB.Prepare(sqlStr)
	if err != nil{
		fmt.Println("prepare failed, err: ", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil{
		fmt.Println("scan failed, err: ", err)
		return
	}
	defer rows.Close()

	judgeExist := 0
	var user Information
	for rows.Next() {
		err := rows.Scan(&user.MyName, &user.MyPhone, &user.MyEmail, &user.MyPart, &user.MyAbout, &user.MyStatus)
		if err != nil{
			fmt.Println("scan failed, err: ", err)
			return
		}
		if sMyPhone == user.MyPhone{
			judgeExist = 1
			//先转为rune再转回string确保截取中文完整
			aboutRune := []rune(user.MyAbout)
			number := len(aboutRune)
			if number >= 90 {
				user.MyAbout = string(aboutRune[0:90]) + "······"
			}
			break
		}
	}

	if judgeExist == 1 {
		res := NewJsonResponse(200,"Request success", user)
		request, _ :=json.Marshal(res)
		w.Write(request)
		fmt.Println("查询成功")
	}else {
		w.Write([]byte("查询失败"))
		fmt.Println("查询失败")
	}

}