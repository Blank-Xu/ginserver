package error

// code 2000-2999 for database
const (
	CodeDB = 2000
)

func init() {
	regErrCode(CodeDB, "succ")
}
