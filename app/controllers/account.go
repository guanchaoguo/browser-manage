package controllers

import (
	"browser-manage/app/helper"
	"browser-manage/app/models"
	"browser-manage/lang"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"math"
	"regexp"
	"strconv"
	"time"
)

/*
	后台用户管理操作
*/
type Account struct {
}

/**
* @api {Get} /accountList 账号列表
* @apiDescription 账号列表
* @apiGroup account
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
*		  {
*			"id": 1,
*			"user_name": "admin",
*			"password": "fd3c1140e2660132dc39e8788884cc8d",
*			"salt": "123456",
*			"account": "开发用户",
*			"last_date": "0000-00-00 00:00:00",
*			"last_login_date": "2017-11-07 16:13:17",
*			"last_login_ip": "127.0.0.1"
*		  },
*		],
*		"last_page": 1,
*		"per_page": 15,
*		"total": 7
*	  }
*	}
*
 */
func (Account) List(w context.Context) {
	search := w.FormValue("search")
	per_page := w.URLParamDefault("per_page", "15")
	page := w.URLParamDefault("page", "1")
	p, _ := strconv.Atoi(page)
	nowPage := int32(p)

	//查询条件组装
	user := new(models.User)
	var count int64
	//分页结算
	if search != "" {
		count, _ = models.User{}.GetObj().Where("user_name =? or account =? ", search, search).Count(user)
	} else {
		count, _ = models.User{}.GetObj().Count(user)
	}

	//count := models.User{}.GetCount(user) //总的数据量
	number, _ := strconv.Atoi(per_page) //每页显示的数据条数
	compute := float64(count) / float64(number)
	countPage := math.Ceil(compute)                //计算总的页数
	startPosition := (nowPage - 1) * int32(number) //计算开始位置

	//获取数据
	list := make([]models.User, 0)
	var errlist error
	if search != "" {
		errlist = models.User{}.GetObj().Where("user_name =? or account = ? ", search, search).Limit(number, int(startPosition)).Find(&list)
	} else {
		errlist = models.User{}.GetObj().Limit(number, int(startPosition)).Find(&list)
	}

	//数据为空
	if errlist != nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Failure"),
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
			"data":         list,
			"current_page": nowPage,   //当前页码
			"last_page":    countPage, //总共多少页
			"per_page":     number,    //每页多少条数据
			"total":        count,     //总的记录条数
		},
	})
}

/**
	* @api {Post} /addAccount 添加账号
	* @apiDescription 添加账号
	* @apiGroup account
	* @apiPermission JWT
	* @apiVersion 1.0.0
	* @apiParam {String} locale 语言
	* @apiParam {String} token token
	* @apiParam {String} user_name 用户名
	* @apiParam {String} password 密码
	* @apiParam {String} account 真实姓名
	* @apiSuccessExample {json} Success-Response:
	*  {
	*   "code": 0,
	*   "text": "操作成功",
	*   "result": "",
    *  }
	*
*/
func (Account) AddAccount(w context.Context) {
	user_name := w.FormValue("user_name")
	password := w.FormValue("password")
	account := w.FormValue("account")

	//进行表单验证
	verifyError := checkUserForm(w, "add")

	//验证不通过
	if len(verifyError) != 0 {
		w.JSON(iris.Map{
			"code":    400,
			"message": verifyError,
			"result":  "",
		})
		return
	}

	//验证通过进行密码生成
	salt := helper.GetRandomString(10)
	newPassword := password + salt
	h := sha1.New()
	h.Write([]byte(newPassword))
	enPassword := hex.EncodeToString(h.Sum(nil))
	m := md5.New()
	m.Write([]byte(enPassword))
	md5Password := hex.EncodeToString(m.Sum(nil))

	//验证通过进行添加操作
	insertData := new(models.User)
	insertData.UserName = user_name
	insertData.Password = md5Password
	insertData.Account = account
	insertData.Salt = salt
	insertData.LastDate = time.Now().Format("2006-01-02 15:04:05")
	insertData.LastLoginDate = "0000-00-00 00:00:00"
	add := models.User{}.Add(insertData)

	//操作失败
	if add <= 0 {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Failure"),
			"result":  "",
		})
		return
	}

	//操作成功
	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  "",
	})
}

/*
	表单验证
*/
func checkUserForm(w context.Context, action string) map[string]map[string]string {
	user_name := w.FormValue("user_name")
	password := w.FormValue("password")
	account := w.FormValue("account")

	verifyError := make(map[string]map[string]string)
	userNameError := make(map[string]string)
	passwordError := make(map[string]string)
	accountError := make(map[string]string)

	if user_name == "" {
		userNameError["required"] = lang.Lang{}.GetLang(w, "account", "UsernameRequired")
	}
	//修改操作时不进行密码验证
	if action == "add" {
		if password == "" {
			passwordError["required"] = lang.Lang{}.GetLang(w, "account", "CaptchaRequired")
		}
		//验证密码长度，最少为6位长度
		if len(password) < 6 {
			passwordError["length"] = lang.Lang{}.GetLang(w, "account", "PasswordLength")
		}
	}

	if account == "" {
		accountError["required"] = lang.Lang{}.GetLang(w, "account", "AccountRequired")
	}

	//验证用户名的长度和格式（大小写字母 + 数字组合）
	layout, _ := regexp.MatchString("^[0-9a-zA-Z]{6,50}$", user_name)
	if len(user_name) > 50 || len(user_name) < 6 || !layout {
		userNameError["length"] = lang.Lang{}.GetLang(w, "account", "UserNameLength")
	}

	//验证用户名数据库中是否已经存在
	if user_name != "" {
		user := new(models.User)
		user.UserName = user_name
		id := w.Params().Get("id")
		newId, _ := strconv.Atoi(id)
		notIn := make([]int, 0)
		notInList := append(notIn, newId)
		oneFind, _ := models.User{}.GetAccountByOne(user, notInList)
		if oneFind != nil && oneFind.UserName != "" {
			userNameError["exist"] = lang.Lang{}.GetLang(w, "account", "UserNameExist")
		}
	}

	if len(userNameError) > 0 {
		verifyError["user_name"] = userNameError
	}
	if len(passwordError) > 0 {
		verifyError["password"] = passwordError
	}
	if len(accountError) > 0 {
		verifyError["account"] = accountError
	}
	return verifyError
}

/**
* @api {Get} /showAccount/{id} 修改时获取账号信息
* @apiDescription 修改时获取账号信息
* @apiGroup account
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiParam {String} token token
* @apiSuccessExample {json} Success-Response:
*	{
*	  "code": 400,
*	  "message": "操作成功",
*	  "result": {
*		"data": {
*		  "account": "黄飞鸿",
*		  "id": 2,
*		  "user_name": "huang"
*		}
*	  }
*	}
*
 */
func (Account) GetAccount(w context.Context) {
	id := w.Params().Get("id")
	newId, _ := strconv.Atoi(id)

	user := new(models.User)
	user.Id = int32(newId)
	oneUser, err := models.User{}.GetAccountByOne(user, make([]int, 0))
	//用户不存在，证明参数错误
	if err != nil || oneUser.UserName == "" {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	//正常返回数据，只返回需要的数据
	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result": iris.Map{
			"data": iris.Map{
				"id":        oneUser.Id,
				"user_name": oneUser.UserName,
				"account":   oneUser.Account,
			},
		},
	})
}

/**
	* @api {Put} /saveAccount/{id} 修改保存账号
	* @apiDescription 修改保存账号
	* @apiGroup account
	* @apiPermission JWT
	* @apiVersion 1.0.0
	* @apiParam {String} locale 语言
	* @apiParam {String} token token
	* @apiParam {String} user_name 用户名
	* @apiParam {String} password 密码
	* @apiParam {String} account 真实姓名
	* @apiSuccessExample {json} Success-Response:
	*  {
	*   "code": 0,
	*   "text": "操作成功",
	*   "result": "",
    *  }
	*
*/
func (Account) SaveAccount(w context.Context) {
	id := w.Params().Get("id")
	user_name := w.FormValue("user_name")
	account := w.FormValue("account")

	//进行表单验证
	verifyError := checkUserForm(w, "update")
	//验证不通过
	if len(verifyError) != 0 {
		w.JSON(iris.Map{
			"code":    400,
			"message": verifyError,
			"result":  "",
		})
		return
	}

	//判断修改的数据是否存在
	saveId, _ := strconv.Atoi(id)
	user := new(models.User)
	user.Id = int32(saveId)
	oneUser, err := models.User{}.GetAccountByOne(user, make([]int, 0))
	if err != nil || oneUser == nil || oneUser.UserName == "" {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	//验证通过进行修改操作
	updateData := new(models.User)
	updateData.UserName = user_name
	//updateData.Password = oneUser.Password
	updateData.Account = account
	updateData.LastDate = time.Now().Format("2006-01-02 15:04:05")

	upd := models.User{}.UpdateById(int32(saveId), updateData)

	//操作失败
	if !upd {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Failure"),
			"result":  "",
		})
	}

	//操作成功
	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  "",
	})

}

/**
	* @api {Put} /savePwd/{id} 修改账户密码
	* @apiDescription 修改账户密码
	* @apiGroup account
	* @apiPermission JWT
	* @apiVersion 1.0.0
	* @apiParam {String} locale 语言
	* @apiParam {String} token token
	* @apiParam {String} password 密码
	* @apiSuccessExample {json} Success-Response:
	*  {
	*   "code": 0,
	*   "text": "操作成功",
	*   "result": "",
    *  }
	*
*/
func (Account) SavePwd(w context.Context) {
	pwd := w.FormValue("password")
	id := w.Params().Get("id")

	deleteId, _ := strconv.Atoi(id)
	//参数错误
	if deleteId <= 0 {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	//查看参数条件数据是否存在
	user := new(models.User)
	user.Id = int32(deleteId)
	findOne, err := models.User{}.GetAccountByOne(user, make([]int, 0))
	if err != nil || findOne == nil || findOne.UserName == "" {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	//验证密码长度
	if len(pwd) < 6 {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "account", "PasswordLength"),
			"result":  "",
		})
		return
	}

	//重新生成密码和盐值
	//验证通过进行密码生成
	salt := helper.GetRandomString(10)
	newPassword := pwd + salt
	h := sha1.New()
	h.Write([]byte(newPassword))
	enPassword := hex.EncodeToString(h.Sum(nil))
	m := md5.New()
	m.Write([]byte(enPassword))
	md5Password := hex.EncodeToString(m.Sum(nil))

	saveId, _ := strconv.Atoi(id)
	user = new(models.User)
	user.Password = md5Password
	user.Salt = salt
	save := models.User{}.UpdateById(int32(saveId), user)

	if !save {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Failure"),
			"result":  "",
		})
		return
	}

	//操作成功
	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  "",
	})
}

/**
	* @api {Delete} /deleteAccount/{id} 删除账号
	* @apiDescription 删除账号
	* @apiGroup account
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
func (Account) DeleteAccount(w context.Context) {
	id := w.Params().Get("id")
	deleteId, _ := strconv.Atoi(id)
	//参数错误
	if deleteId <= 0 {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	//查看参数条件数据是否存在
	user := new(models.User)
	user.Id = int32(deleteId)
	findOne, err := models.User{}.GetAccountByOne(user, make([]int, 0))
	if err != nil || findOne == nil || findOne.UserName == "" {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	//进行删除操作
	del := models.User{}.DeleteById(findOne.Id)

	//操作失败
	if del <= 0 {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Failure"),
			"result":  "",
		})
		return
	}

	//操作成功
	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  "",
	})
	return
}
