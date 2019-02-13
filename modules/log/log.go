package log

type Type int

const (
	TypeLogin Type = iota + 1
	TypeLogout
)

type Level int

const (
	// 0:panic, 1:fatal, 2:error, 3:warn, 4:info, 5:debug, 6:trace
	LevelPanic Level = iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
)