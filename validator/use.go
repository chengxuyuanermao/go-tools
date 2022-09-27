package validator

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

func Use() {
	validate := validator.New()
	// 验证变量
	email := "admin#admin.com" // 报错：Key: '' Error:Field validation for '' failed on the 'email' tag
	//email := "" // 报错：Key: '' Error:Field validation for '' failed on the 'required' tag
	err := validate.Var(email, "required,email")
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		fmt.Println(validationErrors)
		return
	}
}

func UseStruct() {
	// 获取一个验证器实例
	validate := validator.New()

	// 注册函数，获取结构体字段的备用名称
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return "j"
		}
		return name
	})
	// 将错误信息翻译成中文：
	zh1 := zh.New()
	uni := ut.New(zh1)
	trans, _ := uni.GetTranslator("zh")
	_ = zh_translations.RegisterDefaultTranslations(validate, trans)

	// 定义一个结构体
	type User struct {
		ID     int64  `json:"id" validate:"gt=0"`
		Name   string `json:"name" validate:"required"`
		Gender string `json:"gender" validate:"required,oneof=man woman"`
		Age    uint8  `json:"age" validate:"required,gte=0,lte=130"`
		Email  string `json:"email" validate:"required,email"`
	}
	user := &User{
		ID:     1,
		Name:   "frank",
		Gender: "boy",
		Age:    180,
		Email:  "gopher@88.com",
	}
	// 验证结构体值
	err := validate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)

		// 没翻译前,报错信息：
		//Key: 'User.gender' Error:Field validation for 'gender' failed on the 'oneof' tag
		//Key: 'User.age' Error:Field validation for 'age' failed on the 'lte' tag
		fmt.Println(validationErrors)

		// 翻译后报错信息：map[User.age:age必须小于或等于130 User.gender:gender必须是[man woman]中的一个]
		fmt.Println(validationErrors.Translate(trans))
		return
	}
}
