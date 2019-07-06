package apiv1

import (
	"github.com/qaqzzl/happy"
	"happy-demo/app"
	"happy-demo/app/model"
)

func ArticleList(c *happy.Context) {
	where := "class_name = ?"

	if list, err := model.GetArticleList(1,3,where,"golang1"); err == nil {
		c.WJson(app.Success(2000,list))
	}

}
