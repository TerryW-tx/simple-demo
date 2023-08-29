package main

import (
	"os"
	"io"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
	"github.com/RaymondCode/simple-demo/dal"
	"gorm.io/gorm"
        "gorm.io/driver/mysql"
)

func main() {
	// model.InitDatabase()
	gormdb, _ := gorm.Open(mysql.Open("test:123456@(120.55.103.230:3306)/douyin"))
        // g.UseDB(gormdb) // reuse your gorm db
	dal.SetDefault(gormdb)

	go service.RunMessageServer()

	gin.DisableConsoleColor()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
