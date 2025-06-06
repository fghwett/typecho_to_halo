// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTypechoContents = "typecho_contents"

// TypechoContents mapped from table <typecho_contents>
type TypechoContents struct {
	Cid          uint32  `gorm:"column:cid;type:int(10) unsigned;primaryKey;autoIncrement:true" json:"cid"`
	Title        *string `gorm:"column:title;type:varchar(150)" json:"title"`
	Slug         *string `gorm:"column:slug;type:varchar(150);uniqueIndex:slug,priority:1" json:"slug"`
	Created      *uint32 `gorm:"column:created;type:int(10) unsigned;index:created,priority:1" json:"created"`
	Modified     *uint32 `gorm:"column:modified;type:int(10) unsigned" json:"modified"`
	Text         *string `gorm:"column:text;type:longtext" json:"text"`
	Order_       *uint32 `gorm:"column:order;type:int(10) unsigned" json:"order"`
	AuthorID     *uint32 `gorm:"column:authorId;type:int(10) unsigned" json:"authorId"`
	Template     *string `gorm:"column:template;type:varchar(32)" json:"template"`
	Type         *string `gorm:"column:type;type:varchar(16);default:post" json:"type"`
	Status       *string `gorm:"column:status;type:varchar(16);default:publish" json:"status"`
	Password     *string `gorm:"column:password;type:varchar(32)" json:"password"`
	CommentsNum  *uint32 `gorm:"column:commentsNum;type:int(10) unsigned" json:"commentsNum"`
	AllowComment *string `gorm:"column:allowComment;type:char(1)" json:"allowComment"`
	AllowPing    *string `gorm:"column:allowPing;type:char(1)" json:"allowPing"`
	AllowFeed    *string `gorm:"column:allowFeed;type:char(1)" json:"allowFeed"`
	Parent       *uint32 `gorm:"column:parent;type:int(10) unsigned" json:"parent"`
}

// TableName TypechoContents's table name
func (*TypechoContents) TableName() string {
	return TableNameTypechoContents
}
