###博客后端系统

这个博客项目是用来练习的。本博客使用gin框架进行开发，简单开发了几个接口。

####目录介绍

* config:配置文件
* does:文档集合
* global:全局变量
* internal:内部模块
	* dao:数据访问层，所有与数据相关的操作都会在dao层
	* middleware:HTTP中间件
	* model:模型层
	* routers:路由相关逻辑
	* service:相关核心业务逻辑
* pkg:项目相关的模块包
* storage:项目生产的临时文件
* scripts:各类构建、安装、分析等操作的脚本
* third_party:第三方的资源工具


####中间件介绍
* 文件配置开源库viper
* 数据库相关操作开源库gorm
* 日志组件开源库lumberjack
* 接口文档资源库swagger

