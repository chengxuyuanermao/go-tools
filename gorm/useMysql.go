package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BlogArticle struct {
	gorm.Model
	ID      uint8
	Title   string
	Content string
}

func (ba BlogArticle) TableName() string { // gorm会自动把table带上s，此处重新定义表名
	return "blog_article"
}

func UseMysql() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	ba := &BlogArticle{Title: "testGorm", Content: "testing"}
	result := db.Select("Title", "Content").Create(&ba)
	fmt.Println(result)
}
