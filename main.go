package main

import (
	"github.com/Avash027/ratings/controllers/test"
	"github.com/beego/beego/v2/server/web"
)


func main() {

	web.Router("/", &test.TestController{})
	web.Run()
}