package controllers

import (
	"browser-manage/app/helper"
	"browser-manage/app/models"
	"browser-manage/lang"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"math"
	"strconv"
	"time"
	"fmt"
)

/*
	代理服务器管理
*/
type Proxy struct {
}

/**
* @api {Get} /proxyList 代理服务器列表
* @apiDescription 代理服务器列表
* @apiGroup proxy
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiParam {String} token token
* @apiParam {String} search 输入框条件
* @apiParam {Number} per_page 每页显示数据条数，默认15条
* @apiParam {Number} page 当前的所在页码，默认第1页
* @apiSuccessExample {json} Success-Response:
*	{
*	  "code": 0,
*	  "message": "操作成功",
*	  "result": {
*		"current_page": 1,
*		"data": [
*		 {
*			"id": 2,
*			"addr": "192.168.31.155",
*			"port": 8080,
*			"user_name": "sunwukong",
*			"password": "123456",
*			"last_update_date": "0000-00-00 00:00:00"
*		  }
*		],
*		"last_page": 1,
*		"per_page": 15,
*		"total": 2
*	  }
*	}
 */
func (Proxy) List(w context.Context) {
	search := w.FormValue("search")
	per_page := w.URLParamDefault("per_page", "15")
	page := w.URLParamDefault("page", "1")
	newPage, _ := strconv.Atoi(page)
	nowPage := int32(newPage)
	proxy := new(models.Proxy)

	//分页结算
	var count int64
	if search != "" {
		count, _ = models.Proxy{}.GetObj().Where("user_name = ? or addr = ? ", search).Count(proxy) //总的数据量
	} else {
		count, _ = models.Proxy{}.GetObj().Count(proxy) //总的数据量
	}

	//count := models.Proxy{}.GetCount(proxy)
	number, _ := strconv.Atoi(per_page) //每页显示的数据条数
	compute := float64(count) / float64(number)
	countPage := math.Ceil(compute)                //计算总的页数
	startPosition := (nowPage - 1) * int32(number) //计算开始位置

	proxyList := make([]models.Proxy, 0)
	var errlist error
	if search != "" {
		errlist = models.Proxy{}.GetObj().Where("user_name = ? or addr = ? ", search).Limit(number, int(startPosition)).Find(&proxyList)
	} else {
		errlist = models.Proxy{}.GetObj().Limit(number, int(startPosition)).Find(&proxyList)
	}

	if errlist != nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Success"),
			"result": iris.Map{
				"data":         "",
				"current_page": nowPage,   //当前页码
				"last_page":    countPage, //总共多少页
				"per_page":     number,    //每页多少条数据
				"total":        count,     //总的记录条数
			},
		})
		return
	}

	//正常返回数据
	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result": iris.Map{
			"data":         proxyList,
			"current_page": nowPage,   //当前页码
			"last_page":    countPage, //总共多少页
			"per_page":     number,    //每页多少条数据
			"total":        count,     //总的记录条数
		},
	})
}

/**
	* @api {Post} /addProxy 添加代理服务器
	* @apiDescription 添加代理服务器
	* @apiGroup proxy
	* @apiPermission JWT
	* @apiVersion 1.0.0
	* @apiParam {String} locale 语言
	* @apiParam {String} token token
	* @apiParam {String} addr IP地址
	* @apiParam {Number} port 端口
	* @apiParam {String} user_name 用户名
	* @apiParam {String} password 密码
	* @apiSuccessExample {json} Success-Response:
	*  {
	*   "code": 0,
	*   "text": "操作成功",
	*   "result": "",
    *  }
	*
*/
func (Proxy) AddProxy(w context.Context) {
	addr := w.FormValue("addr")
	port := w.FormValue("port")
	user_name := w.FormValue("user_name")
	password := w.FormValue("password")
	newPort, _ := strconv.Atoi(port)
	//进行表单验证
	checkError := checkProxyForm(w)
	if len(checkError) > 0 {
		w.JSON(iris.Map{
			"code":    400,
			"message": checkError,
			"result":  "",
		})
		return
	}

	//进行数据写入操作
	proxy := new(models.Proxy)
	proxy.UserName = user_name
	proxy.Password = password
	proxy.Addr = addr
	proxy.Port = int32(newPort)
	proxy.LastUpdateDate = time.Now().Format("2006-01-02 15:04:05")
	insert := models.Proxy{}.Add(proxy)
	fmt.Println(insert)
	if insert <= 0 {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Failure"),
			"result":  "",
		})
		return
	}

	//操作成功
	body := map[string]interface{}{"type": 1, "adds": []int32{int32(insert)}, "dels": []int32{}, "updates": []int32{}}
	//异步通知MQ
	go White{}.PushMq(body)

	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  "",
	})

}

func checkProxyForm(w context.Context) map[string]map[string]string {
	addr := w.FormValue("addr")
	port := w.FormValue("port")
	user_name := w.FormValue("user_name")
	password := w.FormValue("password")

	checkError := make(map[string]map[string]string)
	addrError := make(map[string]string)
	portError := make(map[string]string)
	userNameError := make(map[string]string)
	passwordError := make(map[string]string)

	if addr == "" {
		addrError["required"] = lang.Lang{}.GetLang(w, "proxy", "AddrRRequired")
	} else {
		if ok := helper.CheckIp(addr); !ok {
			addrError["error"] = lang.Lang{}.GetLang(w, "proxy", "IpError")
		}
	}

	if port == "" {
		portError["required"] = lang.Lang{}.GetLang(w, "proxy", "PortRequired")
	}
	if user_name == "" {
		userNameError["required"] = lang.Lang{}.GetLang(w, "proxy", "UserNameRequired")
	}
	if password == "" {
		passwordError["required"] = lang.Lang{}.GetLang(w, "proxy", "PasswordRequired")
	}

	if len(addrError) > 0 {
		checkError["addr"] = addrError
	}
	if len(portError) > 0 {
		checkError["port"] = portError
	}
	if len(userNameError) > 0 {
		checkError["user_name"] = userNameError
	}
	if len(passwordError) > 0 {
		checkError["password"] = passwordError
	}
	return checkError
}

/**
* @api {Get} /proxy/{id} 修改时获取代理服务器信息
* @apiDescription 修改时获取代理服务器信息
* @apiGroup proxy
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiParam {String} token token
* @apiSuccessExample {json} Success-Response:
*	{
*	  "code": 0,
*	  "message": "操作成功",
*	  "result": {
*		"id": 2,
*		"addr": "192.168.31.155",
*		"port": 8080,
*		"user_name": "sunwukong",
*		"password": "123456",
*		"last_update_date": "0000-00-00 00:00:00"
*	  }
*	}
*
 */
func (Proxy) GetProxy(w context.Context) {
	id := w.Params().Get("id")
	newId, _ := strconv.Atoi(id)
	//判断数据是否存在
	proxy := new(models.Proxy)
	proxy.Id = int32(newId)
	oneProxy := models.Proxy{}.GetOne(proxy)
	if newId <= 0 || oneProxy == nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	//操作成功
	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  oneProxy,
	})
}

/**
	* @api {Put} /saveProxy/{id} 修改保存代理服务器
	* @apiDescription 修改保存代理服务器
	* @apiGroup proxy
	* @apiPermission JWT
	* @apiVersion 1.0.0
	* @apiParam {String} locale 语言
	* @apiParam {String} token token
	* @apiParam {String} addr IP地址
	* @apiParam {Number} port 端口
	* @apiParam {String} user_name 用户名
	* @apiParam {String} password 密码
	* @apiSuccessExample {json} Success-Response:
	*  {
	*   "code": 0,
	*   "text": "操作成功",
	*   "result": "",
    *  }
	*
*/
func (Proxy) SaveProxy(w context.Context) {
	id := w.Params().Get("id")
	newId, _ := strconv.Atoi(id)
	addr := w.FormValue("addr")
	port := w.FormValue("port")
	user_name := w.FormValue("user_name")
	password := w.FormValue("password")
	newPort, _ := strconv.Atoi(port)

	//判断数据是否存在
	proxy := new(models.Proxy)
	proxy.Id = int32(newId)
	oneProxy := models.Proxy{}.GetOne(proxy)
	if newId <= 0 || oneProxy == nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	//进行表单验证
	checkError := checkProxyForm(w)
	if len(checkError) > 0 {
		w.JSON(iris.Map{
			"code":    400,
			"message": checkError,
			"result":  "",
		})
		return
	}

	//进行数据修改操作
	updateProxy := new(models.Proxy)
	updateProxy.UserName = user_name
	updateProxy.Password = password
	updateProxy.Addr = addr
	updateProxy.Port = int32(newPort)
	updateProxy.LastUpdateDate = time.Now().Format("2006-01-02 15:04:05")

	update := models.Proxy{}.UpdateById(int32(newId), updateProxy)
	if !update {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Failure"),
			"result":  "",
		})
		return
	}

	//操作成功
	body := map[string]interface{}{"type": 1, "adds": []int32{}, "dels": []int32{int32(newId)}, "updates": []int32{}}
	//异步通知MQ
	go White{}.PushMq(body)

	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  "",
	})
}

/**
	* @api {Delete} /deleteProxy/{id} 删除代理服务器
	* @apiDescription 删除代理服务器
	* @apiGroup proxy
	* @apiPermission JWT
	* @apiVersion 1.0.0
	* @apiParam {String} locale 语言
	* @apiParam {String} token token
	* @apiSuccessExample {json} Success-Response:
	*  {
	*   "code": 0,
	*   "text": "操作成功",
	*   "result": "",
    *  }
	*
*/
func (Proxy) DeleteProxy(w context.Context) {
	id := w.Params().Get("id")
	newId, _ := strconv.Atoi(id)

	//判断数据是否存在
	proxy := new(models.Proxy)
	proxy.Id = int32(newId)
	oneProxy := models.Proxy{}.GetOne(proxy)
	if newId <= 0 || oneProxy == nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	//数据存在进行删除操作
	del := models.Proxy{}.Delete(proxy)
	if !del {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Failure"),
			"result":  "",
		})
		return
	}

	//操作成功

	body := map[string]interface{}{"type": 1, "adds": []int32{}, "dels": []int32{int32(newId)}, "updates": []int32{}}
	//异步通知MQ
	go White{}.PushMq(body)

	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  "",
	})
}
