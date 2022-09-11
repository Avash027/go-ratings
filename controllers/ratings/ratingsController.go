package ratings

import (
	"encoding/json"

	"github.com/Avash027/ratings/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type RatingsControllers struct {
	web.Controller
}

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`
}


// @router / [post]
func (c *RatingsControllers) AddRating() {
	var UserRating models.UserRatings

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &UserRating)

	logs.Info(err)

	if err != nil {
		c.Ctx.ResponseWriter.Status = 400
		c.Data["json"] = Response{false, "Unable to process json"}
		c.ServeJSON()
		return
	}

	if UserRating.Rating <=0 || UserRating.Rating > 5 {
		c.Ctx.Output.Status = 400
		c.Data["json"] = Response{false, "Invalid UserRatings"}
		c.ServeJSON()
		return
	}

	_,err = models.UpdateRating(&UserRating)

	if err != nil {
		c.Ctx.Output.Status = 500
		c.Data["json"] = Response{false, "Internal Server Error"}
		c.ServeJSON()
		return
	}

	
		c.Ctx.Output.Status = 200
		c.Data["json"] = Response{true, "Rating Updated Successfully"}
		c.ServeJSON()
	

}