package main

import (
	"github.com/RaymondCode/simple-demo/model"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

// Dynamic SQL
// type Querier interface {
  // SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
  // FilterWithNameAndRole(name, role string) ([]gen.T, error)
// }

func main() {
  	g := gen.NewGenerator(gen.Config{
    		OutPath: "./dal",
    		Mode: gen.WithDefaultQuery, // generate mode
		ModelPkgPath: "./model/entity",
  	})

  	gormdb, _ := gorm.Open(mysql.Open("test:123456@(120.55.103.230:3306)/douyin"))
  	g.UseDB(gormdb) // reuse your gorm db

  	g.ApplyBasic(
		g.GenerateModel("users"),
		g.GenerateModel("videos"),
	 	g.GenerateModel("comments"),
	 	g.GenerateModel("messages"),
	 	g.GenerateModel("follows"),	
	)
  	// Generate basic type-safe DAO API for struct `model.User` following conventions
  	// g.ApplyBasic(model.User{}, model.Video{}, model.Comment{}, model.Message{}, model.Follow{})

  	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
  	// g.ApplyInterface(func(Querier){}, model.User{}, model.Passport{})

  	// Generate the code
  	g.Execute()

	// create tables
	gormdb.AutoMigrate(model.User{}, model.Video{}, model.Comment{}, model.Message{}, model.Follow{})
}
