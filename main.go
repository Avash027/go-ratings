package main

import (
	"os"

	"github.com/Avash027/ratings/controllers/test"
	"github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load()
	web.Router("/", &test.TestController{})
	web.Run(":"+os.Getenv("PORT"))
}