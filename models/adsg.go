package models
import (
	"github.com/Kreditech/go-cbes"
	"github.com/astaxie/beego"
)

type Adsg struct {
	Id   int64
	Type string `type:"string" analyzer:"keyword" index:"analyzed"`
	User *User `type:"object"`
	SecurityGroups []*SecurityGroup `type:"object"`
}

func init() {
	err := cbes.RegisterModel(new(Adsg))

	if err != nil {
		beego.Critical(err)
	}
}