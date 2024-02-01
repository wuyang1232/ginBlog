package model

import "gorm.io/gorm"

// 结构体首字母大写全局
type User struct {
	//gorm提供的四个参数：id,更新时间，创建时间，删除时间
	//role：角色：管理员，普通用户
	//json用于前后端交互
	gorm.Model
	Username string `gorm:"type: varchar(20); not null" json:"username"`
	Password string `gorm:"type: varchar(20); not null" json:"password"`
	Role     int    `gorm:"type: int" json:"role"`
}
