package main

import (
	"log"

	"github.com/xiaowuzai/goexample/casbin/dao"
	"github.com/xiaowuzai/goexample/casbin/server"
	"github.com/xiaowuzai/goexample/casbin/service"
	"github.com/xiaowuzai/goexample/db"
)

func main() {

	configPath := "./configs/model.conf"
	dsn := "root:123456@(localhost:3308)/goexample?charset=utf8&parseTime=true&loc=Local"

	l := db.NewDBLogger("offline")
	database, err := db.NewDB(dsn, l)
	if err != nil {
		panic(err)
	}

	repo, err := dao.NewCasbinRepoByDB(database, configPath, "", "")
	if err != nil {
		log.Fatal(err)
	}
	casbinService := service.NewCasbinService(repo)
	server := server.NewServer(casbinService)

	if err := server.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
