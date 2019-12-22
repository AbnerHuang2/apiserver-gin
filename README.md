### apidemo-gin
基于gin的api构键脚手架，整个项目参考了生产项目的布局架构，也参考了golang项目标准布局，还有学习golang过程中见过的比较好的项目结构。
[golang项目标准布局](https://github.com/golang-standards/project-layout)
整个项目拉下来，基本做下简单的改造，就可以实际使用在生成中，特别对于新手学习搭建项目，熟悉项目规划，了解生产项目结构和细节设计感觉会有很大帮助。

### 下一步计划
- 加入Makefile, Dockerfile, 版本管理，安全

### 目前整合组件及实现功能
- 加入viper使用yml配置文件来配置项目信息，启动时可根据环境指定不同的配置文件
```html
eg.
go run main.go -config test.yml
```
- gorm 并自定义JsonTime 解决Json序列化和反序列话只支持UTC时间的问题  
提供了部分demo代码，可以按照样板代码在项目中直接使用。
- 整合redis，开箱即用
- 整合zap，lumerjack 完善日志输出，日志分割，部分参数可配置。
- 集成jwt，提供demo代码
- 实现部分工具类的封装，md5, bcrypt, uuid生成
- 应用统一封装响应格式，基本参照笔者大型项目经验规范
- 项目全局错误处理封装，使go在写业务代码时也能规范统一
- 全局应用常量
- 应用统一入口日志记录中间件实现，类似java中拦截器或者过滤器，可以很方面的在入口处记录访问日志。

### 文档和项目完善中...
