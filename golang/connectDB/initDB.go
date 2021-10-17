package connectDB

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

//初始化数据库
func InitDB(){
	dsn := "root:root@tcp(127.0.0.1)/sign_up?charset=utf8mb4&parseTime=True"
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil{
		fmt.Println("open failed, err: ", err)
		return
	}
	err = DB.Ping()
	if err != nil{
		fmt.Println("ping failed, err: ", err)
		return
	}
	fmt.Println("数据库初始化完成")
}


