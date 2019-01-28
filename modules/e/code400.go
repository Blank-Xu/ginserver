package e

// code 4000-4999
const (
	CodeParamInvalid = 4001
)

func init() {
	registerErrMsg(CodeParamInvalid, "param invalid")
}
