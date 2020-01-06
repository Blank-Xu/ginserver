package log

type Type int

const (
	TypeInternalServerError Type = 10000 + iota
	TypeDBError
)

const (
	TypeOther Type = iota
	TypeLogin
	TypeLogout
	TypeChangePwd
)
