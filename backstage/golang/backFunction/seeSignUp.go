package backFunction

import (
	"connectDB"
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct {
	Code	int
	Msg		string
	Result  interface{}
}

func NewJsonResponse(code int, msg string, result interface{}) *JsonResponse  {
	return &JsonResponse{
		Code:   code,
		Msg:    msg,
		Result: result,
	}
}

type Application struct {
	Id 		int
	Name 	string
	Phone 	string
	Email 	string
	Part 	string
	Status 	string
}

func NewApplication(id int, name string, phone string, email string, part string, status string) *Application {
	return &Application{
		Id:     id,
		Name:   name,
		Phone:  phone,
		Email:  email,
		Part:   part,
		Status: status,
	}
}

func SeeSignUp(w http.ResponseWriter, r *http.Request)  {


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


	sqlStr := "select id, my_name, my_phone, my_email, my_part, my_status from sign_up where id > ?"
	stmt, err := connectDB.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("Prepare failed, err:", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(0)
	if err != nil {
		fmt.Println("scan failed, err: ", err)
		return
	}
	defer rows.Close()

	var needData []*Application

	for rows.Next(){
		var data Application
		err := rows.Scan(&data.Id, &data.Name, &data.Phone, &data.Email, &data.Part, &data.Status)
		if err != nil {
			fmt.Println("scan failed, err: ", err)
			return
		}
		needData = append(needData, NewApplication(data.Id, data.Name, data.Phone, data.Email, data.Part, data.Status))
	}

	res := NewJsonResponse(200, "success", needData)
	request, _ := json.Marshal(res)
	w.Write(request)
}
