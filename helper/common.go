package helper

import (
	beego "github.com/beego/beego/v2/server/web"
)

func SendResponse(c *beego.Controller, httpCode int, message string, data interface{}) {
	c.Data["json"] = map[string]interface{}{
		STATUS:  httpCode,
		DATA:    data,
		MESSAGE: message,
	}
	c.ServeJSON()
}
