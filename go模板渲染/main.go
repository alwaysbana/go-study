package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./go模板渲染/hello.tmpl")
	if err != nil {
		fmt.Println("Parse has an err:%v", err)
		return
	}
	err = t.Execute(w, "fyf")
	if err != nil {
		fmt.Println("execute has an err:%v", err)
		return
	}
}
func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("the is an err:%v", err)
		return

	}
}
