### 环境
wsl2 Ubuntu-22.04

### 使用Hertz+thrift的方法构筑

* 构筑命令
    * `hz new -module simple_douyin -idl idl/douoyin_service.thrift`
* 更新命令
    * `hz update -idl idl/douoyin_service.thrift`

* 编译运行
    * go build && ./simple_douyin
* 测试
    * 例子
        * `curl --location --request GET "http://127.0.0.1:8888/douyin/feed/?token=100&latest_time=1000"`
        * 该命令将调用`./biz/handler/douyin_service/feed_service.go`中的`Feed`方法并且返回一个`douyin_feed_response`的json
* 具体接口定义
    * 参考`https://www.apifox.cn/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707523`和 `https://bytedance.feishu.cn/docs/doccnKrCsU5Iac6eftnFBdsXTof#bIeHvH`