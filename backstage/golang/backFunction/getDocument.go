package backFunction

import (
	"connectDB"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"net/http"
	"strconv"
)

type information struct {
	id 		string
	name 	string
	phone 	string
	email 	string
	part 	string
	about 	string
	status	string
}

func GetDocument(w http.ResponseWriter, r *http.Request)  {


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



	var documentData = make(map[string]string, 32)
	xlsx := excelize.NewFile()
	index := xlsx.NewSheet("报名表")
	documentData["A1"] = "id"
	documentData["B1"] = "姓名"
	documentData["C1"] = "电话"
	documentData["D1"] = "邮箱"
	documentData["E1"] = "部门"
	documentData["F1"] = "自我介绍"
	documentData["G1"] = "状态"


	sqlStr := "select id, my_name, my_phone, my_email, my_part, my_about, my_status from sign_up where id > ?"
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
	var number = 1
	for rows.Next(){
		var data information
		err := rows.Scan(&data.id, &data.name, &data.phone, &data.email, &data.part, &data.about, &data.status)
		if err != nil {
			fmt.Println("scan failed, err: ", err)
			return
		}
		number += 1
		AStr := "A" + strconv.Itoa(number)
		BStr := "B" + strconv.Itoa(number)
		CStr := "C" + strconv.Itoa(number)
		DStr := "D" + strconv.Itoa(number)
		EStr := "E" + strconv.Itoa(number)
		FStr := "F" + strconv.Itoa(number)
		GStr := "G" + strconv.Itoa(number)
		documentData[AStr] = data.id
		documentData[BStr] = data.name
		documentData[CStr] = data.phone
		documentData[DStr] = data.email
		documentData[EStr] = data.part
		documentData[FStr] = data.about
		documentData[GStr] = data.status
	}

	for k, v := range documentData{
		xlsx.SetCellValue("报名表", k, v)
	}
	xlsx.SetActiveSheet(index)
	err = xlsx.SaveAs("../../document/报名表.xlsx")
	if err != nil {
		fmt.Println("write failed, err: ", err)
		return
	}
	fmt.Println("成功导出报名数据为excel")
	w.Write([]byte("/get/报名表.xlsx"))

}
