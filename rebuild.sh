#!/bin/bash
rm -rf biz main.go router_gen.go router.go
hz new -module simple_douyin -idl idl/douyin_service.thrift
go mod tidy
go build