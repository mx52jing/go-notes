package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	UserName string `json:"user_name" form:"user_name" uri:"user_name"`
	Age      int    `json:"age" form:"age" uri:"age"`
	Sex      string `json:"sex" form:"sex" uri:"sex"`
}

func bodyBindHandler(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		fmt.Println("err", err)
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "JSON绑定出错"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func queryBindHandler(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBindQuery(&user)
	if err != nil {
		fmt.Println("err", err)
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "JSON绑定出错"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func uriBindHandler(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBindUri(&user)
	if err != nil {
		fmt.Println("err", err)
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "JSON绑定出错"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func bindHandler(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBind(&user)
	if err != nil {
		fmt.Println("err", err)
		ctx.JSON(http.StatusOK, gin.H{"code": -1, "message": "JSON绑定出错"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

type Student struct {
	Name           string    `json:"name" form:"name" binding:"required" msg:"name字段必传"`                               // name字段必须传
	Score          int       `json:"score" form:"score" binding:"required,gt=0" msg:"score字段必传" gt_msg:"score字段必须大于0"` // score 必须大于0
	AdmissionDate  time.Time `json:"admission_date" form:"admission_date" binding:"required,before_today" time_format:"2006-01-02" time_utc:"8" msg:"admission_date字段必传" before_today_msg:"admission_date日期必须在当日之前"`
	GraduationDate time.Time `json:"graduation_date" form:"graduation_date" binding:"required,gtfield=AdmissionDate" time_format:"2006-01-02" time_utc:"8" msg:"graduation_date字段必传" gtfield_msg:"graduation_date必须大于admission_date"`
}

// 自定义验证器 校验日期必须在今天之前
var beforeToday validator.Func = func(fl validator.FieldLevel) bool {
	if date, ok := fl.Field().Interface().(time.Time); ok { // 通过反射获得结构体Field的值
		todayDate := time.Now()
		return date.Before(todayDate)
	}
	return false
}

func studentHandler(ctx *gin.Context) {
	var student Student
	// 绑定参数并且完成参数校验
	if err := ctx.ShouldBind(&student); err != nil {
		errMsg := parseErrorMessage(err, &student)
		ctx.JSON(http.StatusOK, gin.H{"msg": errMsg})
		return
	}
	ctx.JSON(http.StatusOK, student)
}

// 获取参数报错具体信息
func parseErrorMessage(err error, originalStruct any) string {
	errs, ok := err.(validator.ValidationErrors)
	// 如果不是校验器校验失败的错误，直接返回错误信息
	if !ok {
		return err.Error()
	}
	var result string
	// 利用反射获取结构体的值
	structData := reflect.TypeOf(originalStruct).Elem()
	for _, errItem := range errs {
		// 获取错误字段在结构体中的key 比如Name/Score
		errFiledName := errItem.Field()
		// 获取当前错误具体的tag名称 比如是required 还是 gtfield
		errFieldTag := errItem.Tag()
		// 获取当前field的tag信息
		var errMsg string
		if field, ok := structData.FieldByName(errFiledName); ok {
			// 先获取当前错误tag的key
			errMsg = field.Tag.Get(fmt.Sprintf("%s_msg", errFieldTag))
			// 当错误不是其他约束错误时 取默认的错误信息
			if len(errMsg) == 0 {
				errMsg = field.Tag.Get("msg")
			}
			var separator string
			if len(errMsg) > 0 {
				separator = "，"
			}
			result += fmt.Sprintf("%s%s", separator, errMsg)
		}
	}
	return result
}

func main() {
	// gin中的bind可以很方便的将 前端传递 来的数据与 结构体 进行 参数绑定 ，以及参数校验
	// 在使用这个功能的时候，需要给结构体加上Tag json form uri xml yaml
	// Must Bind
	router := gin.Default()
	// 注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("before_today", beforeToday)
	}
	router.POST("/body_bind", bodyBindHandler)
	router.GET("/query_bind", queryBindHandler)
	router.GET("/uri_bind/:user_name/:age/:sex", uriBindHandler)
	router.POST("/bind", bindHandler)
	router.GET("/bind", bindHandler)
	router.GET("/s", studentHandler)
	router.Run(":9100")
}
