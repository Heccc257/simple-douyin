package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

var unknownUser *UserEntry = &UserEntry{
	UID:  0,
	Name: "unknown",
}

func Init() {
	DB, err = gorm.Open(sqlite.Open("database/douyin.db"), &gorm.Config{})

	// 初始化数据库
	DB.Migrator().DropTable("user_entries", "video_entries", "user_thumb", "comment_entry")

	if err != nil {
		panic("failed to connect database")
	}

	// db.Migrator().DropTable("user_entries", "video_entries")

	// 创建表之后会默认添加一个unknown用户
	// 以及两个作者为unkn的视频
	if exist := DB.Migrator().HasTable("user_entries"); !exist {
		if err := DB.AutoMigrate(&UserEntry{}); err != nil {
			log.Fatal(err)
		}
		// 创建一个默认用户
		DB.Create(unknownUser)
	}
	if exist := DB.Migrator().HasTable("video_entries"); !exist {
		if err := DB.AutoMigrate(&VideoEntry{}); err != nil {
			log.Fatal(err)
		}
		videos := [2]VideoEntry{}
		DB.Create(&videos)
	}
	if exist := DB.Migrator().HasTable("user_thumb"); !exist {
		if err := DB.AutoMigrate(&UserThumb{}); err != nil {
			log.Fatal(err)
		}
	}
	if exist := DB.Migrator().HasTable("comment_entry"); !exist {
		if err := DB.AutoMigrate(&CommentEntry{}); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("database initialized")
	tables, _ := DB.Migrator().GetTables()
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
