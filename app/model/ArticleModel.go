package model

import (
	"fmt"
	"github.com/qaqzzl/happy/db"
	"strconv"
)

type Article struct {
	ArticleId		int		`article_id`
	class_name		string
	headline		string
	summary			string
	content			string
	created_at		int
	updated_at		int
	comm			int
	uv				int
	pv				int
	status			int
}

func GetArticleList(page int,pagecount int, where string, args ...interface{}) (article interface{},err error) {

}

