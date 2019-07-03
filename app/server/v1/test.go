package v1

import (
	"fmt"
	"github.com/qaqzzl/happy"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("进入ControllerTest")
	//return
	//fmt.Printf("进入Controller")
	//fmt.Println(r.URL.Scheme,"Scheme")
	//fmt.Println(r.URL.Opaque,"Opaque")
	//fmt.Println(r.URL.User,"User")
	//fmt.Println(r.URL.Host,"Host")
	//fmt.Println(r.URL.Path,"Path")
	//fmt.Println(r.URL.RawPath,"RawPath")
	//fmt.Println(r.URL.ForceQuery,"ForceQuery")
	//fmt.Println(r.URL.RawQuery,"RawQuery")
	//fmt.Println(r.URL.Fragment,"Fragment")
}

func ControllerTest(this *happy.Context)  {
	fmt.Println("进入ControllerTest")
	var array map[string]interface{}
	array = make(map[string]interface{})
	array["status"] = "ok"
	array["msg"] = "ok"
	array["data"] = "进入ControllerTest"
	array["code"] = 200
	this.WJson(array)
}