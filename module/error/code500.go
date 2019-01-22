package error

const (
	Code500    = 500
	Code500Err = "Internal Server Error"
)

func init() {
	regErrCode(Code500, Code500Err)
}
