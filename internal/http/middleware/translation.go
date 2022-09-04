package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	pt_BR_translations "github.com/go-playground/validator/v10/translations/pt_BR"
	"log"
)

func TranslationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		enLang := en.New()
		ptBrLang := pt_BR.New()

		uni := ut.New(ptBrLang, ptBrLang, enLang)
		val := validator.New()

		locale := c.DefaultQuery("locale", "pt_BR")
		trans, _ := uni.GetTranslator(locale)

		switch locale {
		case "en":
			err := en_translations.RegisterDefaultTranslations(val, trans)
			if err != nil {
				log.Fatal(err)
			}
			break
		default:
			err := pt_BR_translations.RegisterDefaultTranslations(val, trans)
			if err != nil {
				log.Fatal(err)
			}
			break
		}
		c.Set("TranslatorKey", trans)
		c.Set("ValidatorKey", val)
		c.Next()
	}
}
