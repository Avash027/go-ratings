package routers

import (
	"github.com/Avash027/ratings/controllers/course"
	"github.com/Avash027/ratings/controllers/ratings"
	user "github.com/Avash027/ratings/controllers/user"

	"github.com/beego/beego/v2/server/web"
)

func init() {
	ns := web.NewNamespace("/api",
		web.NSNamespace("/user",
			web.NSInclude(
				&user.UserController{},
			),
		),
		web.NSNamespace("/course",
		web.NSInclude(
			&course.CourseController{},
		),
	),
		web.NSNamespace("/ratings",
		web.NSInclude(
			&ratings.RatingsControllers{},

		),
	),

	)
	web.AddNamespace(ns)


}