package main

import (
	"back/cboy"
	"fmt"
	"log"
	"net/http"
)
func solve(w http.ResponseWriter, r *http.Request) {
	log.Println("URL=", r.URL.Path)
	fmt.Fprintf(w, cboy.GetHtml(r.URL.Path)) // 这个写入到 w 的是输出到客户端的
}
func  main()  {
	log.Println("server start----")
	http.HandleFunc("/", solve) // 设置访问的路由
	err := http.ListenAndServe(":80", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}