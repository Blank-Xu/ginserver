package error

// code 230-299
const (
	Code230    = 230
	Code230Msg = "230"
)

func init() {
	registerErrMsg(Code230, Code230Msg)
}
