package models
import (
	"github.com/Kreditech/go-cbes"
	"github.com/astaxie/beego"
	"time"
)

type License struct {
	Id   int64
	Producer string `type:"string" analyzer:"keyword" index:"analyzed"`
	Name string `type:"string" analyzer:"keyword" index:"analyzed"`
	Version string `type:"string"`
	SerialKey string `type:"string"`
	PCNumber *PCNumber `type:"object"`
	User *User `type:"object"`
	ExpirationDate time.Time `type:"string"`
}

func init() {
	err := cbes.RegisterModel(new(License))

	if err != nil {
		beego.Critical(err)
	}
}