package controllers
import "github.com/astaxie/beego"

type UsersController struct {
	beego.Controller
}

func (this *UsersController) Get() {
	this.Data["Title"] = "Users Management"
	this.TplNames = "pages/users/index.tpl"

	this.Layout = "layouts/index.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "layouts/head.tpl"
	this.LayoutSections["Nav"] = "layouts/nav.tpl"
	this.LayoutSections["Sidebar"] = "layouts/sidebar.tpl"
	this.LayoutSections["Header"] = "layouts/header.tpl"
	this.LayoutSections["Footer"] = "layouts/footer.tpl"
}