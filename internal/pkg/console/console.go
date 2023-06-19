package console

const (
	_ = iota
	DebugLevel
	InfoLevel
	WarnLevel
	ErrorLevel

	NoneLevel
)

var level = NoneLevel

func Debug() {

}
