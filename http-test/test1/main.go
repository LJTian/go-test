package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t, err := template.ParseFiles("./html/index.html")
	if err != nil {
		fmt.Printf("func template.ParseFiles faid, err is [%v]", err)
		return
	}
	// 渲染模板
	name := "tianLJ"
	t.Execute(w, name)
}

func sayHello(w http.ResponseWriter, r *http.Request) {

	// 自定义一个夸人的模板函数
	kua := func(arg string) (string, error) {
		return arg + "真帅", nil
	}
	// 采用链式操作在Parse之前调用Funcs添加自定义的kua函数
	tmpl, err := template.New("hello").Funcs(template.FuncMap{"kua": kua}).ParseFiles("./html/index.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}

	name := "tianLJ"
	// 使用user渲染模板，并将结果写入w
	tmpl.Execute(w, name)
}

func indexBlock(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	// 解析模板
	t1, err := template.ParseFiles("./html/base.html", "./html/index.html")
	if err != nil {
		fmt.Printf("func template.ParseFiles faid, err is [%v]\n", err)
		return
	}
	// 渲染模板
	name := "tianLJ"
	t1.ExecuteTemplate(w, "index.html", name)
}

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/indexBlock", indexBlock)
	http.ListenAndServe(":9000", nil)

}
