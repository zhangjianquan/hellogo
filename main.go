package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
)

type IndexData struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData IndexData
	indexData.Title = "码神之路"
	indexData.Desc = "入门教程"
	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)
}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData IndexData
	indexData.Title = "码神之路"
	indexData.Desc = "入门教程"
	t := template.New("index.html")
	//拿到当前的路径
	path, _ := os.Getwd()
	t, _ = t.ParseFiles(path + "/template/index.html")
	t.Execute(w, indexData)
}

func main() {
	//程序如楼，一个项目只能有一个入口
	//web程序，http协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)
	if error := server.ListenAndServe(); error != nil {
		log.Println(error)

	}
}
