## 特点

1. 暂时支持目前市面主流服务商,逐步完善中
2. 统一的返回值格式，方便排查和记录

## 平台支持

* [腾讯云](https://cloud.tencent.com/product/sms)
* [Ucloud](https://www.ucloud.cn)
* [阿里云](https://www.aliyun.com/)
* [百度云](https://cloud.baidu.com/)
* [七牛云](https://www.qiniu.com/)
* [互亿无线](https://www.ihuyi.com/)
* [聚合数据](https://www.juhe.cn/)
* [时代互联](https://www.now.cn)

## 示例使用代码
直接使用
```go
 gateway := gateway.NewGatewayInterface("aliyun", &model.Config{
	 	AppID:     "AppID",
	 	AppSecret: "AppSecret",
})
Param := map[string]string{"0": "1234", "1": "1234", "2": "1234"}
message := model.NewMessage("您正在申请手机注册，验证码为：1234，5分钟内有效！", "", "", Param)
Message, err := api.SendMessage(model.NewPhone("150xxxxxxxx"), message)
fmt.Printf("Message: %v\n", Message)
fmt.Printf("err: %v\n", err)
```

开启Grpc微服务
```go
s := server.NewGrpcServer()
s.Start(50051)
```

## 预期功能  

接入多家短信平台后, 能够免接口加签的困扰, 请求后能统一返回的格式, 可做到直接使用或用grpc实现微服务接收

## 可使用功能  

| 方法              |   描述   |
| ----------------- | :------: |
| SendMessage       | 发送短信 |
| CreateSmsTemplate | 创建模板 |
| QuerySmsTemplate  | 查询模板 |

## 发送格式

总结各大短信平台是有两种发送方式

### 1. 直接发送短信内容, 如互亿无线, 时代互联

### 2. 根据模板和参数进行发送, 如阿里云, 腾讯云, 此方式也是最多平台广泛使用的

## 各平台配置说明以及接口实现情况

| 平台                                            | 发送方式 | 其他描述 | 发送短信 | 创建模板 | 查询模板 |
| ----------------------------------------------- | -------- | -------- | -------- | -------- | :------: |
| [阿里云](https://www.aliyun.com/)               | 模板     |     个人可用     | ✅        | ✅         |   ✅       |
| [腾讯云](https://cloud.tencent.com/product/sms) | 模板     |    个人可用      | ✅        |          |          |
| [Ucloud](https://www.ucloud.cn)                 | 模板     |    个人可用      | ✅        |     ✅     |    ✅      |
| [百度云](https://cloud.baidu.com/)              | 模板     |    企业      | ✅        |          |          |
| [七牛云](https://www.qiniu.com/)                | 模板     |    企业      | ✅        |          |          |
| [互亿无线](https://www.ihuyi.com/)              | 内容     |     个人可用       | ✅        |          |          |
| [聚合数据](https://www.juhe.cn/)                | 模板     |     个人可用       |       |          |          |
| [时代互联](https://www.now.cn)                  | 内容     |     个人可用       | ✅        |          |          |

### 功能逐步完善中,如果有需要接入的其他短信平台可留言备注