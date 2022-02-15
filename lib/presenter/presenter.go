package presenter

const (
	RESPONSE_SUCCESS_GET      = "Success Get Data"
	RESPONSE_SUCCESS_GET_LIST = "Success Get List Data"
	RESPONSE_SUCCESS_CREATE   = "Success Create Data"
	RESPONSE_SUCCESS_UPDATE   = "Success Update Data"
	RESPONSE_SUCCESS_DELETE   = "Success Delete Data"
	RESPONSE_SUCCESS_VOID     = "Success Void Data"

	RESPONSE_ERROR_INVALID_PARAM_NUMBER = "should be a number"
)

type Default struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type List struct {
	Message string      `json:"message"`
	Meta    Meta        `json:"meta"`
	Data    interface{} `json:"data"`
}

type Meta struct {
	TotalData int `json:"total_data"`
	TotalPage int `json:"total_page"`
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
}
