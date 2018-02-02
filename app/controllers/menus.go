package controllers

import (
	"browser-manage/app/models"
	"browser-manage/lang"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"strconv"
	"time"
)

/*
	后台菜单管理
*/
type Menus struct {
}

/**
* @api {Get} /menusList 菜单列表
* @apiDescription 菜单列表
* @apiGroup menus
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiParam {String} token token
* @apiSuccessExample {json} Success-Response:
	{
	  "code": 0,
	  "message": "操作成功",
	  "result": {
		"data": [
		  {
			"id": 1,
			"parent_id": 0,
			"title_cn": "abc",
			"title_en": "",
			"class": 0,
			"desc": "",
			"link_url": "",
			"icon": "",
			"state": 1,
			"sort_id": 1,
			"menu_code": "",
			"_child": [
			  {
				"id": 2,
				"parent_id": 1,
				"title_cn": "cba",
				"title_en": "",
				"class": 0,
				"desc": "",
				"link_url": "",
				"icon": "",
				"state": 1,
				"sort_id": 1,
				"menu_code": "",
				"_child": [
				  {
					"id": 3,
					"parent_id": 2,
					"title_cn": "bnm",
					"title_en": "",
					"class": 0,
					"desc": "",
					"link_url": "",
					"icon": "",
					"state": 1,
					"sort_id": 1,
					"menu_code": "",
					"_child": []
				  }
				]
			  }
			]
		  }
		]
	  }
	}
*
*/
func (Menus) List(w context.Context) {
	list := make([]models.Menus, 0)
	menus := new(models.Menus)
	listMenus := models.Menus{}.List(list, menus)
	if listMenus == nil {
		w.JSON(iris.Map{
			"code":    0,
			"message": lang.Lang{}.GetLang(w, "comm", "Success"),
			"result": iris.Map{
				"data": "",
			},
		})
	}

	//进行菜单等级组装
	treeList := TreeMenues(listMenus, 0)
	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result": iris.Map{
			"data": treeList,
		},
	})

}

/*
	递归整理树形菜单
*/
func TreeMenues(list []models.Menus, parent_id int32) interface{} {
	type newList struct {
		models.Menus
		Child interface{} `json:"_child"`
	}
	tree := make([]newList, 0)
	for _, item := range list {
		var newItem newList
		if item.ParentId == parent_id {
			newItem.ParentId = item.ParentId
			newItem.Id = item.Id
			newItem.LinkUrl = item.LinkUrl
			newItem.MenuCode = item.MenuCode
			newItem.TitleCn = item.TitleCn
			newItem.Icon = item.Icon
			newItem.SortId = item.SortId
			newItem.State = item.State
			newItem.Desc = item.Desc
			newItem.Class = item.Class
			newItem.Child = TreeMenues(list, item.Id)
			tree = append(tree, newItem)
		}

	}
	return tree
}

/**
	* @api {Post} /addMenus 添加菜单
	* @apiDescription 添加菜单
	* @apiGroup menus
	* @apiPermission JWT
	* @apiVersion 1.0.0
	* @apiParam {String} locale 语言
	* @apiParam {String} token token
	* @apiParam {String} title_cn 菜单名称
	* @apiParam {String} link_url 菜单链接
	* @apiParam {String} menu_code 菜单标识符
	* @apiParam {String} icon 菜单图标
	* @apiParam {Number} sort_id 排序 （越小越前）
	* @apiParam {Number} parent_id 父级菜单ID
	* @apiSuccessExample {json} Success-Response:
	*  {
	*   "code": 0,
	*   "text": "操作成功",
	*   "result": "",
    *  }
	*
*/
func (Menus) Add(w context.Context) {
	//进行表单验证
	checkform := checkform(w)
	if checkform != nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": checkform,
			"result":  "",
		})
		return
	}

	title_cn := w.FormValue("title_cn")
	link_url := w.FormValue("link_url")
	menu_code := w.FormValue("menu_code")
	icon := w.FormValue("icon")
	parent_id := w.FormValue("parent_id")
	newParentId, _ := strconv.Atoi(parent_id)
	sortId := w.FormValue("sort_id")
	newSortId, _ := strconv.Atoi(sortId)
	menus := new(models.Menus)
	menus.Icon = icon
	menus.TitleCn = title_cn
	menus.MenuCode = menu_code
	menus.ParentId = int32(newParentId)
	menus.LinkUrl = link_url
	menus.State = 1
	menus.SortId = int32(newSortId)
	menus.UpdateDate = time.Now().Format("2006-01-02 15:04:05")

	//进行添加操作
	aff := models.Menus{}.Add(menus)
	if aff == 0 {
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
	* @api {Put} /saveMenus/{id} 修改保存菜单
	* @apiDescription 修改保存菜单
	* @apiGroup menus
	* @apiPermission JWT
	* @apiVersion 1.0.0
	* @apiParam {String} locale 语言
	* @apiParam {String} token token
	* @apiParam {String} title_cn 菜单名称
	* @apiParam {String} link_url 菜单链接
	* @apiParam {String} menu_code 菜单标识符
	* @apiParam {String} icon 菜单图标
	* @apiParam {Number} parent_id 父级菜单ID
	* @apiParam {Number} sort_id 排序 （越小越前）
	* @apiSuccessExample {json} Success-Response:
	*  {
	*   "code": 0,
	*   "text": "操作成功",
	*   "result": "",
    *  }
	*
*/
func (Menus) Save(w context.Context) {
	id := w.Params().Get("id")
	newId, _ := strconv.Atoi(id)

	//验证数据是否存在
	whereMenus := new(models.Menus)
	whereMenus.Id = int32(newId)
	oneMenu := models.Menus{}.GetOne(whereMenus)
	//数据不存在，参数错误
	if oneMenu == nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	//进行表单验证
	checkform := checkform(w)
	if checkform != nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": checkform,
			"result":  "",
		})
		return
	}

	title_cn := w.FormValue("title_cn")
	link_url := w.FormValue("link_url")
	menu_code := w.FormValue("menu_code")
	icon := w.FormValue("icon")
	parent_id := w.URLParamDefault("parent_id", "0")
	newParentId, _ := strconv.Atoi(parent_id)
	sortId := w.FormValue("sort_id")
	newSortId, _ := strconv.Atoi(sortId)

	menus := new(models.Menus)
	menus.Icon = icon
	menus.TitleCn = title_cn
	menus.MenuCode = menu_code
	menus.ParentId = int32(newParentId)
	menus.LinkUrl = link_url
	menus.State = 1
	menus.SortId = int32(newSortId)
	menus.UpdateDate = time.Now().Format("2006-01-02 15:04:05")

	//进行修改操作
	aff := models.Menus{}.UpdateById(int32(newId), menus)
	if !aff {
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
func checkform(w context.Context) map[string]map[string]string {
	title_cn := w.FormValue("title_cn")
	link_url := w.FormValue("link_url")
	menu_code := w.FormValue("menu_code")
	icon := w.FormValue("icon")

	checkError := make(map[string]map[string]string)
	titleError := make(map[string]string)
	linkError := make(map[string]string)
	codeError := make(map[string]string)
	iconError := make(map[string]string)

	if title_cn == "" {
		titleError["required"] = lang.Lang{}.GetLang(w, "menus", "TitleRequired")
		checkError["title_cn"] = titleError
	}
	if link_url == "" {
		linkError["required"] = lang.Lang{}.GetLang(w, "menus", "LinkRequired")
		checkError["link_url"] = linkError
	}
	if menu_code == "" {
		codeError["required"] = lang.Lang{}.GetLang(w, "menus", "CodeRequired")
		checkError["menu_code"] = codeError
	}
	if icon == "" {
		iconError["required"] = lang.Lang{}.GetLang(w, "menus", "IconRequired")
		checkError["icon"] = iconError
	}

	//返回值
	if len(checkError) > 0 {
		return checkError
	}
	return nil

}

/**
* @api {Get} /getMenus/{id} 修改时获取菜单信息
* @apiDescription 修改时获取菜单信息
* @apiGroup menus
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiParam {String} token token
* @apiSuccessExample {json} Success-Response:
	{
	  "code": 0,
	  "message": "操作成功",
	  "result": {
		"id": 112,
		"parent_id": 0,
		"title_cn": "系统管理",
		"title_en": "",
		"class": 0,
		"desc": "",
		"link_url": "/systemManage",
		"icon": "icon-xitongguanli",
		"state": 0,
		"sort_id": 0,
		"menu_code": "M4001",
		"update_date": "2017-11-14 10:09:24"
	  }
	}
*
*/
func (Menus) Show(w context.Context) {
	id := w.Params().Get("id")
	newId, _ := strconv.Atoi(id)

	//查看数据是否存在
	menus := new(models.Menus)
	menus.Id = int32(newId)
	oneMenus := models.Menus{}.GetOne(menus)
	if oneMenus == nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}
	//正常返回数据
	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  oneMenus,
	})
}

/**
	* @api {Delete} /deleteMenus/{id} 删除菜单
	* @apiDescription 删除菜单
	* @apiGroup menus
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
func (Menus) Delete(w context.Context) {
	id := w.Params().Get("id")
	newId, _ := strconv.Atoi(id)

	//查看数据是否存在
	menus := new(models.Menus)
	menus.Id = int32(newId)
	oneMenus := models.Menus{}.GetOne(menus)
	if oneMenus == nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	//进行正常删除操作
	del := models.Menus{}.DeleteById(int32(newId), menus)
	if !del {
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
