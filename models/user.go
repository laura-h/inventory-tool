package models
import (
	"github.com/Kreditech/go-cbes"
	"github.com/astaxie/beego"
)

type User struct {
	Id   int64
	Username string `type:"string" analyzer:"keyword" index:"analyzed"`
	Email string `type:"string" analyzer:"keyword" index:"analyzed"`
	Role string `type:"string"`
	PCNumber *PCNumber `type:"object"`
}

func init() {
	err := cbes.RegisterModel(new(User))

	if err != nil {
		beego.Critical(err)
	}
}