package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	// 获取 cycles 请求参数
	// 获取方式 1：不存在会报错
	// value := r.Form["cycles"][0]
	// 获取方式 2
	cycles := r.URL.Query().Get("cycles")
	fmt.Fprintf(w, cycles)
}
