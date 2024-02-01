package model

import "gorm.io/gorm"

// category指向category分类类型；cid指向分类的id；Desc为描述;content为文章内容；
// foreignkey指定用作联接表中外键的当前模型的列名
type Article struct {
	Category Category `gorm:"foreignKey:Cid"`
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}
