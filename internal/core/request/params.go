package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func makeErrorResponse(msg string, code int) (ResponseErrorDto, error) {
	err := errors.New(msg)
	return ResponseErrorDto{
		ResponseDTO: ResponseDTO{
			Message: err.Error(),
			Code:    code,
		},
	}, err
}

func DefaultGetValidParams(c *gin.Context, params interface{}) (ResponseErrorDto, error) {
	if err := c.ShouldBind(params); err != nil {
		return makeErrorResponse(err.Error(), http.StatusBadRequest)
	}

	valid, err := GetValidator(c)
	if err != nil {
		return makeErrorResponse(err.Error(), http.StatusInternalServerError)
	}

	trans, err := GetTranslation(c)
	if err != nil {
		return makeErrorResponse(err.Error(), http.StatusInternalServerError)
	}

	err = valid.Struct(params)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		return HandleValidation(errs, trans)
	}

	return ResponseErrorDto{}, nil
}

func HandleValidation(errs validator.ValidationErrors, trans ut.Translator) (ResponseErrorDto, error) {
	errorData := map[string]interface{}{}

	for _, e := range errs {

		if _, ok := errorData[e.Field()]; ok {
			continue
		}

		var localData []string
		for _, e2 := range errs {
			if e.Namespace() == e2.Namespace() {
				localData = append(localData, e2.Translate(trans))
			}
		}

		errorData[e.Field()] = localData
	}

	text, _ := trans.T("UnprocessableEntity")
	err := errors.New(text)
	responseErrorDto := ResponseErrorDto{
		ResponseDTO: ResponseDTO{
			Message: err.Error(),
			Code:    http.StatusUnprocessableEntity,
		},
		Errors: errorData,
	}

	return responseErrorDto, err
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
