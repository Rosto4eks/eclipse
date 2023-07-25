package logger

import (
	"fmt"
	"time"
)

const (
	reset  = "\033[0;0m"
	black  = "\033[1;30m"
	red    = "\033[1;31m"
	green  = "\033[1;32m"
	yellow = "\033[1;33m"
	cyan   = "\033[1;36m"
	purple = "\033[1;35m"

	redBg    = "\033[41m"
	greenBg  = "\033[42m"
	yellowBg = "\033[43m"
	cyanBg   = "\033[46m"
	purpleBg = "\033[45m"
)

type Ilogger interface {
	Error(string, string, error)
	Info(string, string, string)
	Warning(string, string, string)
	Debug(string, string, string)
}

type logger struct {
	level int // 1 - DEBUG 2 - DEV 3 - PROD
}

func New(level string) *logger {
	var lvl int
	switch level {
	case "DEBUG":
		lvl = 1
	case "DEV":
		lvl = 2
	case "PROD":
		lvl = 3
	default:
		lvl = 1
	}
	return &logger{
		level: lvl,
	}
}

func (l logger) print(level int, levelStr, module, function string, str string) {
	if level >= l.level {
		fmt.Printf("%s %v  %s%s %s %s -> %s%s %s %s -> %s\n", levelStr, time.Now().Format("2006-01-02 15:04:05"), cyanBg, black, module, reset, yellowBg, black, function, reset, str)
	}
}

func (l logger) Error(module, function string, err error) {
	l.print(1, red+"ERROR"+reset, module, function, err.Error())
}

func (l logger) Debug(module, function string, str string) {
	l.print(1, yellow+"DEBUG"+reset, module, function, str)
}

func (l logger) Warning(module, function string, str string) {
	l.print(2, yellow+"WARN "+reset, module, function, str)
}

func (l logger) Info(module, function string, str string) {
	l.print(3, green+"INFO "+reset, module, function, str)
}
