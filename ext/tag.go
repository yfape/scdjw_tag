package ext

import (
	"log"
	"scdj/pool"
)

type Article struct{
	Aid int `db:"aid"`
	Catid int `db:"catid"`
	Title string `db:"title"`
	Content string `db:"content"`
	Dateline int64 `db:"dateline"`
}

type Tag struct{
	article Article
	tag string
}

func Run(){
	var articles []Article
	err := pool.Mysqlpool.Select (&articles, `SELECT a.aid, a.catid, a.title, b.content, a.dateline 
	FROM pre_portal_article_title a JOIN pre_portal_article_content b ON a.aid=b.aid where a.aid>=? and a.aid<=?`, 79821, 79830)
	if err != nil{
		log.Println(err)
	}
}