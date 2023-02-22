### 环境
wsl2 Ubuntu-22.04

### 框架
* 使用Hertz+thrift
* 构筑命令
    * `hz new -module simple_douyin -idl idl/douoyin_service.thrift`
    * 也可以使用`./rebuild.sh`
* 更新命令
    * `hz update -idl idl/douoyin_service.thrift`
    * 也可以使用`./update.sh`

* 编译运行
    * go build && ./simple_douyin

### 数据库
* 使用gorm+sqlite
* 每次重启服务器会清空数据库并新建表项
* 默认为添加一个`unknown`用户以及属于该用户的两个视频

### 上传视频
* 用户上传视频后会新建一个`public/user_name`的文件夹，用于保存上传的视频

### 测试
* 目前只在本地以及远程使用Http请求测试过。
* 在`./test`目录下使用`publish`接口可以测试上传文件功能