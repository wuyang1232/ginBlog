package model

import (
	"fmt"
	"ginBlog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

// 配置数据库连接参数
// 引入两个全局包类参数
var db *gorm.DB
var err error

func InitDB() {
	//连接数据库"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{

		NamingStrategy: schema.NamingStrategy{
			//使用单数表名，启用此选项后，“User”的表将为“user”否则表名为复数"users"
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
	}

	//AutoMigrate 用于自动迁移您的 schema，保持您的 schema 是最新的。
	//注意： AutoMigrate 会创建表、缺失的外键、约束、列和索引。 如果大小、精度、是否为空可以更改，则 AutoMigrate 会改变列的类型。 出于保护您数据的目的，它 不会 删除未使用的列
	db.AutoMigrate(&User{}, &Category{}, &Article{})

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。需要保证不会大于框架的连接时间，既router.go中的r.Run(utils.HttpPort)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	//sqlDB.Close()
}
