package backFunction

import (
	"fmt"
	"net/http"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret"))

func LoginBack(w http.ResponseWriter, r *http.Request)  {
	password := r.FormValue("password")
	if password == "wonderful" {
		//登录成功存一个session
		username := "root"
		//给session取个名
		session, err := store.Get(r, "loginUser")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		//设置session键值对
		session.Values["user"] = username
		//保存session
		session.Save(r, w)

		w.Write([]byte("/backstage/"))
		fmt.Println("成功登录后台系统")
	}else {
		w.Write([]byte("fail"))
		fmt.Println("后台密码错误")
	}
}

func GetSession(w http.ResponseWriter, r *http.Request)  {
	//拿到一个session
	session, err := store.Get(r, "loginUser")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		w.Write([]byte("fail"))
		return
	}
	username := session.Values["user"]
	str := username.(string)
	w.Write([]byte(str))
}
