package main

import (
	"flag"
	"fmt"

	"github.com/anguloc/zet/internal/pkg/conf"
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
		OutPath:           "./internal/dao/zet_query",
		ModelPkgPath:      "./internal/dao/zet_model",
		FieldNullable:     true,
		FieldCoverable:    true,
		FieldSignable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		Mode:              gen.WithQueryInterface,
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
