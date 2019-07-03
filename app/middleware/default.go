package middleware

import (
	"fmt"
	"github.com/qaqzzl/happy"
	"time"
)


func MiddlewareTest(this *happy.Context) {
	//code ...
	this.Handler(this)

	ticker := time.NewTicker(5 * time.Second)
	go func(ticker *time.Ticker) {
		//for range ticker.C {
		//	fmt.Println("Ticker2....")
		//}
		time.Sleep(5 * time.Second)
		fmt.Println("Ticker5 Stop")
	}(ticker)
}


