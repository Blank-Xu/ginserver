package resp

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
