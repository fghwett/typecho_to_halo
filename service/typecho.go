package service

import (
	"fmt"
	"github.com/fghwett/typecho-to-halo/config"
	"github.com/fghwett/typecho-to-halo/pkg/typecho/model"
	"github.com/fghwett/typecho-to-halo/pkg/typecho/query"
	"github.com/spf13/cast"
	"github.com/tidwall/gjson"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Typecho struct {
	conf  *config.Typecho
	db    *gorm.DB
	query *query.Query
}

func NewTypecho(conf *config.Typecho) (*Typecho, error) {

	var open func(string) gorm.Dialector
	switch conf.Type {
	case "mysql":
		open = mysql.Open
	case "postgres":
		open = postgres.Open
	case "sqlite3":
		open = sqlite.Open
	}
	dialector := open(conf.Dsn)

	db, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	q := query.Use(db)

	t := &Typecho{
		conf:  conf,
		db:    db,
		query: q,
	}

	return t, nil
}

func (typecho *Typecho) GetTags() ([]*model.TypechoMetas, error) {
	return typecho.getMetas("tag")
}

func (typecho *Typecho) GetCategories() ([]*model.TypechoMetas, error) {
	return typecho.getMetas("category")
}

func (typecho *Typecho) getMetas(t string) ([]*model.TypechoMetas, error) {
	m := typecho.query.TypechoMetas
	return m.Where(m.Type.Eq(t)).Find()
}

func (typecho *Typecho) GetRelationShips() ([]*model.TypechoRelationships, error) {
	m := typecho.query.TypechoRelationships
	return m.Find()
}

func (typecho *Typecho) GetAttachments() ([]*model.TypechoContents, error) {
	m := typecho.query.TypechoContents
	return m.Where(m.Type.Eq("attachment")).Find()
}

func (typecho *Typecho) GetAttachmentUrl(file *model.TypechoContents) string {
	return fmt.Sprintf("%s%s", typecho.conf.BaseUrl, typecho.GetFilePath(file))
}

func (typecho *Typecho) GetFilePath(file *model.TypechoContents) string {
	return gjson.Get(cast.ToString(file.Text), "path").String()
}

func (typecho *Typecho) GetPosts() ([]*model.TypechoContents, error) {
	m := typecho.query.TypechoContents
	return m.Where(m.Type.In("post", "post_draft")).Find()
}

func (typecho *Typecho) GetPage() ([]*model.TypechoContents, error) {
	m := typecho.query.TypechoContents
	return m.Where(m.Type.In("page")).Find()
}

func (typecho *Typecho) GetComments() ([]*model.TypechoComments, error) {
	m := typecho.query.TypechoComments
	return m.Where(m.Type.Eq("comment")).Find()
}
