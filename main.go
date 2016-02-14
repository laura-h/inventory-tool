package main

import (
	_ "inventory-tool/routers"
	"github.com/astaxie/beego"
	"github.com/Kreditech/go-cbes"
)

func registerDatabase() {
	beego.Debug(beego.AppConfig.String("elasticsearch::url"))
	settings := new(cbes.Settings)
	settings.ElasticSearch.Urls = []string{beego.AppConfig.String("elasticsearch::url")}
	settings.ElasticSearch.Index = beego.AppConfig.String("elasticsearch::index")
	settings.ElasticSearch.NumberOfShards = 5
	settings.ElasticSearch.NumberOfReplicas = 1
	settings.ElasticSearch.RefreshInterval = "10ms"
	settings.ElasticSearch.CheckOnStartup = true

	settings.CouchBase.Host = beego.AppConfig.String("couchbase::host")
	settings.CouchBase.UserName = beego.AppConfig.String("couchbase::username")
	settings.CouchBase.Pass = beego.AppConfig.String("couchbase::password")

	bucket := new(cbes.Bucket)
	bucket.Name = beego.AppConfig.String("couchbase::bucket")
	bucket.Pass = ""
	bucket.OperationTimeout = 5 // seconds

	settings.CouchBase.Bucket = bucket

	viewsOptions := new(cbes.ViewsOptions)
	viewsOptions.UpdateInterval = 5000
	viewsOptions.UpdateMinChanges = 5
	viewsOptions.ReplicaUpdateMinChanges = 5

	settings.CouchBase.ViewsOptions = viewsOptions

	err := cbes.RegisterDataBase(settings)
	if err != nil {
		beego.Critical(err)
	}
}

func main() {
//	registerDatabase()
	beego.Run()
}

