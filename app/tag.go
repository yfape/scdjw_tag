package app

import (
	"fmt"
	"net/http"
	"scdj/pool"
	"strings"

	"github.com/gin-gonic/gin"
)

type Tag struct{
	Tid int `db:"tid"`
	Name string `db:"name"`
	Entries string `db:"entries"`
}

var Tags []Tag

func Load(){
	err := pool.Mysqlpool.Select(&Tags, `SELECT * FROM tag`)
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("比对表加载完毕，共%d个比对项", len(Tags)))
}

type RequestContent struct{
	Content string `json:"content"`
}

/**
 * @description: 
 * @param {*gin.Context} c
 * @return {*}
 */
func CompareTag(c *gin.Context){
	var req RequestContent
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errorMsg": err.Error()})
		c.Abort()
		return
	}
	content := req.Content
	if content == ""{
		c.JSON(http.StatusBadRequest, gin.H{"errorMsg": "参数错误"})
		c.Abort()
		return
	}
	var restag []string
	for _, tag := range Tags{
		tagarr := strings.Split(tag.Entries, ",")
		for2: for _, item := range tagarr{
			if strings.Index(content, item) != -1{
				restag = append(restag, tag.Name)
				break for2
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"content": content, 
		"tag": restag,
	})
}

