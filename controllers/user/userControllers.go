package user

import (
	"encoding/json"

	"github.com/Avash027/ratings/models"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
web.Controller
}

type Response struct {
	Success bool `json:"success"`
	Message string `json:"message"`

}

// @Title Add User
// @Description Add User
// @Success 200 {object} Response{success:true, message:"User created"}
// @Failure 400 {object} Response{success:false, message:"Unable to process json"}
// @Failure 500 {object} Response{success:false, message:"Internal Server Errror"}
// @router / [post]

func (c *UserController) AddUser(){
	var User models.User

	err := json.Unmarshal(c.Ctx.Input.RequestBody , &User)

	if err != nil {
		c.Ctx.Output.Status = 400
		c.Data["json"] = Response{false, "Unable to unmarshal json"}
		c.ServeJSON()
		return
	}

	err = models.AddUser(&User)

	if err!= nil {
		logs.Critical(err.Error())
		c.Ctx.Output.Status = 500
		c.Data["json"] = Response{false, "Internal Server Error"}
		c.ServeJSON()
		return
	}

		c.Ctx.Output.Status = 200
		c.Data["json"] = Response{false, "User created"}
		c.ServeJSON()


}