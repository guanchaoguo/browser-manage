package controllers

//验证码控制器

import (
	"browser-manage/app/helper"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/lifei6671/gocaptcha"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Captcaha struct {
}

const (
	dx = 154
	dy = 58
)

/**
* @api {Get} /getCaptcha 获取验证码
* @apiDescription 获取验证码
* @apiGroup captcha
* @apiPermission JWT
* @apiVersion 1.0.0
* @apiSuccessExample {json} Success-Response:
*	{
*	  "code": 0,
*	  "message": "操作成功",
*	  "result": {
*		"url": "127.0.0.1:8080/captcha/njp674vTtzlvolXVptyU7qA5zzqNiiW2"
*	  }
*	}
*
 */
func (Captcaha) GetCaptcha(w context.Context) {

	//生成随机数
	urlCode := helper.GetRandomString(32)
	//存放到redis中
	r := helper.GetRedis()
	_, err := r.Do("SET", urlCode, "")
	defer r.Close()

	if err != nil {
		w.JSON(iris.Map{
			"code":    400,
			"message": "操作失败",
			"result":  "",
		})
		w.Request().Body.Close()
		return
	}
	//成功返回地址
	w.JSON(iris.Map{
		"code":    0,
		"message": "操作成功",
		"result": iris.Map{
			"url": "/captcha/" + urlCode,
		},
	})

}

func (Captcaha) Get(w context.Context) {
	urlCode := w.Params().Get("code")
	//验证code是否合法
	r := helper.GetRedis()
	is_key_exit, _ := redis.Bool(r.Do("EXISTS", urlCode))
	if is_key_exit == false {
		defer r.Close()
		w.StatusCode(500)
		return
	}
	//获取当前执行文件的路径
	file, _ := exec.LookPath(os.Args[0])
	AppPath, _ := filepath.Abs(file)
	fontPath, _ := filepath.Split(AppPath)
	//fontFils, err := ListDir(fontPath+"/fonts", ".ttf")
	//if err != nil {
	//	defer r.Close()
	//	fmt.Println(err)
	//	return
	//}
	//gocaptcha.SetFontFamily(fontFils...)
	gocaptcha.SetFontFamily(fontPath + "/fonts/captcha0.ttf")
	captchaImage, err := gocaptcha.NewCaptchaImage(dx, dy, gocaptcha.RandLightColor())
	//captchaImage.DrawNoise(gocaptcha.CaptchaComplexHigh)
	//captchaImage.DrawTextNoise(gocaptcha.CaptchaComplexHigh)
	randText := gocaptcha.RandText(4)
	captchaImage.DrawText(randText)
	//captchaImage.Drawline(3)
	captchaImage.DrawBorder(gocaptcha.ColorToRGB(0x17A7A7A))
	captchaImage.DrawHollowLine()
	if err != nil {
		defer r.Close()
		fmt.Println(err)
		w.Request().Body.Close() //关闭本次请求
		return
	}
	//把验证码转成大写存放到redis中
	_, err = r.Do("SET", urlCode, strings.ToUpper(randText))
	defer r.Close()

	if err != nil {
		fmt.Println(err)
		w.Request().Body.Close() //关闭本次请求
		return
	}
	captchaImage.SaveImage(w, gocaptcha.ImageFormatJpeg)
}

//检测验证码是否正确
func (Captcaha) CheckCaptcha(code string, captcha string) bool {
	//获取redis中的code
	r := helper.GetRedis()
	redis_captcha, _ := redis.String(r.Do("GET", code))
	//删除本次key(一次性)
	r.Do("DEL", code)
	defer r.Close() //关闭redis

	if redis_captcha != strings.ToUpper(captcha) {
		return false
	}
	return true
}

//获取指定目录下的所有文件，不进入下一级目录搜索，可以匹配后缀过滤。
func ListDir(dirPth string, suffix string) (files []string, err error) {
	files = make([]string, 0, 10)
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}
	PthSep := string(os.PathSeparator)
	suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写
	for _, fi := range dir {
		if fi.IsDir() { // 忽略目录
			continue
		}
		if strings.HasSuffix(strings.ToUpper(fi.Name()), suffix) { //匹配文件
			files = append(files, dirPth+PthSep+fi.Name())
		}
	}
	return files, nil
}
