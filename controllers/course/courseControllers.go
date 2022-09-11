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

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Courses []*models.CourseData `json:"courses"`
}


func (c *CourseController) GetCourses(){


	courses, err := models.GetCourses()


	if err != nil {
		logs.Critical(err.Error())
		c.Ctx.Output.Status = 500
		c.Data["json"] = Response{false, "Internal Server Errror" , nil}
		c.ServeJSON()
		return
	}

	c.Ctx.Output.Status = 200
	c.Data["json"] = Response{true, "All courses" , courses}
	c.ServeJSON()


}

func (c *CourseController) AddCourse() {
	var courseModel models.CourseData
	
	err:= json.Unmarshal(c.Ctx.Input.RequestBody , &courseModel)



	if err != nil {
		logs.Critical(err.Error())
		c.Ctx.Output.Status = 400
		c.Data["json"] = Response{false, "Cannot unmarshal json" , nil}
		c.ServeJSON()
		return
	}

	if courseModel.CourseName == "" {
		c.Ctx.Output.Status = 400
		c.Data["json"] = Response{false, "Invalid course name" , nil}
		c.ServeJSON()
		return
	}

	courseModel.CourseRating = 0

	courseModel.TotalRatingNum = 0

	_, err = models.AddCourse(&courseModel)

	if err != nil {
		c.Ctx.Output.Status = 500
		c.Data["json"] = Response{false, "Internal Server Errror" , nil}
		c.ServeJSON()
		return
	}

		c.Ctx.Output.Status = 200
		c.Data["json"] = Response{true, "Course created Successfully" , nil}
		c.ServeJSON()

}



