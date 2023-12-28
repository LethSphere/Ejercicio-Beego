package main

import (
	"context"

	_ "github.com/LethSphere/Ejercicio-Beego/agenda_pruebas_v1/routers"

	"github.com/astaxie/beego/plugins/cors"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/lib/pq"
)

func corsHandler(ctx *context.Context) {
	cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "x-requested-with", "content-type", "accept", "origin", "authorization", "x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})(ctx)
}

func main() {
	orm.RegisterDataBase("default", "postgres", "postgres://"+
		beego.AppConfig.String("PGuser")+":"+
		beego.AppConfig.String("PGpass")+":"+
		beego.AppConfig.String("PGhost")+"@"+
		beego.AppConfig.String("PGport")+"/"+
		beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+
		beego.AppConfig.String("PGschema")+"")

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  false,
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "x-requested-with", "content-type", "accept", "origin", "authorization", "x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           0,
	}))
	beego.Run()
}
