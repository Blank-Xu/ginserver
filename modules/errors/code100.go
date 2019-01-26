package errors

// code 104-199 for status
const (
	Code104    = 104
	Code104Msg = "104"
)

func init() {
	registerErrMsg(Code104, Code104Msg)
}
