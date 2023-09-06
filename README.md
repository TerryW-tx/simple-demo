# simple-demo

## 抖音项目服务端简单示例

[接口信息链接](https://apifox.com/apidoc/shared-09d88f32-0b6c-4157-9d07-a36d32d7a75c/api-50707524)

工程无其他依赖，直接编译运行即可

```shell
go build && ./simple-demo
```

### TODO

* 增加log功能
* 反复判定err的代码过于冗余，待优化

### 功能说明

仅支持mysql作为数据库，配置参数在config/config.go中修改

* Register: 新用户注册
* Login: 用户登录，每次登录重新生成唯一token
* Publish: 发布视频，接收文件默认存在./public下，返回用户发布视频列表
* Feed: 推送视频，按提交时间返回该时间之前的视频，若无满足条件视频返回发布最早的视频
* Favorite: 点赞视频，返回用户喜欢视频列表
* Comment: 发布视频评论，返回视频评论列表，不支持对评论进行评论
* Relation: 关注用户，返回用户关注列表，粉丝列表
* Message: 发送消息，返回消息列表
 

### 测试

仅支持功能性测试，不支持并发量测试

通过以下命令行执行测试（依赖test包）

```shell
cd test
go test common.go [test_file.go]
```

test 目录下为不同场景的功能测试case，可用于验证功能实现正确性

其中 common.go 中的 _serverAddr_ 为服务部署的地址，默认为本机地址，可以根据实际情况修改

测试数据写在 demo_data.go 中，用于列表接口的 mock 测试
