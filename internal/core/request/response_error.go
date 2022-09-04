package request

type ResponseErrorDto struct {
	ResponseDTO
	Errors map[string]interface{} `json:"errors"`
}
