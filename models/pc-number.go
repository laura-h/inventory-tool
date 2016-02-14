package models
import (
	"github.com/Kreditech/go-cbes"
	"github.com/astaxie/beego"
)

type PCNumber struct {
	Id   int64
	PCNumber string `default:"TestName" type:"string" analyzer:"keyword" index:"analyzed"`
	User *User `type:"object"`
}

func init() {
	err := cbes.RegisterModel(new(PCNumber))

	if err != nil {
		beego.Critical(err)
	}
}