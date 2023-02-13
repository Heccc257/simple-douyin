package databases

import (
	"fmt"
	"simple_douyin/biz/model/common"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	common.Video
}

func Init() {
	fmt.Println("233333333333")
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Video{})
	db.Create(&Video{ID: 123})
	db.Create(&Video{ID: 123})
	var video Video
	db.First(&video, 1) // 根据整型主键查找
	fmt.Println(video)
	db.First(&video, "id = ?", 123) // 查找 code 字段值为 D42 的记录
}
