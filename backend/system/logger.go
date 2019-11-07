package system

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/rs/zerolog"
)

type AppLogger interface {
	NewLogger(funcName string) zerolog.Logger
}

type appLogger struct {
	output zerolog.ConsoleWriter
}

// For Local
func NewAppLogger() AppLogger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05.000000"}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatMessage = func(i interface{}) string {
		if i == nil {
			return ""
		}
		return fmt.Sprintf("***%s****", i)
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	output.FormatCaller = func(i interface{}) string {
		t := fmt.Sprintf("%s", i)
		s := strings.Split(t, ":")
		if 2 != len(s) {
			return t
		}
		d := filepath.Dir(s[0])
		_, pkg := filepath.Split(d)
		f := filepath.Base(s[0])
		return fmt.Sprintf("%s/%s:%s", pkg, f, s[1])
	}

	return &appLogger{output: output}
}

func (l *appLogger) NewLogger(funcName string) zerolog.Logger {
	return zerolog.New(l.output).With().Str("function", funcName).Timestamp().Caller().Logger()
}

func fileWithLineNum() string {
	_, name, line, ok := runtime.Caller(2)
	if !ok {
		return "-"
	}
	return fmt.Sprintf("%s:%d", name, line)
}
