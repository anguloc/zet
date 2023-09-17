package main

import (
	"flag"
	"fmt"

	"github.com/anguloc/zet/pkg/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var config string

func main() {
	flag.StringVar(&config, "config", "conf/conf.yml", "config path")
	flag.Parse()

	if err := conf.Init(config); err != nil {
		fmt.Println(err)
		return
	}
	c := gen.Config{
		OutPath:           "./internal/app/repo/zetdao/query",
		ModelPkgPath:      "./internal/app/repo/zetdao/model",
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
	}

	g := gen.NewGenerator(c)

	db, err := gorm.Open(mysql.Open(conf.Conf().DBZet.Dsn()))
	if err != nil {
		fmt.Println(err)
		return
	}

	g.UseDB(db)

	tables, err := db.Migrator().GetTables()
	if err != nil {
		fmt.Println(err)
		return
	}
	models := make([]interface{}, len(tables))
	for i, v := range tables {
		models[i] = g.GenerateModel(v)
	}

	g.ApplyBasic(models...)

	g.Execute()
}
