package entity

type Blog struct {
	ID       	int
	Author   	Author `gorm:"embedded"`
	Author2  	Author `gorm:"embedded;embeddedPrefix:vlog_"`
	Upvotes  	int32
	Tag	 	 	[]*Tag `gorm:"many2many:blog_tags;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT"`
	LocaleTags 	[]*Tag `gorm:"many2many:locale_blog_tags;ForeignKey:ID;JoinForeignKey:BlogID;References:Locale;JoinReferences:RenamedLocale"`
	/*
		ForeignKey:ID and JoinForeignKey:BlogID, it means the column's name will be blog_id and references(Blog.ID)
		References:Locale and JoinReferences:RenamedLocale, it means the column's name will be renamed_locale and references(Tag.Locale)
	*/
}
