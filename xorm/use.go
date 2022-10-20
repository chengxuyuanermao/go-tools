package xorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

import (
	"log"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Level   int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

var engine *xorm.Engine

// https://darjun.github.io/2020/05/07/godailylib/xorm/
func Main() {
	conn()
	//createTable()
	//insert()
	//update()
	//delete()
	//doRawSql()
	selectData()
}

// 查询&统计
func selectData() {
	//selectGet()
	//selectFind()
	//selectIterate()
	selectRows()
}

func selectGet() {
	user1 := &User{}
	has, _ := engine.ID(1).Get(user1)
	if has {
		fmt.Printf("user1:%v\n", user1)
	}

	user2 := &User{}
	has, _ = engine.Where("name=?", "dj").Get(user2)
	if has {
		fmt.Printf("user2:%v\n", user2)
	}

	user3 := &User{Id: 5}
	has, _ = engine.Get(user3)
	if has {
		fmt.Printf("user3:%v\n", user3)
	}

	user4 := &User{Name: "pipi"}
	has, _ = engine.Get(user4)
	if has {
		fmt.Printf("user4:%v\n", user4)
	}
}

// Get()方法只能返回单条记录，其生成的 SQL 语句总是有LIMIT 1。Find()方法返回所有符合条件的记录。Find()需要传入对象切片的指针或 map 的指针：
func selectFind() {
	slcUsers := make([]User, 0)
	engine.Where("age > ? and age < ?", 12, 30).Find(&slcUsers)
	fmt.Println("users whose age between [12,30]:", slcUsers)

	// map的键为主键，所以如果表为复合主键就不能使用这种方式了。
	mapUsers := make(map[int64]User)
	engine.Where("length(name) = ?", 3).Find(&mapUsers)
	fmt.Println("users whose has name of length 3:", mapUsers)
}

func selectIterate() {
	engine.Where("age > ? and age < ?", 12, 30).Iterate(&User{}, func(i int, bean interface{}) error {
		fmt.Printf("user%d:%v\n", i, bean.(*User))
		return nil
	})
}

func selectRows() {
	rows, _ := engine.Where("age > ? and age < ?", 12, 30).Rows(&User{})
	defer rows.Close()

	u := &User{}
	for rows.Next() {
		rows.Scan(u)

		fmt.Println(u)
	}
}

// ----------------

func conn() {
	engineRes, err := xorm.NewEngine("mysql", "root:@tcp(127.0.0.1:3306)/blog?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	engine = engineRes
}

func createTable() {
	err := engine.Sync2(new(User))
	if err != nil {
		log.Fatal(err)
	}
}

func insert() {
	user := &User{Name: "lzy", Age: 50}
	affected, _ := engine.Insert(user)
	fmt.Printf("%d records inserted, user.id:%d\n", affected, user.Id)

	users := make([]*User, 2)
	users[0] = &User{Name: "xhq", Age: 41}
	users[1] = &User{Name: "lhy", Age: 12}
	affected, _ = engine.Insert(&users)
	// 需要注意的是，批量插入时，每个对象的Id字段不会被自动赋值，所以上面最后一行输出id1和id2均为 0
	fmt.Printf("%d records inserted, id1:%d, id2:%d", affected, users[0].Id, users[1].Id)
}

func update() {
	/**
	由于使用map[string]interface{}类型的参数，xorm无法推断表名，必须使用Table()方法指定。第一个Update()方法只会更新name字段，其他空字段不更新。第二个Update()方法会更新name和age两个字段，age被更新为 0。
	*/
	engine.ID(1).Update(&User{Name: "ldj"})
	engine.ID(1).Cols("name", "age").Update(&User{Name: "dj"})
	engine.Table(&User{}).ID(1).Update(map[string]interface{}{"age": 18})
}

func delete() {
	affected, _ := engine.Where("name = ?", "lhy").Delete(&User{})
	fmt.Printf("%d records deleted", affected)
}

// 执行原始的sql语句
func doRawSql() {
	querySql := "select * from user limit 1"
	results, _ := engine.Query(querySql)
	for _, record := range results {
		for key, val := range record {
			fmt.Println(key, string(val))
		}
	}

	updateSql := "update `user` set name=? where id=?"
	res, _ := engine.Exec(updateSql, "ldj", 1)
	fmt.Println(res.RowsAffected())
}
