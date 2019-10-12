# 使用gin 框架 搭建自己的 web api

[gin 参考资料](https://github.com/gin-gonic/gin)


---------
## 简述
- 集成golog ，用于业务日志输出
- 集成gprofile,可参数化系统配置
- 日志中间件：输出请求的 uri 相关信息
- 集成swagger。自动生成在线api 文档 。访问地址：http://ip:port/[context-path]/api/index.html （eg:http://localhost:8188/dsp/swagger/index.html）
- 增加 中间件 middlewares，在Bootstrapper 中全局使用
- 修改日志输出模式：同时输出 stdout，file
---
## swagger 具体使用
[swagger](./swagger.md)

[swagger api具体使用 ](https://github.com/swaggo/swag)




