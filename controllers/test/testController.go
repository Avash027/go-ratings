package test

import (
	"github.com/beego/beego/v2/server/web"
)

type TestController struct {
	web.Controller
};

func (c *TestController) Get() {
	c.Ctx.WriteString("hello world")
}
