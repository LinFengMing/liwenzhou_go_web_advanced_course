package controllers

import (
	"bluebell/models"
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func InitTrans(locale string) (err error) {
	// 修改 gin 的 Validator 的翻譯器，使其支持多國語言
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 註冊一個獲得 json tag 的自定義方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		// ParamSigUp 註冊自定義驗證函數
		v.RegisterStructValidation(ParamSigUpStructLevelValidation, models.ParamSigUp{})
		zhT := zh.New()
		enT := en.New()
		// 第一個參數是備用（fallback）的語言翻譯器
		// 後面的參數是應該支持的語言翻譯器
		uni := ut.New(enT, zhT, enT)
		// locale 通常取決於 http header 的 'Accept-Language'
		var ok bool
		// 也可以使用 uni.FindTranslator(...) 傳入多個 locale 進行尋找
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		// 註冊翻譯器
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

// 去除提示訊息中的 struct 名稱
func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// 自定義 ParamSigUp struct 驗證函數
func ParamSigUpStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(models.ParamSigUp)
	if su.Password != su.RePassword {
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}
