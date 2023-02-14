package database

import (
	"fmt"
	"os"
	"simple_douyin/biz/model/common"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestMain(m *testing.M) {
	db, err = gorm.Open(sqlite.Open("douyin.db"), &gorm.Config{})
	code := m.Run()
	os.Exit(code)
}

func TestInit(t *testing.T) {
	Init()
}

func TestFindUserEntry(t *testing.T) {
	ue := FindUserEntry(&common.User{
		Name: "unknown",
	})
	fmt.Println(*ue)

	ue = FindUserEntry(&common.User{
		Name: "2333",
	})
	fmt.Println(*ue)
}

func TestFeed(t *testing.T) {
	videos, latest_time, err := Feed(9223372036854775807)
	fmt.Println("max time ", time.Unix(9223372036854775807, 0))
	fmt.Println("time: ", latest_time)
	fmt.Println("err: ", err)
	fmt.Println("videos: ", videos)

}
