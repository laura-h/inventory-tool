package routers

import (
	"inventory-tool/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.MainController{}, "get:Login")

	beego.Router("/access-management", &controllers.AccessManagementController{})
	beego.Router("/licenses", &controllers.LicensesController{})

	beego.Router("/users", &controllers.UsersController{})
	beego.Router("/security-groups", &controllers.SecurityGroupController{})
	beego.Router("/pc-management", &controllers.PCController{})

    beego.Router("/login", &controllers.MainController{}, "get:Login")


}
