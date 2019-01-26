package resp

type ResponseData struct {
	Data interface{} `json:"data"`
}

type ResponseErr struct {
	Error interface{} `json:"error"`
}

type ResponseCallback struct {
	CallBack string `json:"callback"`
}
