package controllers
import "github.com/astaxie/beego"

type SecurityGroupController struct {
	beego.Controller
}

func (this *SecurityGroupController) Get() {
	this.Data["Title"] = "Security Groups Management"
	this.TplNames = "pages/security-groups/index.tpl"

	this.Layout = "layouts/index.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["HtmlHead"] = "layouts/head.tpl"
	this.LayoutSections["Nav"] = "layouts/nav.tpl"
	this.LayoutSections["Sidebar"] = "layouts/sidebar.tpl"
	this.LayoutSections["Header"] = "layouts/header.tpl"
	this.LayoutSections["Footer"] = "layouts/footer.tpl"
}