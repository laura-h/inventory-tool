package models
import (
	"github.com/Kreditech/go-cbes"
	"github.com/astaxie/beego"
	"time"
)

type History struct {
	Id   int64
	SecurityGroup *SecurityGroup `type:"object"`
	Date time.Time `type:"string"`
}

func init() {
	err := cbes.RegisterModel(new(History))

	if err != nil {
		beego.Critical(err)
	}
}
