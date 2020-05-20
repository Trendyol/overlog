package over

import (
	"github.com/rs/zerolog"
	"strconv"
	"strings"
)

func LowercaseLevelEncoder() func(l zerolog.Level) string {
	return func(l zerolog.Level) string {
		return l.String()
	}
}

func UppercaseLevelEncoder() func(l zerolog.Level) string {
	return func(l zerolog.Level) string {
		switch l {
		case zerolog.TraceLevel:
			return "TRACE"
		case zerolog.DebugLevel:
			return "DEBUG"
		case zerolog.InfoLevel:
			return "INFO"
		case zerolog.WarnLevel:
			return "WARN"
		case zerolog.ErrorLevel:
			return "ERROR"
		case zerolog.FatalLevel:
			return "FATAL"
		case zerolog.PanicLevel:
			return "PANIC"
		case zerolog.NoLevel:
			return ""
		}
		return ""
	}
}

func FullCallerEncoder() func(file string, line int) string {
	return func(file string, line int) string {
		return file + ":" + strconv.Itoa(line)
	}
}

func ShortCallerEncoder() func(file string, line int) string {
	return func(file string, line int) string {
		return TrimmedPath(file) + ":" + strconv.Itoa(line)
	}
}

func TrimmedPath(file string) string {
	idx := strings.LastIndexByte(file, '/')
	if idx == -1 {
		return file
	}
	idx = strings.LastIndexByte(file[:idx], '/')
	if idx == -1 {
		return file
	}
	return file[idx+1:]
}
