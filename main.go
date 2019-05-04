package main

import (
	"github.com/qaqzzl/happy"
	"happy-demo/app/server/home"
)

func main()  {
	//http.HandleFunc("/", httptts)
	//http.ListenAndServe(":8888", nil)
	h := happy.New()
	h.RouteAny("/", home.Test)
	h.Send(":8088")
}

func init() {
	//初始化路由

}