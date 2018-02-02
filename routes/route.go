package routes

import (
	"browser-manage/app/controllers"
	"browser-manage/app/middleware"
	"github.com/kataras/iris"
)

type WebRoutes struct {
}

func (WebRoutes) StartRoute(app *iris.Application) {
	app.Get("/", controllers.Test{}.Hello)

	//测试mq
	app.Get("/getmq",controllers.White{}.TestMq)

	//测试获取数据
	app.Get("/user", controllers.Test{}.GetUser)

	//测试redis
	app.Get("/redis", controllers.Test{}.GetRedis)

	//测试链接mongodb
	app.Get("/mongodb", controllers.Test{}.GetMongodb)

	//新token
	app.Get("/token", controllers.Test{}.MyToken)

	//测试jwt
	app.Get("/ping", middleware.CheckJwt, controllers.Test{}.MyToken)

	/*
		项目正式路由
	*/

	//获取验证码操作
	app.Get("/getCaptcha", controllers.Captcaha{}.GetCaptcha)

	//显示验证码
	app.Get("/captcha/{code}", controllers.Captcaha{}.Get)

	//用户登录操作
	app.Post("/login", controllers.UserLogin{}.Login)

	//用户登出操作
	app.Post("/loginOut", middleware.CheckJwt, controllers.UserLogin{}.LoginOut)

	/*
		系统账户模块路由
	*/

	//系统账号列表
	app.Get("/accountList", middleware.CheckJwt, controllers.Account{}.List)

	//添加系统账号操作
	app.Post("/addAccount", middleware.CheckJwt, controllers.Account{}.AddAccount)

	//修改时获取账号信息操作
	app.Get("/showAccount/{id}", middleware.CheckJwt, controllers.Account{}.GetAccount)

	//编辑保存账号信息操作
	app.Put("/saveAccount/{id}", middleware.CheckJwt, controllers.Account{}.SaveAccount)

	//修改密码操作
	app.Put("/savePwd/{id}", middleware.CheckJwt, controllers.Account{}.SavePwd)

	//删除账号操作
	app.Delete("/deleteAccount/{id}", middleware.CheckJwt, controllers.Account{}.DeleteAccount)

	/*
		白名单管理模块
	*/

	//获取白名单列表
	app.Get("/whiteList", middleware.CheckJwt, controllers.White{}.List)

	//添加白名单
	app.Post("/addWhite", middleware.CheckJwt, controllers.White{}.AddWhite)

	//修改白名单时获取数据
	app.Get("/showWhite/{id}", middleware.CheckJwt, controllers.White{}.ShowWhite)

	//编辑保存白名单
	app.Put("/saveWhite/{id}", middleware.CheckJwt, controllers.White{}.SaveWhite)

	//修改白名单状态
	app.Put("/saveWhiteStatus/{id}", middleware.CheckJwt, controllers.White{}.SaveWhiteStatus)

	//删除白名单
	app.Delete("/deleteWhite/{id}", middleware.CheckJwt, controllers.White{}.DeleteWhite)

	//导入白名单操作
	app.Post("/importWhite", middleware.CheckJwt, iris.LimitRequestBodySize(10<<20), controllers.White{}.Import)

	//获取导入记录
	app.Get("/getImportLog", middleware.CheckJwt, controllers.White{}.GetUploadLog)

	//下载上传记录文件
	app.Get("/down", middleware.CheckJwt, controllers.White{}.Down)

	/*
		代理服务器管理
	*/
	//代理服务器列表
	app.Get("/proxyList", middleware.CheckJwt, controllers.Proxy{}.List)

	//添加代理服务器
	app.Post("/addProxy", middleware.CheckJwt, controllers.Proxy{}.AddProxy)

	//修改时获取数据
	app.Get("/proxy/{id}", middleware.CheckJwt, controllers.Proxy{}.GetProxy)

	//修改保存数据
	app.Put("/saveProxy/{id}", middleware.CheckJwt, controllers.Proxy{}.SaveProxy)

	//删除数据
	app.Delete("/deleteProxy/{id}", middleware.CheckJwt, controllers.Proxy{}.DeleteProxy)

	/*
		后台菜单管理模块
	*/
	//菜单列表
	app.Get("/menusList", middleware.CheckJwt, controllers.Menus{}.List)

	//添加菜单
	app.Post("/addMenus", middleware.CheckJwt, controllers.Menus{}.Add)

	//修改菜单时获取数据
	app.Get("/getMenus/{id}", middleware.CheckJwt, controllers.Menus{}.Show)

	//编辑保存菜单
	app.Put("/saveMenus/{id}", middleware.CheckJwt, controllers.Menus{}.Save)

	//删除菜单
	app.Delete("/deleteMenus/{id}", middleware.CheckJwt, controllers.Menus{}.Delete)
}
