package model

import (
	"fmt"
	"github.com/qaqzzl/happy/db"
	"strconv"
)


func GetArticleList(page int,pagecount int, where string, args ...string) (article interface{},err error) {
	if article, err = db.Table("article").
		Select("article_id,class_name,summary,comm,headline,pv,updated_at").
		Limit( strconv.Itoa(((page-1)*pagecount))+","+strconv.Itoa(pagecount) ).
		Where(where, args).
		Get(); err != nil {
		fmt.Println("错误处理... ")
	}
	return article,err
}

