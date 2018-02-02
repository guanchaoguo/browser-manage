package controllers

import (
	"browser-manage/app/helper"
	"browser-manage/app/models"
	"browser-manage/lang"
	"encoding/json"
	"github.com/Luxurioust/excelize"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/streadway/amqp"
	"io"
	"log"
	"math"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
	白名单管理
*/
type White struct {
}

/*
	mq通知消息结构体
*/
type PushMq struct {
	Type    int     `json:"type"`
	Adds    []int32 `json:"adds"`
	Dels    []int32 `json:"dels"`
	Updates []int32 `json:"updates"`
}

/**
* @api {Get} /whiteList 白名单列表
* @apiDescription 白名单列表
* @apiGroup white
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiParam {String} token token
* @apiParam {String} search 输入框条件
* @apiParam {Number} channel 查询模式 1为IP，2为代理
* @apiParam {Number} status 状态，2为锁定，1为启用状态，默认为0（全部）
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
*			"id": 55,
*			"domain": "www.alibaba.com",
*			"hall_name": "马云",
*			"status": 1, //启用状态，2为锁定，1为启用状态，默认为1
*			"channel": 1, //1：走IP(默认)，2：走代理
*			"ips" : "",  //IP字段，用英文半角分号隔开
*			"create_date": "0000-00-00 00:00:00" //创建时间
*		  },
*		],
*		"last_page": 1,
*		"per_page": 15,
*		"total": 2
*	  }
*	}
 */
func (White) List(w context.Context) {
	search := w.FormValue("search")
	per_page := w.URLParamDefault("per_page", "15")
	status := w.URLParamDefault("status", "0")
	newStatus, _ := strconv.Atoi(status)
	channel := w.FormValue("channel")
	newChannel, _ := strconv.Atoi(channel)
	page := w.URLParamDefault("page", "1")
	newPage, _ := strconv.Atoi(page)
	nowPage := int32(newPage)

	white := new(models.White)

	//分页结算
	var whiteSession *xorm.Session
	if search != "" {
		whiteSession = models.White{}.GetObj().Where("domain =? or hall_name=? or ips like ?", search, search, "%"+search+"%")
	} else {
		whiteSession = models.White{}.GetObj().Where("")
	}

	if newStatus > 0 && newStatus < 3 {
		whiteSession.And("status =? ", newStatus)
	}
	if newChannel > 0 && newChannel < 3 {
		whiteSession.And("channel=?", newChannel)
	}
	count, _ := whiteSession.Count(white) //总的数据量
	number, _ := strconv.Atoi(per_page)   //每页显示的数据条数
	compute := float64(count) / float64(number)
	countPage := math.Ceil(compute)                //计算总的页数
	startPosition := (nowPage - 1) * int32(number) //计算开始位置

	whiteList := make([]models.White, 0)
	var whiteSessionList *xorm.Session
	if search != "" {
		whiteSessionList = models.White{}.GetObj().Where("domain =? or hall_name=? or ips like ?", search, search, "%"+search+"%")
	} else {
		whiteSessionList = models.White{}.GetObj().Where("")
	}

	if newStatus > 0 && newStatus < 3 {
		whiteSessionList.And("status = ?", newStatus)
	}
	if newChannel > 0 && newChannel < 3 {
		whiteSessionList.And("channel = ?", newChannel)
	}
	listErr := whiteSessionList.Desc("id").Limit(number, int(startPosition)).Find(&whiteList)
	//whiteList,err := models.White{}.List(list,white,int32(number),startPosition)
	if listErr != nil {
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
			"data":         whiteList,
			"current_page": nowPage,   //当前页码
			"last_page":    countPage, //总共多少页
			"per_page":     number,    //每页多少条数据
			"total":        count,     //总的记录条数
		},
	})
}

/**
	* @api {Post} /addWhite 添加白名单
	* @apiDescription 添加白名单
	* @apiGroup white
	* @apiPermission JWT
	* @apiVersion 1.0.0
	* @apiParam {String} locale 语言
	* @apiParam {String} token token
	* @apiParam {String} domain 域名
	* @apiParam {String} hall_name 所属人
	* @apiParam {String} ips IP
	* @apiParam {Number} channel 1：走IP(默认)，2：走代理
	* @apiParam {Number} status 1：启用(默认)，2：锁定
	* @apiSuccessExample {json} Success-Response:
	*  {
	*   "code": 0,
	*   "text": "操作成功",
	*   "result": "",
    *  }
	*
*/
func (White) AddWhite(w context.Context) {
	domain := w.FormValue("domain")
	hall_name := w.FormValue("hall_name")
	ips := w.FormValue("ips")
	if len(ips) > 0 {
		ips = IpReplace(ips) //替换和去除空格
	}
	formChannel := w.FormValue("channel")
	channel, _ := strconv.Atoi(formChannel)
	formStatus := w.FormValue("status")
	status, _ := strconv.Atoi(formStatus)
	if channel != 1 && channel != 2 {
		channel = 1
	}
	if status != 1 && status != 2 {
		status = 1
	}

	//先进行表单数据验证
	checkError := checkForm(w)
	if len(checkError) > 0 {
		w.JSON(iris.Map{
			"code":    400,
			"message": checkError,
			"result":  "",
		})
		return
	}

	//验证通过进行写入操作
	insertData := new(models.White)
	insertData.Domain = domain
	insertData.HallName = hall_name
	insertData.Status = int32(status)
	insertData.Ips = ips
	insertData.Channel = int32(channel)
	insertData.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	insertData.LastUpdateDate = "0000-00-00 00:00:00"
	aff := models.White{}.Add(insertData)

	//操作失败
	if aff <= 0 {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Failure"),
			"result":  "",
		})
		return
	}
	//操作成功
	body := map[string]interface{}{"type": 0, "adds": []int32{int32(aff)}, "dels": []int32{}, "updates": []int32{}}
	//异步通知MQ
	go White{}.PushMq(body)

	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  "",
	})
}

//IP处理函数，把全角符号替换成半角符号，并且去掉两端的空格
func IpReplace(str string) string {
	str = strings.Replace(str, "；", ";", -1) //替换全角的分号
	str = strings.Trim(str, " ")             //去除两端的空格
	return str
}

//IP根据指定字符串切割，返回数组
func IpSplit(ipStr string) []string {
	ipSlices := strings.Split(ipStr, ";")
	return ipSlices
}

//验证是否为IP格式
func checkIp(ip string) bool {
	layout, err := regexp.MatchString("(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)\\.(25[0-5]|2[0-4]\\d|[0-1]\\d{2}|[1-9]?\\d)", ip)
	if !layout || err != nil {
		return false
	}
	return true
}

//表单验证
func checkForm(w context.Context) map[string]map[string]string {
	domain := w.FormValue("domain")
	hall_name := w.FormValue("hall_name")
	ips := w.FormValue("ips")
	var newIps []string
	if len(ips) > 0 {
		ips = IpReplace(ips)  //符号转换
		newIps = IpSplit(ips) //IP字段截取
	}
	checkError := make(map[string]map[string]string)
	domainError := make(map[string]string)
	hallNameError := make(map[string]string)
	ipsError := make(map[string]string)

	if domain == "" {
		domainError["required"] = lang.Lang{}.GetLang(w, "white", "DomainRequired")
	}
	if len(domain) > 0 {
		//验证域名是否已经存在
		id := w.Params().Get("id")
		newId, _ := strconv.Atoi(id)
		notIn := make([]int, 0)
		notInList := append(notIn, newId)
		white := new(models.White)
		white.Domain = domain
		isHas := models.White{}.GetOne(white, notInList)
		if isHas != nil {
			domainError["exist"] = lang.Lang{}.GetLang(w, "white", "DomainExist")
		}
	}

	if hall_name == "" {
		hallNameError["required"] = lang.Lang{}.GetLang(w, "white", "HallNameRequired")
	}
	//验证IP格式
	if len(newIps) > 0 {
		for _, item := range newIps {
			//校验是否为IP格式
			if ok := checkIp(item); !ok {
				ipsError["ipError"] = lang.Lang{}.GetLang(w, "white", "IpError")
			}
			break
		}
	}

	//验证错误组装
	if len(domainError) > 0 {
		checkError["domain"] = domainError
	}

	if len(hallNameError) > 0 {
		checkError["hall_name"] = hallNameError
	}

	if len(ipsError) > 0 {
		checkError["ips"] = ipsError
	}
	return checkError
}

/**
* @api {Get} /showWhite/{id} 修改时获取白名单信息
* @apiDescription 修改时获取白名单信息
* @apiGroup white
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiParam {String} token token
* @apiSuccessExample {json} Success-Response:
*	{
*	  "code": 0,
*	  "message": "操作成功",
*	  "result": {
*		"id": 54,
*		"domain": "www.baidu002.com",
*		"hall_name": "baidu002",
*		"status": 1,
*		"channel": 1,
*		"ips" : ""
*	  }
*	}
*
 */
func (White) ShowWhite(w context.Context) {
	id := w.Params().Get("id")
	newId, _ := strconv.Atoi(id)

	//验证数据是否存在
	white := new(models.White)
	white.Id = int32(newId)
	oneWhite := models.White{}.GetOne(white, make([]int, 0))
	if oneWhite == nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}
	//返回数据
	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  oneWhite,
	})
}

/**
	* @api {Put} /saveWhite/{id} 编辑保存白名单
	* @apiDescription 编辑保存白名单
	* @apiGroup white
	* @apiPermission JWT
	* @apiVersion 1.0.0
	* @apiParam {String} locale 语言
	* @apiParam {String} token token
	* @apiParam {String} domain 域名
	* @apiParam {String} hall_name 所属人
	* @apiParam {String} ips IP
	* @apiParam {Number} channel 1：走IP(默认)，2：走代理
	* @apiSuccessExample {json} Success-Response:
	*  {
	*   "code": 0,
	*   "text": "操作成功",
	*   "result": "",
    *  }
	*
*/
func (White) SaveWhite(w context.Context) {
	domain := w.FormValue("domain")
	hall_name := w.FormValue("hall_name")
	id := w.Params().Get("id")
	newId, _ := strconv.Atoi(id)
	ips := w.FormValue("ips")
	ips = IpReplace(ips) //替换和去除空格
	formChannel := w.FormValue("channel")
	channel, _ := strconv.Atoi(formChannel)

	if channel != 1 && channel != 2 {
		channel = 1
	}

	//验证数据是否存在
	white := new(models.White)
	white.Id = int32(newId)
	oneWhite := models.White{}.GetOne(white, make([]int, 0))
	if oneWhite == nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	//先进行表单数据验证
	checkError := checkForm(w)
	if len(checkError) > 0 {
		w.JSON(iris.Map{
			"code":    400,
			"message": checkError,
			"result":  "",
		})
		return
	}

	//验证通过进行写入操作
	updateData := new(models.White)
	updateData.Domain = domain
	updateData.HallName = hall_name
	updateData.Ips = ips
	updateData.Channel = int32(channel)
	updateData.LastUpdateDate = time.Now().Format("2006-01-02 15:04:05")
	aff := models.White{}.UpdateById(int32(newId), updateData)

	//操作失败
	if !aff {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Failure"),
			"result":  "",
		})
		return
	}
	//操作成功
	body := map[string]interface{}{"type": 0, "adds": []int32{}, "dels": []int32{}, "updates": []int32{int32(newId)}}
	//异步通知MQ
	go White{}.PushMq(body)

	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  "",
	})
}

/**
	* @api {Delete} /deleteWhite/{id} 删除白名单
	* @apiDescription 删除白名单
	* @apiGroup white
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
func (White) DeleteWhite(w context.Context) {
	id := w.Params().Get("id")
	newId, _ := strconv.Atoi(id)

	//验证数据是否存在
	white := new(models.White)
	white.Id = int32(newId)
	oneWhite := models.White{}.GetOne(white, make([]int, 0))
	if oneWhite == nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	//进行删除操作
	del := models.White{}.Delete(white)
	//操作失败
	if !del {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Failure"),
			"result":  "",
		})
		return
	}
	//操作成功

	body := map[string]interface{}{"type": 0, "adds": []int32{}, "dels": []int32{int32(newId)}, "updates": []int32{}}
	//异步通知MQ
	go White{}.PushMq(body)

	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  "",
	})
}

/**
	* @api {Put} /saveWhiteStatus/{id} 修改白名单状态
	* @apiDescription 修改白名单状态
	* @apiGroup white
	* @apiPermission JWT
	* @apiVersion 1.0.0
	* @apiParam {String} locale 语言
	* @apiParam {String} token token
	* @apiParam {Number} status 状态值，2为锁定，1为启用状态
	* @apiParam {String} lock_remark 锁定原因
	* @apiSuccessExample {json} Success-Response:
	*  {
	*   "code": 0,
	*   "text": "操作成功",
	*   "result": "",
    *  }
	*
*/
func (White) SaveWhiteStatus(w context.Context) {
	status := w.FormValue("status")
	lock_remark := w.FormValue("lock_remark")
	id := w.Params().Get("id")
	newId, _ := strconv.Atoi(id)
	newStatus, _ := strconv.Atoi(status)

	//验证数据是否存在
	white := new(models.White)
	white.Id = int32(newId)
	oneWhite := models.White{}.GetOne(white, make([]int, 0))
	if oneWhite == nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Invalid"),
			"result":  "",
		})
		return
	}

	//进行状态更新操作
	statusData := new(models.White)
	statusData.Status = int32(newStatus)
	statusData.LockRemark = lock_remark
	save := models.White{}.SaveStatusById(int32(newId), statusData)
	if !save {
		w.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(w, "comm", "Failure"),
			"result":  "",
		})
		return
	}

	//操作成功

	//操作成功
	body := map[string]interface{}{"type": 0, "adds": []int32{}, "dels": []int32{}, "updates": []int32{int32(newId)}}
	//异步通知MQ
	go White{}.PushMq(body)

	w.JSON(iris.Map{
		"code":    0,
		"message": lang.Lang{}.GetLang(w, "comm", "Success"),
		"result":  "",
	})
}

/**
* @api {Post} /importWhite 导入白名单
* @apiDescription 导入白名单
* @apiGroup white
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiParam {String} token token
* @apiParam {File} uploadfile 上传文件
* @apiSuccessExample {json} Success-Response:
*	{
*	  "code": 0,
*	  "message": "操作成功! 总共 3 条记录！ 成功导入 0 条记录; 失败 3 条记录！",
*	  "result": ""
*	}
*
 */
func (White) Import(w iris.Context) {
	fileName, newFileName := White{}.Upload(w)
	if fileName == "" {
		return
	}
	xlsx, err := excelize.OpenFile(fileName)
	//defer os.Remove(fileName)

	if err != nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": "操作错误",
			"result":  "",
		})
		return
	}

	insertData := make([]models.White, 0)
	var mData models.White
	index := xlsx.GetSheetIndex("Sheet1")
	// Get all the rows in a sheet.
	rows := xlsx.GetRows("Sheet" + strconv.Itoa(index))

	//获取所有数据库中已有的域名
	whiteMap := make(map[string]string)
	whiteList := make([]models.White, 0)
	engine := models.White{}.GetEngine()
	engine.Cols("domain").Find(&whiteList)
	if len(whiteList) > 0 {
		for _, item := range whiteList {
			whiteMap[item.Domain] = item.Domain
		}
	}
	for k, row := range rows {
		if k != 0 {
			var errString string
			for kk, colCell := range row {
				if kk == 0 {
					//验证域名格式
					layout, _ := regexp.MatchString("^([a-zA-Z0-9]([a-zA-Z0-9\\-]{0,61}[a-zA-Z0-9])?\\.)+[a-zA-Z]{2,6}$", colCell)
					if !layout {
						errString = lang.Lang{}.GetLang(w, "white", "DomainFormatError")
						break
					} else if _, is_ok := whiteMap[colCell]; is_ok {
						errString = lang.Lang{}.GetLang(w, "white", "DomainExist")
						break
					} else {
						mData.Domain = colCell
					}

				} else if kk == 1 { //所属人
					mData.HallName = colCell
				} else if kk == 2 { //IP
					if len(colCell) > 0 {
						//进行符号的替换和切割
						ips := IpReplace(colCell)
						newips := IpSplit(ips)
						//验证IP格式是否正确
						ipformat := true
						for _, v2 := range newips {
							if oks := checkIp(v2); !oks {
								errString = lang.Lang{}.GetLang(w, "white", "IpError")
								ipformat = false
								break
							}
						}
						if !ipformat {
							break
						} else {
							mData.Ips = ips
						}
					} else {
						mData.Ips = ""
					}

				} else if kk == 3 { //反劫持模式
					if ok := strings.Contains(colCell, "IP"); ok {
						mData.Channel = 1
					} else {
						mData.Channel = 2
					}
				}
				mData.Status = 1
			}
			mData.CreateDate = time.Now().Format("2006-01-02 15:04:05")
			mData.LastUpdateDate = "0000-00-00 00:00:00"
			if len(errString) <= 0 {
				insertData = append(insertData, mData)
			}

			//状态标识
			axis := "E" + strconv.Itoa(k+1)
			if len(errString) > 0 {
				xlsx.SetCellValue("Sheet1", axis, errString)
			} else {
				xlsx.SetCellValue("Sheet1", axis, "Successful")
			}

		}
	}
	xlsx.Save() //保存excel文件

	//进行白名单数据写入操作
	var aff int64 //写入成功的记录条数
	if len(insertData) > 0 {
		aff = models.White{}.AddBatch(&insertData)
	} else {
		aff = 0
	}

	//记录导入记录
	uploadLog := new(models.Upload)
	//uploadLog.FileName = helper.Substr(fileName,10,0)
	uploadLog.FileName = newFileName
	uploadLog.SucceedNumber = aff
	uploadLog.FailureNumber = int64(len(rows)) - aff - 1
	uploadLog.CreateDate = time.Now().Format("2006-01-02 15:04:05")
	models.Upload{}.AddOne(uploadLog)

	w.JSON(iris.Map{
		"code":    0,
		"message": "成功导入 " + strconv.Itoa(int(aff)) + " 条记录; 失败 " + strconv.Itoa(len(rows)-int(aff)-1) + " 条记录！",
		"result":  "",
	})
}

/*
	文件上传操作
*/
func (White) Upload(ctx iris.Context) (string, string) {
	// or use ctx.SetMaxRequestBodySize(10 << 20)
	// to limit the uploaded file(s) size.

	// Get the file from the request
	file, info, err := ctx.FormFile("uploadfile")

	if err != nil {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "Failure"),
			"result":  "",
		})
		return "", ""
	}

	fname := info.Filename

	//只允许上传 .xlsx后缀的文件
	filenameWithSuffix := path.Base(fname)
	fileSuffix := path.Ext(filenameWithSuffix)
	if fileSuffix != ".xlsx" {
		ctx.JSON(iris.Map{
			"code":    400,
			"message": lang.Lang{}.GetLang(ctx, "comm", "UploadFileTypeError_XLSX"),
			"result":  "",
		})
		return "", ""
	}

	//判断上传目录是否存在，不存在则进行创建操作
	fileMain, _ := exec.LookPath(os.Args[0])
	AppPath, _ := filepath.Abs(fileMain)
	mainPath, _ := filepath.Split(AppPath)
	//mainPath,_ := os.Getwd()
	uploadPath := mainPath + "/uploads/"
	_, pathErr := os.Stat(uploadPath)
	if pathErr != nil {
		os.MkdirAll(uploadPath, 0755)
	}

	// Create a file with the same name
	// assuming that you have a folder named 'uploads'
	filePath := uploadPath + fname
	newFileName := helper.GetRandomString(20) + fname
	out, err := os.OpenFile(filePath,
		os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
		return "", ""
	}
	defer out.Close()

	io.Copy(out, file)
	out.Close()
	os.Rename(filePath, uploadPath+newFileName)
	return uploadPath + newFileName, newFileName
}

/**
* @api {Get} /getImportLog 获取导入记录
* @apiDescription 获取导入记录
* @apiGroup white
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiParam {String} token token
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
*			"id": 17,
*			"file_name": "cIBJPVRSH8TLFv4foRKQ001.xlsx",
*			"succeed_number": 0,
*			"failure_number": 3,
*			"create_date": "2017-11-08 11:04:52"
*		  },
*		],
*		"last_page": 1,
*		"per_page": 15,
*		"total": 2
*	  }
*	}
 */
func (White) GetUploadLog(w context.Context) {
	per_page := w.URLParamDefault("per_page", "15")
	page := w.URLParamDefault("page", "1")
	newPage, _ := strconv.Atoi(page)
	nowPage := int32(newPage)

	//分页结算
	upload := new(models.Upload)
	count := models.Upload{}.GetCount(upload) //总的数据量
	number, _ := strconv.Atoi(per_page)       //每页显示的数据条数
	compute := float64(count) / float64(number)
	countPage := math.Ceil(compute)                //计算总的页数
	startPosition := (nowPage - 1) * int32(number) //计算开始位置

	list := make([]models.Upload, 0)
	list, err := models.Upload{}.List(list, upload, int32(number), int32(startPosition))

	//重新声明一个结构体进行重组数据
	type newUploadList struct {
		models.Upload
		OldFileName string `json:"old_file_name"`
	}

	newList := make([]newUploadList, 0)
	var newItemList newUploadList
	for _, item := range list {
		newItemList.CreateDate = item.CreateDate
		newItemList.FailureNumber = item.FailureNumber
		newItemList.SucceedNumber = item.SucceedNumber
		newItemList.FileName = item.FileName
		newItemList.Id = item.Id
		newItemList.OldFileName = helper.Substr(item.FileName, 20, 0)
		newList = append(newList, newItemList)
	}

	if err != nil {
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
			"data":         newList,
			"current_page": nowPage,   //当前页码
			"last_page":    countPage, //总共多少页
			"per_page":     number,    //每页多少条数据
			"total":        count,     //总的记录条数
		},
	})
}

/**
* @api {Get} /down 下载导入文件
* @apiDescription 下载导入文件
* @apiGroup white
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiParam {String} locale 语言
* @apiParam {String} token token
* @apiParam {String} file_name 文件名
*
 */
func (White) Down(w context.Context) {
	file_name := w.FormValue("file_name")
	newFileName := helper.Substr(file_name, 20, 0)
	//获取当前执行文件的路径
	file, _ := exec.LookPath(os.Args[0])
	AppPath, _ := filepath.Abs(file)
	losPath, _ := filepath.Split(AppPath)
	filePath := losPath + "/uploads/" + file_name
	w.SendFile(filePath, newFileName)
}

/*
	推送到MQ操作
*/
func (White) PushMq(body map[string]interface{}) {
	conn, ch := helper.GetRabbitMq()
	defer conn.Close()
	defer ch.Close()
	ch.ExchangeDeclare(
		"browser-manage", // name
		"fanout",         // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	//var body PushMq

	//body := map[string]interface{}{"adds":[]int32{1,3},"dels":[]int32{4},"type":0,"updates":[]int32{5}}
	pushBody, _ := json.Marshal(body)
	err := ch.Publish(
		"browser-manage", // exchange
		"",               // routing key
		false,            // mandatory
		false,            // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        pushBody,
		})
	if err != nil {
		panic(err)
	}
}

/*
	测试消费MQ
*/
func (White) TestMq(w context.Context) {
	conn, ch := helper.GetRabbitMq()
	defer conn.Close()
	defer ch.Close()
	ch.ExchangeDeclare(
		"browser-manage", // name
		"fanout",         // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		nil,              // arguments
	)
	q, err := ch.QueueDeclare(
		"browser_login_queue", // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		panic(err)
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}
}
