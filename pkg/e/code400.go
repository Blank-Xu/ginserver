package e

// code 4000-4999
const (
	CodeInvalidParams = 4001
)

func init() {
	registerErrMsg(CodeInvalidParams, "Invalid Params")
}
