package main

import (
	"fmt"
	"net/http"
	"online_reg/controller"
)

func main() {
	// 处理静态资源路径
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages/"))))

	http.HandleFunc("/", controller.IndexHandle1)
	http.HandleFunc("/login", controller.LoginHandle)
	http.HandleFunc("/logout", controller.LogoutHandle)
	http.HandleFunc("/register", controller.RegisterHandle)
	http.HandleFunc("/checkUsername", controller.CheckUsername)

	fmt.Println("服务开启成功：地址为", "http://127.0.0.1:8082")
	http.ListenAndServe(":8082", nil)
}
