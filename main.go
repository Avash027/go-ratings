package main

import (
	"os"
	"strings"

	"github.com/Avash027/ratings/controllers/course"
	"github.com/Avash027/ratings/controllers/ratings"
	"github.com/Avash027/ratings/controllers/test"
	user "github.com/Avash027/ratings/controllers/user"
	_ "github.com/Avash027/ratings/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


func main() {
	godotenv.Load()


	logs.Async()
	logs.SetLogger(logs.AdapterConsole)

	if err := orm.RegisterDriver("postgres", orm.DRPostgres); err != nil {
		logs.Critical("register driver failed")
		panic(err)
	}

	if err := orm.RegisterDataBase("default", "postgres", strings.TrimSpace(os.Getenv("DB_URL"))); err != nil {
		logs.Critical("register database failed")
		panic(err)
	}


	if err := orm.RunSyncdb("default", false, true); err != nil {
		logs.Critical("run sync db fudged up", err.Error())
		panic(err.Error())
	}

	

	web.Router("/", &test.TestController{})
	web.Router("/api/course" ,&course.CourseController{} , "post:AddCourse" )
	web.Router("/api/course",&course.CourseController{} , "get:GetCourses")

	web.Router("/api/user",&user.UserController{}, "post:AddUser")

	web.Router("/api/rating",&ratings.RatingsControllers{} , "post:AddRating")
	web.Run(":"+os.Getenv("PORT"))
}