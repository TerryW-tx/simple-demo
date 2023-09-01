// +build ignore

package main

import (
	"github.com/RaymondCode/simple-demo/model/dto"
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
	
	// create tables
	gormdb.AutoMigrate(
		dto.User{}, 
		dto.Video{}, 
		dto.Comment{}, 
		dto.Message{}, 
		dto.Follow{},
		dto.Favorite{},
	)

  	// Generate basic type-safe DAO API for struct `model.User` following conventions
  	// g.ApplyBasic(
	// 	dto.User{}, 
	// 	dto.Video{}, 
	// 	dto.Comment{}, 
	// 	dto.Message{}, 
	// 	dto.Follow{},
	// 	dto.Favorite{},
	// )

  	g.ApplyBasic(
	 	g.GenerateModel("users"),
	 	g.GenerateModel("videos"),
	 	g.GenerateModel("comments"),
	  	g.GenerateModel("messages"),
	  	g.GenerateModel("follows"),	
	  	g.GenerateModel("favorites"),	
	)

  	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
  	// g.ApplyInterface(func(Querier){}, model.User{}, model.Passport{})

  	// Generate the code
  	g.Execute()

	// create tables
	// gormdb.AutoMigrate(dto.User{}, dto.Video{}, dto.Comment{}, dto.Message{}, dto.Follow{})
}
