package main

import (
	"github.com/RaymondCode/simple-demo/config"
	"github.com/RaymondCode/simple-demo/dal"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io"
	"os"
)

func main() {
	gormdb, _ := gorm.Open(mysql.Open(config.MysqlDns))
	dal.SetDefault(gormdb)

	go service.RunMessageServer()

	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
