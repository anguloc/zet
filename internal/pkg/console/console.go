package console

import (
	"github.com/gookit/color"
)

type Level uint32

const (
	DebugLevel Level = 1 << iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

var level Level

func Debug(a ...any) {
	if level&DebugLevel == DebugLevel {
		return
	}
	color.Blue.Println(a...)
}

func Debugf(format string, a ...any) {
	if level&DebugLevel == DebugLevel {
		return
	}
	color.Blue.Printf(format, a...)
}

func Info(a ...any) {
	if level&InfoLevel == InfoLevel {
		return
	}
	color.Green.Println(a...)
}

func Infof(format string, a ...any) {
	if level&InfoLevel == InfoLevel {
		return
	}
	color.Green.Printf(format, a...)
}

func Warn(a ...any) {
	if level&WarnLevel == WarnLevel {
		return
	}
	color.Yellow.Println(a...)
}

func Warnf(format string, a ...any) {
	if level&WarnLevel == WarnLevel {
		return
	}
	color.Yellow.Printf(format, a...)
}

func Error(a ...any) {
	if level&ErrorLevel == ErrorLevel {
		return
	}
	color.Red.Println(a...)
}

func Errorf(format string, a ...any) {
	if level&ErrorLevel == ErrorLevel {
		return
	}
	color.Red.Printf(format, a...)
}

func SetLevel(ls ...Level) {
	for _, v := range ls {
		level += v
	}
}
