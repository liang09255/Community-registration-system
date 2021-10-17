package backFunction

import (
	"connectDB"
	"fmt"
	"net/http"
)

func ChangeStatus(w http.ResponseWriter, r *http.Request)  {

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



	id := r.FormValue("id")
	changeInfo := r.FormValue("myStatus")
	sqlStr := "update sign_up set my_status = ? where id = ?"
	stmt, err := connectDB.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed, err: ", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(changeInfo, id)
	if err != nil {
		fmt.Println("update failed, err: ", err)
		return
	}
	w.Write([]byte("success"))
	fmt.Println("成功更新报名状态")
}