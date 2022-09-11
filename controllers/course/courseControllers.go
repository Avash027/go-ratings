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

// @Title Get All Courses
// @Description Get All Courses
// @Success 200 {object} Response{success:true, message:"All courses", courses:[]}
// @Failure 400 {object} Response{success:false, message:"Unable to process json"}
// @Failure 500 {object} Response{success:false, message:"Internal Server Errror"}
// @router / [get]
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

// @Title Add Course
// @Description Add Course
// @Success 200 {object} Response{success:true, message:"Course created Successfully"}
// @Failure 400 {object} Response{success:false, message:"Unable to process json"}
// @Failure 400 {object} Response{success:false, message:"Invalid course name"}
// @Failure 500 {object} Response{success:false, message:"Internal Server Errror"}
// @router / [post]
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



