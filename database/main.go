package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {
	db, err := gorm.Open(sqlite.Open("douyin.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// db.Migrator().DropTable("user_entries", "video_entries")

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
	// db.First(&user)
	// fmt.Println("````````````````````````")
	// fmt.Println("user: ", user)
	// fmt.Println("````````````````````````")

	// var vs []VideoEntry
	// db.Find(&vs)
	// fmt.Println("````````````````````````")
	// fmt.Println("videos: ", vs)
	// fmt.Println("````````````````````````")
}
