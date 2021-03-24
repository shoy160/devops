package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

func GroupDockerHandler(c *gin.Context) {
	groups := make(map[string]string)
	groups["community"] = "智慧社区"
	groups["i-town"] = "智慧城市"
	groups["hainan"] = "海南环岛"
	groups["d-net"] = "数据通信"
	groups["idc"] = "IDC业务"
	groups["yzworld"] = "云智天下"
	groups["yzcloud"] = "云智云"
	c.JSON(http.StatusOK, ResultSucc(groups))
}

func GroupProjectHandler(c *gin.Context) {
	r := ResultNew()
	c.JSON(http.StatusOK, r)
}

func GennerateHandler(c *gin.Context) {
	var json DevOps
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusOK, ResultError(401, err.Error()))
	} else {
		dir := fmt.Sprintf("template/%s", json.Type)
		temp := fmt.Sprintf("_temp/%s_%s_%d", json.Type, json.Workload, rand.Int31())
		// copy(dir, temp, json)
		zip_path, err := zipTemplate(dir, temp, json)
		if err != nil {
			c.JSON(http.StatusOK, ResultError(500, err.Error()))
		} else {
			c.FileAttachment(zip_path, path.Base(zip_path))
		}
	}
}
