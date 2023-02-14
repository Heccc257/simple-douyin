package database

import (
	"fmt"
	"log"
	"simple_douyin/biz/model/common"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open(sqlite.Open("database/douyin.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// db.Migrator().DropTable("user_entries", "video_entries")

	// 创建表之后会默认添加一个unknown用户
	// 以及两个作者为unkn的视频
	if exist := db.Migrator().HasTable("user_entries"); !exist {
		if err := db.AutoMigrate(&UserEntry{}); err != nil {
			log.Fatal(err)
		}
		// 创建一个默认用户
		db.Create(&UserEntry{
			UID:  0,
			Name: "unknown",
		})
	}
	if exist := db.Migrator().HasTable("video_entries"); !exist {
		if err := db.AutoMigrate(&VideoEntry{}); err != nil {
			log.Fatal(err)
		}
		videos := [2]VideoEntry{}
		db.Create(&videos)
	}

	log.Println("database initialized")
	tables, _ := db.Migrator().GetTables()
	log.Println("tables: ", tables)

	// var user UserEntry
	// result := db.First(&user, 1)
	// fmt.Println("````````````````````````")
	// fmt.Println("res: ", result.Error)
	// fmt.Println("user: ", user)
	// fmt.Println("````````````````````````")

	// var vs []VideoEntry
	// db.Find(&vs)
	// fmt.Println("````````````````````````")
	// fmt.Println("videos: ", vs)
	// fmt.Println("````````````````````````")
}

func FindUserEntry(u *common.User) *UserEntry {
	var ue UserEntry
	// 保证名字唯一，所以可以按名字来查找
	if result := db.Where("name = ?", u.Name).First(&ue); result.Error != nil {
		// 未找到时默认返回unknown用户
		db.First(&ue, 1)
	}

	return &ue
}

func FindVideosBefore(latest_time int64) ([]*common.Video, int64, error) {
	var ves []VideoEntry
	if result := db.Where("created_at < ?", time.Unix(latest_time, 0)).Order("created_at desc").Limit(30).Find(&ves); result.Error != nil {
		return nil, 0, result.Error
	}
	var videos []*common.Video
	for _, ve := range ves {
		fmt.Printf("%+v\n", ve)
		videos = append(videos, VideoEntry2Video(&ve))
	}
	if len(ves) == 0 { // 未找到任何视频
		return nil, 0, fmt.Errorf("%s", "No videos")
	}
	return videos, ves[0].Model.CreatedAt.Unix(), nil
}
