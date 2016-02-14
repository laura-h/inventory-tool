package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplNames = "pages/index.tpl"

	this.Layout = "layouts/index.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "layouts/head.tpl"
	this.LayoutSections["Nav"] = "layouts/nav.tpl"
	this.LayoutSections["Sidebar"] = "layouts/sidebar.tpl"
	this.LayoutSections["Header"] = "layouts/header.tpl"
	this.LayoutSections["Footer"] = "layouts/footer.tpl"
}

func (this *MainController) Login() {
	this.TplNames = "pages/login.tpl"
}
