package main

import (
	"bookstat/controller"
	"html/template"
	"net/http"
)

func Indexhandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("views/index.html"))

	t.Execute(w, "")

}
func main() {

	http.HandleFunc("/main", Indexhandler)
	//解析文件模板
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	http.HandleFunc("login", controller.Login)
	http.HandleFunc("/regist", controller.regist)
	http.ListenAndServe(":8080", nil)
}
