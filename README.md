# simple_regist_login

## 功能
基于gin,gorm,go-redis实现了简单的用户注册、登录与验证。
用户注册后会跳转到登录页面，登录后将跳转到个人信息页面（未开发，只返回了用户名和ID）。
服务端将随机生成token key保存到浏览器的cookie中，然后将token key以及token保存到redis数据库中。
当浏览器使用url直接访问/user/info时，会首先进行token验证，验证不通过将不会展示用户信息。


## 路由
  * GET   /regist       用户注册页面
  * GET   /login        用户登录页面
  * POST  /user/regist  注册表单提交
  * POST  /user/login   登录表单提交
  * GET   /user/info    用户信息

## 运行方式
  1. 修改配置文件config.ini
  2. 在主目录下执行`go run main.go`
