package entity

type Tag struct {
	ID     uint   `gorm:"primaryKey;"`
	Locale string `gorm:"primaryKey;unique"`
	Value  string
	Blog   []*Blog `gorm:"many2many:blog_tags"`
}