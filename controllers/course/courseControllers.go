package course

import (
	"encoding/json"

	"github.com/Avash027/ratings/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)


type CourseController struct {
	web.Controller
}


func (c *CourseController) GetCourses(){


	courses, err := models.GetCourses()

	if err != nil {
		logs.Critical(err.Error())
		c.Ctx.ResponseWriter.Status = 500
		c.Ctx.ResponseWriter.Write([]byte("Internal Server Error"))
		return
	}

	c.Data["json"] = courses
	c.ServeJSON()

}

func (c *CourseController) AddCourse() {
	var courseModel models.CourseData
	
	err:= json.Unmarshal(c.Ctx.Input.RequestBody , &courseModel)



	if err != nil {
		logs.Critical(err.Error())
		c.Ctx.ResponseWriter.Status = 400
		c.Ctx.ResponseWriter.Write([]byte("Bad Request"))
		return
	}

	if courseModel.CourseName == "" || courseModel.CourseDesc == "" {
		c.Ctx.ResponseWriter.Status = 400
		c.Ctx.ResponseWriter.Write([]byte("Bad Request"))
		return
	}

	courseModel.CourseRating = 0

	courseModel.TotalReviews = 0

	_, err = models.AddCourse(&courseModel)

	if err != nil {
		c.Ctx.ResponseWriter.Status = 500
		c.Ctx.ResponseWriter.Write([]byte("Internal Server Error"))
		return
	}

	c.Ctx.ResponseWriter.Status = 201
	c.Ctx.ResponseWriter.Write([]byte("Created"))
	return
}



