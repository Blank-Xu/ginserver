package e

// code 5000-5999
const (
	CodeDBErr = 5011
)

func init() {
	registerErrMsg(CodeDBErr, "Database Error")
}
