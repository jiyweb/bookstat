package controller

import (
	"bookstat/dao"
	"net/http"
	"text/template"
)

func Login(w http.ResponseWriter, r *http.Request) {

	//获取前端传过来的值
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.ID > 0 {

		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, "")
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "")
	}
}

//注册

func Regist(w http.ResponseWriter, r *http.Request) {

	//获取前端传过来的值
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {

		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已经存在")
	} else {
		t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
		t.Execute(w, "用户名密码不正确或者未注册")
	}
}

//CheckUserName
//通过发送Ajax请求验证用户名的存在
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {

		w.Write([]byte("用户名存在"))
	} else {
		w.Write([]byte("用户名不存在"))
	}
}
