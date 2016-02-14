package models
import (
	"github.com/Kreditech/go-cbes"
	"github.com/astaxie/beego"
)

type SecurityGroup struct {
	Id   int64
	Name string `type:"string" analyzer:"keyword" index:"analyzed"`
	Folders []*Folder `type:"object"`
}

func init() {
	err := cbes.RegisterModel(new(SecurityGroup))

	if err != nil {
		beego.Critical(err)
	}
}
