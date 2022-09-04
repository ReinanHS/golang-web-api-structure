package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

func DefaultGetValidParams(c *gin.Context, params interface{}) (map[string]interface{}, error) {
	errorData := map[string]interface{}{}
	if err := c.ShouldBind(params); err != nil {
		return errorData, err
	}

	valid, err := GetValidator(c)
	if err != nil {
		return errorData, err
	}

	trans, err := GetTranslation(c)
	if err != nil {
		return errorData, err
	}

	err = valid.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		return HandleValidation(errs, trans)
	}

	return errorData, nil
}

func HandleValidation(errs validator.ValidationErrors, trans ut.Translator) (map[string]interface{}, error) {
	errorData := map[string]interface{}{}

	for _, e := range errs {

		if _, ok := errorData[e.Field()]; ok {
			continue
		}

		var localData = []string{}

		for _, e2 := range errs {
			if e.Namespace() == e2.Namespace() {
				localData = append(localData, e2.Translate(trans))
			}
		}

		errorData[e.Field()] = localData
	}

	return errorData, errors.New("there was a validation error")
}

func GetValidator(c *gin.Context) (*validator.Validate, error) {
	val, ok := c.Get("ValidatorKey")
	if !ok {
		return nil, errors.New("could not find the variable: ValidatorKey")
	}
	valid, ok := val.(*validator.Validate)
	if !ok {
		return nil, errors.New("unable to perform type conversion")
	}
	return valid, nil
}

func GetTranslation(c *gin.Context) (ut.Translator, error) {
	trans, ok := c.Get("TranslatorKey")
	if !ok {
		return nil, errors.New("could not find the variable: TranslatorKey")
	}
	translator, ok := trans.(ut.Translator)
	if !ok {
		return nil, errors.New("unable to perform type conversion")
	}
	return translator, nil
}
