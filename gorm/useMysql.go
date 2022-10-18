package gorm

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BlogArticle struct {
	//gorm.Model，加上时间戳等字段了
	ID      uint8
	Title   string
	Content string
}

func (ba BlogArticle) TableName() string { // gorm会自动把table带上s，此处重新定义表名
	return "blog_article"
}

func UseMysql() {
	dsn := "root:@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// 单个插入
	//ba := &BlogArticle{Title: "testGorm", Content: "testing"}
	//result := db.Select("Title", "Content").Create(&ba)
	//fmt.Println(result.Error, result.RowsAffected)

	// 批量插入
	//blogs := []BlogArticle{
	//	{Title: "test1", Content: "c1"},
	//	{Title: "test2", Content: "c2"},
	//}
	//db.Create(&blogs)
	//for _, blog := range blogs {
	//	fmt.Println(blog)
	//}

	// 批量插入 -- map类型
	//blogs2 := []map[string]interface{}{
	//	{"title": "t1", "Content": "c1"},
	//	{"title": "t2", "Content": "c2"},
	//}
	//db.Model(&BlogArticle{}).Create(blogs2)

	// 查询
	var ba = &BlogArticle{}
	res := db.First(&ba)
	if res.Error != nil {
		fmt.Println(res.Error)
	}
	fmt.Println(ba)

}
