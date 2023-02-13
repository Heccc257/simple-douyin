package databases

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {
	fmt.Println("233333333333")
	db, err := gorm.Open(sqlite.Open("videos.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&VideoEntry{})

	// Create
	db.Create(&VideoEntry{})
	// db.Create(&UserEntry{Model: gorm.Model{ID: 1}, UID: 456, Name: "HCH2"})
	// db.Create(&UserEntry{Model: gorm.Model{ID: 2}, UID: 456, Name: "HCH3"})
	// fmt.Println("err: ", result.Error)
	// user := UserEntry{Model: gorm.Model{ID: 100}}
	// fmt.Printf("%+v\n", user)
	// fmt.Println(user.Model.ID)
	// user = UserEntry{}
	// db.First(&user, 1) // 根据整型主键查找
	// fmt.Println("0", user)
	// db.First(&user, "name = ?", "HCH") // 查找 code 字段值为 D42 的记录
	// fmt.Println("1", user)
	// db.First(&user, "name = ?", "HCH3") // 查找 code 字段值为 D42 的记录
	// fmt.Println("2", user)
}
