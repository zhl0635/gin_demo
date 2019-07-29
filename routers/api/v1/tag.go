package v1

import (
	"gin_demo/models"
	"gin_demo/pkg/e"
	"gin_demo/pkg/setting"
	"gin_demo/pkg/util"
	"github.com/Unknwon/com"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTags(c *gin.Context)  {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state, _ = com.StrTo(arg).Int()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})

}

func AddTag(c *gin.Context)  {

}

func EditTag(c *gin.Context)  {
	
}

func DeleteTag(c *gin.Context)  {
	
}