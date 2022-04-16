// @Title : mysqlInit
// @Description :数据库初始化
// @Author : MX
// @Update : 2022/4/14 21:03

package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func ConnDB() {
	dbUrl := "root:123456@/class?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)      // 表名为单数
	db.DB().SetMaxIdleConns(10) // 最大活跃连接数
	db.DB().SetMaxOpenConns(20) // 最大空闲连接数
	db.AutoMigrate()
	DB = db
}
