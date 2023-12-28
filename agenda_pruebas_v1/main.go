package main

import (
	"fmt"

	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/LethSphere/Ejercicio-Beego/agenda_pruebas_v1/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func main() {
	orm.RegisterDataBase("default", "postgres", "postgres://"+
		beego.AppConfig.String("PGuser")+":"+
		beego.AppConfig.String("PGpass")+"@"+
		beego.AppConfig.String("PGhost")+":"+
		beego.AppConfig.String("PGport")+"/"+
		beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+
		beego.AppConfig.String("PGschema"))
	fmt.Println("postgres://" +
		beego.AppConfig.String("PGuser") + ":" +
		beego.AppConfig.String("PGpass") + "@" +
		beego.AppConfig.String("PGhost") + ":" +
		beego.AppConfig.String("PGport") + "/" +
		beego.AppConfig.String("PGdb") + "?sslmode=disable&search_path=" +
		beego.AppConfig.String("PGschema") + "")
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"

	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	beego.Run()
}