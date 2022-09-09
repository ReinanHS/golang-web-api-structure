package request

type ResponseDataDto struct {
	ResponseDTO
	Data interface{} `json:"data"`
}
