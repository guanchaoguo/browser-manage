package lang

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/kataras/iris/context"
)

type Lang struct {
}

func (Lang) GetLang(w context.Context,module string,langCode string) string  {

	locale := w.FormValue("locale")

	langPak :=  map[string]interface{} {
		//中文语言包
		"zh-cn" : map[string]interface{} {
			//公共模块语言包
			"comm"	: map[string]string {
				"Invalid" : "非法请求参数",
				"DataIsEmpty" : "数据为空",
				"Success"	: "操作成功",
				"Failure"	: "操作失败",
				"Unlink"	: "链接已断开，请重新登录",
				"UploadFileTypeError_XLSX" : "文件类型错误，只能上传 .xlsx 后缀的文件",
			},

			//登录模块语言包
			"login"	: map[string]string {
				"ReLogin" : "链接已断开，请重新登录",
				"UsernameRequired" : "用户名不能为空",
				"CaptchaRequired" : "验证码不能为空",
				"PasswordRequired" : "密码不能为空",
				"CaptchaError"  : "验证码错误",
				"UserNotExist" : "用户不存在",
				"PasswordError" : "密码错误",
				"LoginFailure" : "登录失败",
				"LoginSuccess" : "登录成功",
			},

			//系统用户管理模块语言包
			"account" : map[string]string {
				"UsernameRequired" : "用户名不能为空",
				"CaptchaRequired" : "密码不能为空",
				"AccountRequired" : "真实姓名不能为空",
				"UserNameExist" : "用户名已存在",
				"UserNameLength" : "用户名为6-50个字符长度的英文字母数字",
				"PasswordLength" : "密码长度最少为6位",
			},

			//白名单管理模块语言包
			"white" : map[string]string {
				"DomainRequired"	: "域名不能为空",
				"DomainExist"		: "域名已经存在",
				"DomainFormatError" : "域名格式错误",
				"HallNameRequired"	: "厅主不能为空",
				"IpError"			: "IP格式错误",
			},

			//代理服务器模块语言包
			"proxy" : map[string]string {
				"AddrRRequired" : "地址不能为空",
				"IpError"			: "IP格式错误",
				"PortRequired" : "端口号不能为空",
				"PortInvalid" : "无效端口号",
				"UserNameRequired" : "用户名不能为空",
				"PasswordRequired" : "密码不能为空",
			},

			//菜单语言包
			"menus" : map[string]string {
				"TitleRequired" : "菜单名称不能为空",
				"LinkRequired" : "菜单链接不能为空",
				"CodeRequired" : "菜单标识符不能为空",
				"IconRequired" : "菜单图标不能为空",
			},
		},

		//英文语言包
		"en"	: map[string]interface{} {
			//公共模块语言包
			"comm"	: map[string]string {
				"Invalid" : "Illegal request parameters",
				"DataIsEmpty" : "Data is empty",
				"Success"	: "Successful ",
				"Failure"	: "Pperation Failed",
				"Unlink"	: "The link has been disconnected. Please login again",
				"UploadFileTypeError_XLSX" : "File type error, you can only upload the.xlsx suffix file",
			},

			//登录模块语言包
			"login"	: map[string]string {
				"ReLogin" : "The link has been disconnected. Please login again",
				"UsernameRequired" : "The user name cannot be empty",
				"CaptchaRequired" : "Verify that the code cannot be empty",
				"PasswordRequired" : "Password cant be empty",
				"CaptchaError"  : "Verification code error",
				"UserNotExist" : "User does not exists",
				"PasswordError" : "Password Error",
				"LoginFailure" : "Login failure",
				"LoginSuccess" : "Login successfully",
			},

			//系统用户管理模块语言包
			"account" : map[string]string {
				"UsernameRequired" : "The user name cannot be empty",
				"CaptchaRequired" : "Password cant be empty",
				"AccountRequired" : "Real names can't be empty",
				"UserNameExist" : "User name already exists",
				"UserNameLength" : "The user is called a 6-50 character length of English letter + number combination",
				"PasswordLength" : "The password is at least 6 digits long",
			},

			//白名单管理模块语言包
			"white" : map[string]string {
				"DomainRequired"	: "Domain name cannot be empty",
				"DomainExist"		: "The domain name already exists",
				"DomainFormatError" : "Domain Name Form registering sites error",
				"HallNameRequired"	: "The main hall cannot be empty",
				"IpError"			: "IP format error",
			},

			//代理服务器模块语言包
			"proxy" : map[string]string {
				"AddrRRequired" : "The address cannot be empty",
				"IpError"			: "IP format error",
				"PortRequired" : "The port number cannot be empty",
				"UserNameRequired" : "The user name cannot be empty",
				"PasswordRequired" : "Password cant be empty",
			},
		},
	}

	//判断键值是否存在，不存在默认为中文
	if _,ok := langPak[locale]; !ok{
		locale = "zh-cn"
	}
	jsonLang,_ := json.Marshal(langPak)
	return gjson.Get(string(jsonLang),locale + "." + module + "." +langCode).String()
}
