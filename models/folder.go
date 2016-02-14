package models
import (
	"github.com/Kreditech/go-cbes"
	"github.com/astaxie/beego"
)

type Folder struct {
	Id   int64
	Name string `type:"string" analyzer:"keyword" index:"analyzed"`
	Path string `type:"string" `
}

func init() {
	err := cbes.RegisterModel(new(Folder))

	if err != nil {
		beego.Critical(err)
	}
}