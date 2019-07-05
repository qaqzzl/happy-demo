package apiv1

import (
	"github.com/qaqzzl/happy"
	"happy-demo/app"
	"happy-demo/app/model"
)

func ArticleList(c *happy.Context) {
	where := ""
	args := []string{"golang"}
	//
	//var where [2]interface{}
	//var args [...]interface{}
	//where[0] = "class_name = ?"
	//args[0] = "golang"
	//where[1] = args
	if list, err := model.GetArticleList(1,3,where,args...); err == nil {
		c.WJson(app.Success(2000,list))
	}

}
