# 电信CTWing接入

## 注册插件

- 登录到超管后台
- 进入插件管理添加新服务
  - ![插件管理](./image/image.png)
  - 服务名称：电信CTWing（可自定义）
  - 服务标识符：CTWING
  - 类别：接入服务
  - ![新增插件](./image/image-1.png)
- 点击服务配置按钮配置服务
  - HTTP服务地址：127.0.0.1:8381（提供给平台的HTTP服务，填写平台能够访问到的地址和端口）
  - 服务订阅主题前缀：service/ctwing
  - ![服务配置](./image/image-2.png)

## 对接步骤

 查看文档：[电信CTWing对接](http://thingspanel.io/zh-Hans/docs/device-connect/service_connect/ctwing)
