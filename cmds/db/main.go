package main

import (
	"github.com/fghwett/typecho-to-halo/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	conf, err := config.New()
	if err != nil {
		return err
	}

	var dialector gorm.Dialector

	switch conf.Typecho.Type {
	case "mysql":
		dialector = mysql.Open(conf.Typecho.Dsn)
	case "postgres":
		dialector = postgres.Open(conf.Typecho.Dsn)
	case "sqlite3":
		dialector = sqlite.Open(conf.Typecho.Dsn)
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}

	g := gen.NewGenerator(gen.Config{
		ModelPkgPath: "../../pkg/typecho/model",
		OutPath:      "../../pkg/typecho/query",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode

		// 如果您希望空字段生成属性为指针类型，请将FieldNullable设置为true
		FieldNullable: true,
		// 如果你想分配一个在“Create”API中有默认值的字段，将Field Cover able设置为true, reference: https://gorm.io/docs/create.html#Default-Values
		FieldCoverable: true,
		// 如果要生成无符号整数类型的字段，请将FieldSignable设置为true
		FieldSignable: true,
		// 如果你想从数据库生成索引标签，设置FieldWithIndexTag为true
		FieldWithIndexTag: true,
		// 如果要从数据库生成类型标记，请将FieldWithTypeTag设置为true
		FieldWithTypeTag: true,
		// 如果需要对查询代码进行单元测试，请将WithUnitTest设置为true
		WithUnitTest: false,
	})
	g.UseDB(db)

	g.WithOpts(
		gen.FieldRename("table", "table_name"),
	)

	g.ApplyBasic(
		g.GenerateAllTable()...,
	)

	g.Execute()

	return nil
}
