// Code generated by hertz generator.

package main

import (
	"simple_douyin/database"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	database.Init()
	h := server.Default()

	register(h)
	h.Spin()
}
