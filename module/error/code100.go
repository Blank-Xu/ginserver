package error

// code 1000-1999 for status
const (
	CodeSuccess = 1000
)

func init() {
	regErrCode(CodeSuccess, "succ")
}
