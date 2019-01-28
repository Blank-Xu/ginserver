package resp

type ResponseData struct {
	Data interface{} `json:"data"`
}

func NewResponseData(data interface{}) *ResponseData {
	return &ResponseData{Data: data}
}

type ResponseErr struct {
	Error interface{} `json:"error"`
}

func NewResponseErr(err interface{}) *ResponseErr {
	return &ResponseErr{Error: err}
}

type ResponseCallback struct {
	CallBack string `json:"callback"`
}
