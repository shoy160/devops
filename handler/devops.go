package handler

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

func GroupstHandler(c *gin.Context) {
	r := ResultNew()
	config := AppConfig{}
	data, err := ioutil.ReadFile("./config.yml")
	if err == nil {
		err = yaml.Unmarshal(data, &config)
		if err == nil {
			r = ResultSucc(config)
		} else {
			r.Code = 500
			r.Message = "配置文件格式异常"
		}
	}
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
